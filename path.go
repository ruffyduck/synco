package main

import (
	"strings"
)

//CreatePath Create a new path from given path and file name
func CreatePath(path string, fileName string) string {
	if strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\") {
		return path + fileName
	}

	return path + "/" + fileName
}
