package utils

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"os"
)

func GetFileHash(path string) ([]byte, error) {
	f, err := OpenFileRead(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("%x", h.Sum(nil))), nil
}

func ExtractHashFile(f *os.File) (hash []byte, err error) {
	hash = make([]byte, 128)
	n, err := f.Read(hash)
	if err != nil {
		return
	}

	if n < 128 {
		err = errors.New("wrong file this is not a encrypted file too small")
		return
	}

	return
}
func CompareHashs(h1, h2 []byte) bool {
	if len(h1) != len(h2) {
		return false
	}

	for i := 0; i < len(h1); i++ {
		if h1[i] != h2[i] {
			return false
		}
	}

	return true
}
