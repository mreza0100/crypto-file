package main

import (
	"crypto_file/utils"
	"fmt"
	"os"
)

func Decryption(filePath, out, password string) (err error) {
	var (
		origionF, destinationF *os.File
		originalHash           []byte
	)

	{
		origionF, err = utils.OpenFileRead(filePath)
		if err != nil {
			return err
		}
		defer origionF.Close()

		destinationF, err = utils.CreateFile(out)
		if err != nil {
			return err
		}
	}

	{
		originalHash, err = utils.ExtractHashFile(origionF)
		if err != nil {
			return err
		}
	}

	{
		err = utils.XORFile(origionF, destinationF, password)
		if err != nil {
			return
		}
	}

	{
		newHash, err := utils.GetFileHash(out)
		if err != nil {
			return err
		}

		if !utils.CompareHashs(originalHash, newHash) {
			fmt.Println(string(originalHash))
			fmt.Println(string(newHash))
		}
	}

	return nil
}
