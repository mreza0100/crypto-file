package main

import (
	"log"
)

func main() {
	D, E, filePath, out, password := getFlags()

	if out == "" {
		out = renameInPath(filePath, func(fileName string) string {
			if E {
				return "encrypted_" + fileName
			}
			return "decrypt_" + fileName
		})
	}

	var err error
	if E {
		err = Encryption(filePath, out, password)
	} else if D {
		err = Decryption(filePath, out, password)
	}

	if err != nil {
		log.Fatal(err)
	}
}
