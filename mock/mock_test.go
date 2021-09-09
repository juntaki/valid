package mock_test

import (
	"testing"
	"time"

	"github.com/juntaki/valid/mock"
)

func TestMockGenerate(t *testing.T) {
	mock := mock.NewMockSource(16)
	ref := mock.Generate()
	for i := 1; i < 100; i++ {
		got := mock.Generate()
		if ref != got {
			t.Fatal("invalid")
		}
		if mock.Timestamp(got) != mock.Timestamp(ref) {
			t.Fatal("invalid")
		}
		if !mock.IsValid(got) {
			t.Fatal("invalid")
		}
		time.Sleep(time.Millisecond)
	}
}
