package main_test

import (
	cf "crypto_file"
	"os"
	"testing"
)

func createFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}

const (
	testText    = "this is a test text"
	password    = "my_password"
	targetPath  = "./data/text.txt"
	out         = "./data/enctypted"
	finalResult = "./data/decrypted_text"
)

func init() {
	f, err := createFile(targetPath)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(testText))
	if err != nil {
		panic(err)
	}
}

func Test_encryption(t *testing.T) {
	err := cf.Encryption(targetPath, out, password)
	if err != nil {
		t.Error(err)
	}
}

func Test_decryption(t *testing.T) {
	err := cf.Decryption(out, finalResult, password)
	if err != nil {
		t.Error(err)
	}
}
