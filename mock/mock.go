package mock

import (
	"encoding/base32"
	"time"

	"github.com/juntaki/valid"
)

var wordSafeBase32 = base32.NewEncoding("23456789CFGHJMPQRVWXcfghjmpqrvwx").WithPadding(base32.NoPadding)
var dummyTime = time.Date(2021, 9, 4, 0, 0, 0, 0, time.UTC)

func SetDummyTime(t time.Time) {
	dummyTime = t
}

type mockSource struct {
	randomByteLength int
	useChecksum      bool
	useTimestamp     bool
}

func NewMockSource(randomByteLength int) valid.Source {
	return &mockSource{
		randomByteLength: randomByteLength,
		useChecksum:      false,
		useTimestamp:     false,
	}
}

func (m mockSource) WithChecksum() valid.Source {
	m.useChecksum = true
	return m
}

func (m mockSource) WithTimestamp() valid.Source {
	m.useTimestamp = true
	return m
}

func (m mockSource) Generate() string {
	b := make([]byte, m.byteLength())
	return wordSafeBase32.EncodeToString(b)
}

func (m mockSource) Timestamp(id string) time.Time {
	return dummyTime
}

func (m mockSource) IsValid(id string) bool {
	return true
}

func (m mockSource) byteLength() int {
	l := m.randomByteLength
	if m.useTimestamp {
		l += 5
	}
	if m.useChecksum {
		l++
	}
	return l
}
