package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file := flag.String("json", "", "Define multiple source folders, target"+
		"folders and operations in single json file")
	source := flag.String("source", "", "Synchronisation source folder")
	target := flag.String("target", "", "Synchronisation target folder")
	operation := flag.String("operation", "Update", "Operation type:"+
		"Update, Move, Synchronize, Remove")

	flag.Parse()

	operations := []Operation{}
	if len(*file) != 0 {
		data, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal(err)
		}

		code := string(data)
		if err := json.Unmarshal([]byte(code), &operations); err != nil {
			log.Fatal(err)
		}

	} else if len(*source) != 0 && len(*target) != 0 &&
		!strings.EqualFold(*source, *target) {

		operations = append(operations, Operation{*source, *target, *operation})
	}

	run(operations)
}
