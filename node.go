package main

import (
	"os"
	"time"
)

type FileInfo = os.FileInfo
type Time = time.Time

type Node interface {
	visit(jobs chan<- func(), operation int, reference Node)
	remove()

	getModificationTime() Time
	getPath() string
	getName() string
}

type EmptyNode struct {
	path string
	name string
}

func (n EmptyNode) visit(jobs chan<- func(), operation int, reference Node) {}

func (n EmptyNode) remove() {}

func (n EmptyNode) getModificationTime() Time {
	return time.Time{}
}

func (n EmptyNode) getPath() string {
	return n.path
}

func (n EmptyNode) getName() string {
	return n.name
}

func MakeEmptyNode(path string, name string) EmptyNode {
	return EmptyNode{
		CreatePath(path, name),
		name,
	}
}
