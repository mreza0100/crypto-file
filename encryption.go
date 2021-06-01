package main

import (
	"bytes"
)

func Encryption(filePath, out, password string) error {
	var (
		rawData       []byte
		enctyptedData []byte
	)

	{
		f, err := openFile(filePath)
		if err != nil {
			return err
		}

		rawData, err = readFileData(f)
		if err != nil {
			return err
		}

		enctyptedData = XORData(rawData, password)
	}

	{
		dataHash := getHash(rawData)
		enctyptedData = bytes.Join([][]byte{enctyptedData, dataHash}, []byte{})
	}

	{
		writeFileData(out, enctyptedData)
	}

	return nil
}
