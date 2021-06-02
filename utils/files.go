package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func OpenFileRead(path string) (*os.File, error) {
	if !FileExists(path) {
		return nil, errors.New(path + " file dos not exist")
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func OpenFileAppend(path string) (*os.File, error) {
	if !FileExists(path) {
		return nil, errors.New(path + " file dos not exist")
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func CreateFile(path string) (*os.File, error) {
	if FileExists(path) {
		return nil, errors.New(path + " out file alridy exist")
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func ReadFileData(f *os.File) ([]byte, error) {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return data, err
	}
	return data, nil
}
