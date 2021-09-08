package valid

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
	"io"
	"time"
)

var wordSafeBase32 = base32.NewEncoding("23456789CFGHJMPQRVWXcfghjmpqrvwx").WithPadding(base32.NoPadding)

type Source interface {
	WithChecksum() Source
	WithTimestamp() Source
	WithTimestampStartAt(initialTime time.Time) Source
	Generate() string
	Timestamp(id string) time.Time
	IsValid(id string) bool
}

type baseSource struct {
	randomByteLength int
	useChecksum      bool
	useTimestamp     bool
	initialTime      time.Time
}

func (s baseSource) WithChecksum() Source {
	s.useChecksum = true
	return s
}

func (s baseSource) WithTimestamp() Source {
	s.useTimestamp = true
	return s
}

func (s baseSource) WithTimestampStartAt(initialTime time.Time) Source {
	s.useTimestamp = true
	s.initialTime = initialTime
	return s
}

func NewSource(randomByteLength int) Source {
	if randomByteLength < 1 {
		panic("random byte length < 1")
	}
	return baseSource{
		randomByteLength: randomByteLength,
		useChecksum:      false,
		useTimestamp:     false,
		initialTime:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func (s baseSource) Generate() string {
	b := make([]byte, s.byteLength())
	cursor := 0

	// Timestamp
	if s.useTimestamp {
		// 40bit timestamp ~ Millisecond
		ts := (time.Now().UnixNano() - s.initialTime.UnixNano()) >> 20 & ((1 << 40) - 1)
		b[0] = byte(ts >> 32)
		b[1] = byte(ts >> 24)
		b[2] = byte(ts >> 16)
		b[3] = byte(ts >> 8)
		b[4] = byte(ts)
		cursor = 5
	}

	// Random value
	_, err := io.ReadFull(rand.Reader, b[cursor:cursor+s.randomByteLength])
	if err != nil {
		panic("failed to read random value")
	}
	cursor += s.randomByteLength

	// Checksum
	if s.useChecksum {
		for _, bb := range b[:cursor] {
			b[cursor] ^= bb
		}
	}

	return wordSafeBase32.EncodeToString(b)
}

func (s baseSource) Timestamp(id string) time.Time {
	if !s.useTimestamp {
		return time.Time{}
	}
	b, err := wordSafeBase32.DecodeString(id)
	if err != nil {
		return time.Time{}
	}

	ts := make([]byte, 8)
	ts[3] = b[0]
	ts[4] = b[1]
	ts[5] = b[2]
	ts[6] = b[3]
	ts[7] = b[4]

	val := binary.BigEndian.Uint64(ts) << 20
	return time.Unix(0, s.initialTime.UnixNano()+int64(val))
}

func (s baseSource) IsValid(id string) bool {
	if !s.useChecksum {
		return false // No UseChecksum
	}

	b, err := wordSafeBase32.DecodeString(id)
	if err != nil {
		return false
	}

	if len(b) != s.byteLength() {
		return false
	}

	var checksum byte
	for _, bb := range b {
		checksum ^= bb
	}

	return checksum == byte(0)
}

func (s baseSource) byteLength() int {
	l := s.randomByteLength
	if s.useTimestamp {
		l += 5
	}
	if s.useChecksum {
		l++
	}
	return l
}
