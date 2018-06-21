package main

import (
	"strings"
)

func CreatePath(path string, fileName string) string {
	if strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\") {
		return path + fileName
	}

	return path + "/" + fileName
}
