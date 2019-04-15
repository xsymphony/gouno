package utils

import (
	"encoding/base64"
	"crypto/sha256"
	"io"
)

func CalculateHash(r io.Reader) string {
	h := sha256.New()
	io.Copy(h, r)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
