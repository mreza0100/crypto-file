package main

import (
	"crypto/sha512"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func openFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func createFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func readFileData(f *os.File) ([]byte, error) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return data, err
	}
	return data, nil
}
func writeFileData(path string, data []byte) error {
	f, err := createFile(path)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func renameInPath(path string, changer func(string) string) string {
	splitedFilePath := strings.Split(path, "/")

	fileName := splitedFilePath[len(splitedFilePath)-1]

	splitedFilePath[len(splitedFilePath)-1] = changer(fileName)

	return strings.Join(splitedFilePath, "/")
}

func XORData(data []byte, password string) []byte {
	newData := make([]byte, len(data))
	copy(newData, data)

	for i := 0; i < len(newData); i++ {
		for j := 0; j < len(password); j++ {
			newData[i] = newData[i] ^ password[j]
		}
	}

	return newData
}

func getHash(data []byte) []byte {
	hashedData := sha512.Sum512(data)
	return hashedData[:]
}

func separateHashAndData(encryptedData []byte) (data, hash []byte, err error) {
	if len(encryptedData) < 65 {
		return data, hash, errors.New("wrong file this is not a encrypted file")
	}

	data = encryptedData[:len(encryptedData)-64]
	hash = encryptedData[len(encryptedData)-64:]

	return
}

func compareHashs(h1, h2 []byte) bool {
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
