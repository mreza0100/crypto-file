package utils

import (
	"io"
	"os"
)

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

func trim(tempData []byte) []byte {
	newData := make([]byte, 0, len(tempData))

	for i := 0; i < len(tempData); i++ {
		v := tempData[i]

		if v != 0 {
			newData = append(newData, v)
		}

	}

	return newData
}

func XORFile(origionF, destinationF *os.File, password string) error {
	const readSize = 1024 * 1024 * 10

	data := make([]byte, readSize)
	for done := false; ; {
		{
			n, err := origionF.Read(data)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			data = data[:n]
		}
		{
			_, err := destinationF.Write(XORData(data, password))
			if err != nil {
				return err
			}
		}

		if done {
			break
		}
	}

	return nil
}
