package valid_test

import (
	"testing"
	"time"

	"github.com/juntaki/valid"
)

var ds = valid.NewSource(16).WithTimestamp().WithChecksum()

func TestCustomGenerate(t *testing.T) {
	for i := 1; i < 100; i++ {
		ds = valid.NewSource(i).WithTimestamp().WithChecksum()
		got := ds.Generate()
		if time.Since(ds.Timestamp(got)).Milliseconds() > 1 {
			t.Fatal("invalid")
		}
		if !ds.IsValid(got) {
			t.Fatal("invalid")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestCustomIsValid(t *testing.T) {
	bad := "2X75jjvg2zzzGWq9JHjm88PHvR49gQX2FGqH6"
	if ds.IsValid(bad) {
		t.Fatal("invalid")
	}
}

func TestGenerate(t *testing.T) {
	for i := 1; i < 100; i++ {
		got := valid.Generate()
		if time.Since(valid.Timestamp(got)).Milliseconds() > 1 {
			t.Fatal("invalid")
		}
		if !valid.IsValid(got) {
			t.Fatal("invalid")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestMockGenerate(t *testing.T) {
	mock := valid.NewMockSource(16)
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

func BenchmarkIsValid(b *testing.B) {
	id := valid.Generate()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid.IsValid(id)
	}
}

func BenchmarkGenerate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid.Generate()
	}
}
