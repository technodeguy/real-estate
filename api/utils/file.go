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
