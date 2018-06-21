package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Directory struct {
	modTime Time
	path    string
	name    string
}

func (d Directory) visit(jobs chan<- func(), operation int, ref Node) {

	if operation == REMOVE {
		d.remove()
		return
	}

	os.MkdirAll(ref.getPath(), os.ModePerm)

	//Write all existing files into in reference node into a map
	refMap := make(map[string]Node)
	files, err := ioutil.ReadDir(ref.getPath())
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		node := CreateNode(ref.getPath(), file)
		refMap[node.getName()] = node
	}

	//Read current files and iterate through them
	files, err = ioutil.ReadDir(d.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		node := CreateNode(d.path, file)

		if n, present := refMap[node.getName()]; present {
			node.visit(jobs, operation, n)
			delete(refMap, node.getName())

		} else {
			node.visit(jobs, operation,
				MakeEmptyNode(ref.getPath(), node.getName()))
		}
	}

	//Remove all files that still exist in refMap since they won't get replaced
	if operation == SYNC {
		for _, node := range refMap {
			node.visit(jobs, REMOVE, EmptyNode{})
		}
	}
}

func (d Directory) remove() {
	err := os.RemoveAll(d.path)
	if err != nil {
		log.Fatal(err)
	}
}

func (d Directory) getModificationTime() Time {
	return d.modTime
}

func (d Directory) getPath() string {
	return d.path
}

func (d Directory) getName() string {
	return d.name
}

func CreateNode(path string, info FileInfo) Node {
	if info.IsDir() {
		return MakeDirectory(path, info)
	} else {
		return MakeFile(path, info)
	}
}

func MakeDirectory(path string, info FileInfo) Directory {
	if !info.IsDir() {
		log.Fatal("Given file info doesn't belong to a directory")
	}

	return Directory{
		info.ModTime(),
		CreatePath(path, info.Name()),
		info.Name(),
	}
}

func MakeRootDirectory(dirPath string) Directory {
	return Directory{
		Time{},
		dirPath,
		path.Base(dirPath),
	}
}
