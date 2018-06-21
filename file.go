package main

import (
	"io"
	"log"
	"os"
)

type File struct {
	modTime Time
	path    string
	name    string
	size    int64
}

func (f File) visit(jobs chan<- func(), operation int, ref Node) {
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

func (f File) remove() {
	path := f.getPath()

	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func (f File) getModificationTime() Time {
	return f.modTime
}

func (f File) getPath() string {
	return f.path
}

func (f File) getName() string {
	return f.name
}

func MakeFile(path string, info FileInfo) File {
	if info.IsDir() {
		log.Fatal("Given file info doesn't belongs to a file")
	}

	return File{
		info.ModTime(),
		CreatePath(path, info.Name()),
		info.Name(),
		info.Size(),
	}
}
