package utils

import (
	"io"
	"os"
	"path"
)

// IsFileExists判断所给路径文件是否存在
func IsFileExists(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func CopyAndSeek0(f *os.File, w io.Reader) error {
	_, err := io.Copy(f, w)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	return nil
}

func JoinCacheFile(dir, filename string) (string, error) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path.Join(dir, filename), nil
}
