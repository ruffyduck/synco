package main

import (
	"os"
	"time"
)

//FileInfo os.FileInfo alias
type FileInfo = os.FileInfo

//Time time.Time alias
type Time = time.Time

type node interface {
	visit(jobs chan<- func(), operation int, reference node)
	remove()

	getModificationTime() Time
	getPath() string
	getName() string
}

type emptyNode struct {
	path string
	name string
}

func (n emptyNode) visit(jobs chan<- func(), operation int, reference node) {}

func (n emptyNode) remove() {}

func (n emptyNode) getModificationTime() Time {
	return Time{}
}

func (n emptyNode) getPath() string {
	return n.path
}

func (n emptyNode) getName() string {
	return n.name
}

func makeEmptyNode(path string, name string) emptyNode {
	return emptyNode{
		CreatePath(path, name),
		name,
	}
}
