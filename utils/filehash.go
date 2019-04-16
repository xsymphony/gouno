package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func CalculateHash(r io.Reader) (string, error) {
	h := sha256.New()
	_, err := io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(h.Sum(nil)), nil
}
