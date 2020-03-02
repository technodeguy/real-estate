package utils

import (
	"mime"
	"strings"
)

func GetFileMimeType(filename string) string {
	slice := strings.Split(filename, ".")
	dotAndExt := "." + slice[len(slice)-1]
	return mime.TypeByExtension(dotAndExt)
}

func IsAcceptedMimeType(allowedFormats []string, mt string) bool {
	for _, val := range allowedFormats {
		if val == mt {
			return true
		}
	}

	return false
}
