package main

import (
	"crypto_file/utils"
	"os"
)

func Encryption(filePath, out, password string) (err error) {
	var (
		origionF, destinationF *os.File
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
		defer destinationF.Close()
	}

	{
		dataHash, err := utils.GetFileHash(filePath)
		if err != nil {
			return err
		}
		_, err = destinationF.Write(dataHash)
		if err != nil {
			return err
		}
	}

	{
		err = utils.XORFile(origionF, destinationF, password)
		if err != nil {
			return err
		}
	}

	return nil
}
