package valid

import (
	"time"
)

type mockSource struct {
	randomByteLength int
	useChecksum      bool
	useTimestamp     bool
	dummyTimestamp   time.Time
}

func NewMockSource(randomByteLength int) Source {
	return &mockSource{
		randomByteLength: randomByteLength,
		useChecksum:      false,
		useTimestamp:     false,
		dummyTimestamp:   time.Date(2021, 9, 4, 0, 0, 0, 0, time.UTC),
	}
}

func (m mockSource) WithChecksum() Source {
	m.useChecksum = true
	return m
}

func (m mockSource) WithTimestamp() Source {
	m.useTimestamp = true
	return m
}

func (m mockSource) Generate() string {
	b := make([]byte, m.byteLength())
	return wordSafeBase32.EncodeToString(b)
}

func (m mockSource) Timestamp(id string) time.Time {
	return m.dummyTimestamp
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
