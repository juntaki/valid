package valid

import "time"

var defaultSource = NewSource(14).WithChecksum().WithTimestamp()

func Generate() string {
	return defaultSource.Generate()
}

func Timestamp(id string) time.Time {
	return defaultSource.Timestamp(id)
}

func IsValid(id string) bool {
	return defaultSource.IsValid(id)
}
