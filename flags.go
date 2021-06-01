package main

import (
	"flag"
	"log"
)

func getFlags() (D, E bool, filePath, out, password string) {
	var (
		DPtr        = flag.Bool("D", false, "decription")
		EPtr        = flag.Bool("E", false, "encryption")
		passwordPtr = flag.String("password", "", "your secret key")
		filePathPtr = flag.String("filePath", "", "path to your file")
		outPtr      = flag.String("out", "", "path to your out file")
	)

	{
		flag.Parse()

		D, E = *DPtr, *EPtr
		filePath, password, out = *filePathPtr, *passwordPtr, *outPtr
	}

	{
		if (!E && !D) || (E && D) {
			log.Fatal("encryption or decription?", " set a flag for god sake")
		}
		if password == "" {
			log.Fatal("no password?", " how i'am gonna do with this shit!?")
		}

		if filePath == "" {
			log.Fatal("no filePath?", " what i'am gonna do with this shit!?")
		}
	}

	return D, E, filePath, out, password
}
