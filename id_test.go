package valid_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/juntaki/valid"
)

func TestCustomGenerate(t *testing.T) {
	var ds = valid.NewSource(16).WithTimestamp().WithChecksum()
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
	var ds = valid.NewSource(16)
	id := ds.Generate()
	if ds.IsValid(id) {
		t.Fatal("invalid")
	}
}

func TestCustomTimestamp(t *testing.T) {
	var ds = valid.NewSource(16)
	id := ds.Generate()
	if ds.Timestamp(id).UnixNano() != (time.Time{}).UnixNano() {
		t.Fatal("invalid")
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

func TestNewSource(t *testing.T) {
	type args struct {
		randomByteLength int
	}
	tests := []struct {
		name  string
		args  args
		panic string
	}{
		{
			name: "valid",
			args: args{
				randomByteLength: 16,
			},
		},
		{
			name: "bad",
			args: args{
				randomByteLength: 0,
			},
			panic: "random byte length < 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					if fmt.Sprintf("%s", err) != tt.panic {
						t.Errorf("got %v\nwant %v", err, tt.panic)
					}
				}
			}()
		})
	}
}
