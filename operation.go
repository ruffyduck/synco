package main

import (
	"strings"
)

const (
	UPDATE = iota
	COPY
	SYNC
	REMOVE
)

type Operation struct {
	Source    string
	Target    string
	Operation string
}

func GetOperationType(name string) int {
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

	panic("Given name is not a valid operation")
}
