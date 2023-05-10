// Package miniurl is a url shortener
package miniurl

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash generates 32 byte long deterministic string
func Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
