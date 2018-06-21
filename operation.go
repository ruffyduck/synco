package main

import (
	"strings"
)

const (
	UPDATE = iota
	MOVE
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

	if strings.EqualFold(n, "MOVE") {
		return MOVE
	}

	if strings.EqualFold(n, "SYNC") {
		return SYNC

	}
	if strings.EqualFold(n, "REMOVE") {
		return REMOVE
	}

	panic("Given name is not a valid operation")
}
