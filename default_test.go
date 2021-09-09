package valid_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/juntaki/valid"
)

func TestIsValid(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				id: "22MjwpfpMGhr6fh4xwCxJw3h9QHHhWMx",
			},
			want: true,
		},
		{
			name: "1 letter change",
			args: args{
				id: "22MjwpfpMGhr6fh4xwCxJw3h9QHHhWMa",
			},
			want: false,
		},
		{
			name: "long",
			args: args{
				id: "22MjwpfpMGhr6fh4xwCxJw3h9QHHhWMxa",
			},
			want: false,
		},
		{
			name: "short",
			args: args{
				id: "22MjwpfpMGhr6fh4xwCxJw3h9QHHhWM",
			},
			want: false,
		},
		{
			name: "bad letter",
			args: args{
				id: "22MjwpfpMGhr6fh4xwCxJw3h9QHHhWM=",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.IsValid(tt.args.id); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	for i := 1; i < 100; i++ {
		id := valid.Generate()
		if time.Since(valid.Timestamp(id)).Milliseconds() > 1 {
			t.Fatal("invalid")
		}
		if !valid.IsValid(id) {
			t.Fatal("invalid")
		}
		time.Sleep(time.Millisecond)
	}
}

func TestTimestamp(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "valid",
			args: args{
				id: "22Mm4V5HwqC5WxVGPQMPcmPrC9pv33xh",
			},
			want: 1631198571713134592,
		},
		{
			name: "short (valid timestamp and invalid id)",
			args: args{
				id: "22Mm4V5HwqC5WxVG",
			},
			want: time.Time{}.UnixNano(),
		},
		{
			name: "invalid",
			args: args{
				id: "aaaaaa",
			},
			want: time.Time{}.UnixNano(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Timestamp(tt.args.id).UnixNano(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestampWithReferenceTime(t *testing.T) {
	type args struct {
		id string
	}
	valid.SetReferenceTime(time.Date(2020, 9, 4, 0, 0, 0, 0, time.UTC))
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "valid",
			args: args{
				id: "22Mm4V5HwqC5WxVGPQMPcmPrC9pv33xh",
			},
			want: 1599662571713134592,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Timestamp(tt.args.id).UnixNano(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
