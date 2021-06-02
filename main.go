package main

import (
	"crypto_file/utils"
	"fmt"
	"os"
)

func main() {
	D, E, filePath, out, password := getFlags()

	if out == "" {
		out = utils.RenameInPath(filePath, func(fileName string) string {
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
		removeErr := os.Remove(out)
		if removeErr != nil {
			fmt.Println("cant remove generated file")
		}
		panic(err)
	}

	fmt.Println("Done :D")

	// err := utils.XORFile("./data/text.txt", "./data/e", "gogoly")
	// if err != nil {
	// 	panic(err)
	// }

	// f, _ := utils.OpenFileRead("./data/e")
	// data, _ := ioutil.ReadAll(f)
	// fmt.Println(string(utils.XORData(data, "gogoly")))
}
