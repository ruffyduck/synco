package main

import (
	"io"
	"log"
	"os"
	"path"
)

type file struct {
	modTime Time
	path    string
}

func (f file) visit(jobs chan<- func(), operation int, ref node) {
	if operation == REMOVE {
		jobs <- func() { f.remove() }
		return
	}

	t := f.modTime
	b := (operation != UPDATE) ||
		(t.After(ref.getModificationTime()))

	if b {
		currPath := f.getPath()
		refPath := ref.getPath()

		jobs <- func() {
			source, err := os.Open(currPath)
			if err != nil {
				log.Fatal(err)
			}

			defer source.Close()

			target, err := os.Create(refPath)
			if err != nil {
				log.Fatal(err)
			}

			defer target.Close()

			//Move file to reference file location
			//err := os.Rename(currPath, refPath)
			_, err = io.Copy(target, source)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (f file) remove() {
	path := f.getPath()

	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func (f file) getModificationTime() Time {
	return f.modTime
}

func (f file) getPath() string {
	return f.path
}

func (f file) getName() string {
	return path.Base(f.path)
}

func makeFile(path string, info FileInfo) file {
	if info.IsDir() {
		log.Fatal("Given file info doesn't belongs to a file")
	}

	return file{
		info.ModTime(),
		CreatePath(path, info.Name()),
	}
}
