package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

type directory struct {
	modTime Time
	path    string
}

func (d directory) visit(jobs chan<- func(), operation int, ref node) {

	if operation == REMOVE {
		d.remove()
		return
	}

	os.MkdirAll(ref.getPath(), os.ModePerm)

	//Write all existing files in reference node into a map
	refMap := make(map[string]node)
	files, err := ioutil.ReadDir(ref.getPath())
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		node := createNode(ref.getPath(), file)
		refMap[node.getName()] = node
	}

	//Read current files and iterate through them
	files, err = ioutil.ReadDir(d.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		node := createNode(d.path, file)

		if n, present := refMap[node.getName()]; present {
			node.visit(jobs, operation, n)
			delete(refMap, node.getName())

		} else {
			node.visit(jobs, operation,
				makeEmptyNode(ref.getPath(), node.getName()))
		}
	}

	//Remove all files that still exist in refMap since they won't get replaced
	if operation == SYNC {
		for _, node := range refMap {
			node.visit(jobs, REMOVE, emptyNode{})
		}
	}
}

func (d directory) remove() {
	err := os.RemoveAll(d.path)
	if err != nil {
		log.Fatal(err)
	}
}

func (d directory) getModificationTime() Time {
	return d.modTime
}

func (d directory) getPath() string {
	return d.path
}

func (d directory) getName() string {
	return path.Base(d.path)
}

func createNode(path string, info FileInfo) node {
	if info.IsDir() {
		return makeDirectory(path, info)
	}

	return makeFile(path, info)
}

func makeDirectory(path string, info FileInfo) directory {
	if !info.IsDir() {
		log.Fatal("Given file info doesn't belong to a directory")
	}

	return directory{
		info.ModTime(),
		CreatePath(path, info.Name()),
	}
}

func makeRootDirectory(dirPath string) directory {
	return directory{
		Time{},
		dirPath,
	}
}
