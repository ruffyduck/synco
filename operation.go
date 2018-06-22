package main

import (
	"strings"
)

//Enum of possible file operations
const (
	UPDATE = iota
	COPY
	SYNC
	REMOVE
)

type operation struct {
	Source    string
	Target    string
	Operation string
}

func getOperationType(name string) int {
	n := strings.ToUpper(name)

	if strings.EqualFold(n, "UPDATE") {
		return UPDATE
	}

	if strings.EqualFold(n, "COPY") {
		return COPY
	}

	if strings.EqualFold(n, "SYNC") {
		return SYNC

	}

	panic(name + " is not a valid operation")
}
