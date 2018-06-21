package main

import (
	"log"
	"os"
)

type File struct {
	modTime Time
	path    string
	name    string
	size    int64
}

func (f File) visit(operation int, ref Node) {
	switch operation {
	case UPDATE:
	case MOVE:
	case SYNC:
		t := f.modTime
		b := (operation != UPDATE) ||
			(t.After(ref.getModificationTime()))

		if b {
			path := f.getPath()
			refPath := ref.getPath()

			//Move file to reference file location
			err := os.Rename(path, refPath)
			if err != nil {
				log.Fatal(err)
			}
		}
	case REMOVE:
		f.remove()
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
