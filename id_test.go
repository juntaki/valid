package valid_test

import (
	"testing"
	"time"

	"github.com/juntaki/valid"
)

var ds = valid.NewSource(16).WithTimestamp().WithChecksum()

func Test_generateID(t *testing.T) {
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

func Test_validateID(t *testing.T) {
	bad := "2X75jjvg2zzzGWq9JHjm88PHvR49gQX2FGqH6"
	if ds.IsValid(bad) {
		t.Fatal("invalid")
	}
}

func Benchmark_generateID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ds.Generate()
	}
}

func Benchmark_validateID(b *testing.B) {
	got := ds.Generate()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ds.IsValid(got)
	}
}
