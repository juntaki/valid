package valid_test

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/google/uuid"
	"github.com/juntaki/valid"
	"github.com/rs/xid"
)

var ds = valid.NewSource(16).WithTimestamp().WithChecksum()

func Test_generateID(t *testing.T) {
	for i := 1; i < 100; i++ {
		ds = valid.NewSource(14).WithTimestamp().WithChecksum()
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

const encoding = "23456789CFGHJMPQRVWXcfghjmpqrvwx"

func encode(dst, id []byte) {

	for i := 0; i < len(id)/5; i++ {
		dst[7+8*i] = encoding[id[4+5*i]&0x1F]
		dst[6+8*i] = encoding[id[4+5*i]>>5|(id[3+5*i]<<3)&0x1F]
		dst[5+8*i] = encoding[(id[3+5*i]>>2)&0x1F]
		dst[4+8*i] = encoding[id[3+5*i]>>7|(id[2+5*i]<<1)&0x1F]
		dst[3+8*i] = encoding[(id[2+5*i]>>4)&0x1F|(id[1+5*i]<<4)&0x1F]
		dst[2+8*i] = encoding[(id[1+5*i]>>1)&0x1F]
		dst[1+8*i] = encoding[(id[1+5*i]>>6)&0x1F|(id[0+5*i]<<2)&0x1F]
		dst[0+8*i] = encoding[id[0+5*i]>>3]
	}
}

func BenchmarkStdlibBASE32(b *testing.B) {
	var wordSafeBase32 = base32.NewEncoding("23456789CFGHJMPQRVWXcfghjmpqrvwx").WithPadding(base32.NoPadding)
	bb := make([]byte, 20)
	_, err := io.ReadFull(rand.Reader, bb)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = wordSafeBase32.EncodeToString(bb)
	}
}

func BenchmarkOrigEncoder(b *testing.B) {
	bb := make([]byte, 20)
	_, err := io.ReadFull(rand.Reader, bb)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]byte, 32)
		encode(dst, bb)
		_ = string(dst)
	}
}

func BenchmarkOrigEncoder2(b *testing.B) {
	bb := make([]byte, 20)
	_, err := io.ReadFull(rand.Reader, bb)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := make([]byte, 32)
		valid.Encode40bytes2(dst, bb)
		_ = string(dst)
	}
}

func TestEncoding(t *testing.T) {
	var wordSafeBase32 = base32.NewEncoding("23456789CFGHJMPQRVWXcfghjmpqrvwx").WithPadding(base32.NoPadding)
	b := make([]byte, 400)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		t.Fatal(err)
	}
	id1 := wordSafeBase32.EncodeToString(b)
	dst := make([]byte, 640)
	valid.Encode40bytes2(dst, b)
	id2 := string(dst)

	if id1 != id2 {
		fmt.Println(id1)
		fmt.Println(id2)
		t.Fatal(id1, id2)
	}
}
