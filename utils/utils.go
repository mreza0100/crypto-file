package utils

import (
	"strings"
)

func RenameInPath(path string, changer func(string) string) string {
	splitedFilePath := strings.Split(path, "/")

	fileName := splitedFilePath[len(splitedFilePath)-1]

	splitedFilePath[len(splitedFilePath)-1] = changer(fileName)

	return strings.Join(splitedFilePath, "/")
}
