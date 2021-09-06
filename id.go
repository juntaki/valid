package valid

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"time"
)

var wordSafeBase32 = base32.NewEncoding("23456789CFGHJMPQRVWXcfghjmpqrvwx").WithPadding(base32.NoPadding)

type IDSource struct {
	RandomByteLength int
	Checksum         bool
	Sortable         bool
	InitialTime      time.Time
}

func (s IDSource) checksumByteLength() int {
	if s.Checksum {
		return 1
	}
	return 0
}

func (s IDSource) sortableByteLength() int {
	if s.Sortable {
		return 6
	}
	return 0
}

func (s IDSource) totalByteLength() int {
	return s.sortableByteLength() + s.RandomByteLength + s.checksumByteLength()
}

func (s IDSource) randomByteOffset() int {
	return s.sortableByteLength() + s.RandomByteLength
}

var DefaultSource = IDSource{
	RandomByteLength: 16, // 128bit
	Checksum:         true,
	Sortable:         true,
	InitialTime:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
}

func (s IDSource) Generate() string {
	b := make([]byte, s.totalByteLength())

	// Timestamp
	if s.Sortable {
		timestamp := b[:s.sortableByteLength()]
		// 40bit timestamp
		ts := (time.Now().UnixNano() - s.InitialTime.UnixNano()) >> 20 & ((1 << 40) - 1)

		timestamp[0] = byte(ts >> 32)
		timestamp[1] = byte(ts >> 24)
		timestamp[2] = byte(ts >> 16)
		timestamp[3] = byte(ts >> 8)
		timestamp[4] = byte(ts)
	}

	// Random value
	random := b[s.sortableByteLength():s.randomByteOffset()]
	_, err := rand.Reader.Read(random)
	if err != nil {
		panic(fmt.Errorf("failed to read random value"))
	}

	// Checksum
	if s.Checksum {
		checksum := b[s.randomByteOffset():]
		for _, b := range b[:s.randomByteOffset()] {
			checksum[0] = checksum[0] ^ b
		}
	}

	return wordSafeBase32.EncodeToString(b)
}

func (s IDSource) IsValid(id string) bool {
	b, err := wordSafeBase32.DecodeString(id)
	if err != nil {
		return false
	}

	var checksum byte
	checksumPart := b[s.randomByteOffset():]
	for _, b := range b[:s.randomByteOffset()] {
		checksum = checksum ^ b
	}

	return checksum == checksumPart[0]
}
