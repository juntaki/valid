package valid

import (
	"time"
)

type MockSource struct {
	RandomByteLength int
	UseChecksum      bool
	UseTimestamp     bool
	DummyTimestamp   time.Time
}

func NewMockSource(randomByteLength int) *MockSource {
	return &MockSource{
		RandomByteLength: randomByteLength,
		UseChecksum:      false,
		UseTimestamp:     false,
		DummyTimestamp:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func (m MockSource) WithChecksum() Source {
	m.UseChecksum = true
	return m
}

func (m MockSource) WithTimestamp() Source {
	m.UseTimestamp = true
	return m
}

func (m MockSource) WithTimestampStartAt(initialTime time.Time) Source {
	m.UseTimestamp = true
	return m
}

func (m MockSource) Generate() string {
	b := make([]byte, m.byteLength())
	return wordSafeBase32.EncodeToString(b)
}

func (m MockSource) Timestamp(id string) time.Time {
	return m.DummyTimestamp
}

func (m MockSource) IsValid(id string) bool {
	return true
}

func (m MockSource) byteLength() int {
	l := m.RandomByteLength
	if m.UseTimestamp {
		l += 5
	}
	if m.UseChecksum {
		l++
	}
	return l
}
