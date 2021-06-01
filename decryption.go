package main

import (
	"log"
)

func Decryption(filePath, out, password string) error {
	var (
		enctyptedData []byte
		decryptedData []byte

		originalHash []byte
	)

	{
		f, err := openFile(filePath)
		if err != nil {
			return err
		}

		enctyptedData, err = readFileData(f)
		if err != nil {
			return err
		}
	}

	{
		var err error
		enctyptedData, originalHash, err = separateHashAndData(enctyptedData)
		if err != nil {
			return err
		}
	}

	{
		decryptedData = XORData(enctyptedData, password)
	}

	{
		newHash := getHash(decryptedData)
		if !compareHashs(originalHash, newHash) {
			log.Fatal("password is wrong or file has been damaged")
		}
	}

	{
		writeFileData(out, decryptedData)
	}

	return nil
}
