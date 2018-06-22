package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file := flag.String("json", "", "Specify json file that can define "+
		"multiple source folders, target folders and operations\n"+
		`Format: [{"source":"...", "target":"...", "operation":"..."}, ...]`)
	source := flag.String("source", "", "Synchronisation source folder")
	target := flag.String("target", "", "Synchronisation target folder")
	opType := flag.String("operation", "Update", "Specify operation type =>\n"+
		"  Update: Only updates existing files (that have been modified),\n"+
		"    Copy: Copies all files from source to target,\n"+
		"    Sync: Like 'Copy', also deletes all outdated files in target")

	flag.Parse()

	operations := []operation{}
	if len(*file) != 0 {
		data, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(data, &operations); err != nil {
			log.Fatal(err)
		}

	}

	if len(*source) != 0 && len(*target) != 0 &&
		!strings.EqualFold(*source, *target) {

		operations = append(operations, operation{*source, *target, *opType})
	}

	if len(operations) == 0 {
		fmt.Println("No valid arguments given, try -help")
	}

	run(operations)
}
