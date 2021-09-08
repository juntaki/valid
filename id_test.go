package valid_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/google/uuid"
	"github.com/juntaki/valid"
	"github.com/rs/xid"
)

var ds = valid.NewSource(16).WithTimestamp().WithChecksum()

func Example() {
	id := valid.Generate()

	fmt.Println(id, len(id), valid.IsValid(id), valid.Timestamp(id))
	// Output: 2X9P75FX2pJqCcHRVWV2862JW6XhFr6x 32 true 2021-09-08 23:59:17.339250688 +0900 JST
}

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

func BenchmarkIsValid(b *testing.B) {
	id := valid.Generate()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid.IsValid(id)
	}
}

func BenchmarkUUIDv4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uuid.Must(uuid.NewRandom()).String()
	}
}

func BenchmarkXID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xid.New().String()
	}
}

func BenchmarkValidGenerate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = valid.Generate()
	}
}

func BenchmarkULID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader)
	}
}
