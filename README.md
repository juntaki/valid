# Validatable Unique Identifier

[![Go Reference](https://pkg.go.dev/badge/github.com/juntaki/valid.svg)](https://pkg.go.dev/github.com/juntaki/valid)
[![GoCover](http://gocover.io/_badge/github.com/juntaki/valid)](http://gocover.io/github.com/juntaki/valid)

valid is a library for generating random ID, with the default configuration as follows

- 5 bytes timestamp
- 14 bytes of random bytes
- 1 byte checksum

Features:
- Cryptographically secure (use of crypto/rand)
- Lexicographically sortable
- Allows validation of ID
- Word safe and URL safe characters
- Embedded time with millisecond precision

In addition, most of the features can be changed by configuration. 

| Name        | Secure Random Bit Size| String ID Size    | Sortable | Checksum
|-------------|---|----------------|-----------|-------
| [google/uuid]|  122 bits | 36 chars       | no        | no
| [oklog/ulid] |  80 bits (user defined source) | 26 chars      | yes | no
| [rs/xid]     |  n/a (not cryptographically secure) | 20 chars       | yes| no
| valid (default) |  112 bits| 32 chars | yes | yes
| valid (configurable) |  any | depends on config | yes / no | yes / no 

[google/uuid]: https://github.com/google/uuid
[oklog/ulid]:https://github.com/oklog/ulid
[rs/xid]: https://github.com/rs/xid

## Install

    go get github.com/juntaki/valid

## Usage

```go
id := valid.Generate()

fmt.Println(id, len(id), valid.IsValid(id), valid.Timestamp(id))
// Output: 2X9P75FX2pJqCcHRVWV2862JW6XhFr6x 32 true 2021-09-08 23:59:17.339250688 +0900 JST
```
