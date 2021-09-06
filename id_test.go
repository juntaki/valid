package valid

import (
	"fmt"
	"testing"
	"time"
)

func Test_generateID(t *testing.T) {
	for i := 0; i < 100; i++ {
		got := DefaultSource.Generate()
		fmt.Println(got, len(got))
		time.Sleep(time.Millisecond)
		if !DefaultSource.IsValid(got) {
			t.Fatal("invalid")
		}
	}
}

func Benchmark_generateID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DefaultSource.Generate()
	}
}

func Benchmark_validateID(b *testing.B) {
	got := DefaultSource.Generate()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DefaultSource.IsValid(got)
	}
}
