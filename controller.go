package main

import ()

func run(operations []Operation) {

	for _, op := range operations {
		opType := GetOperationType(op.Operation)

		source := MakeRootDirectory(op.Source)
		target := MakeRootDirectory(op.Target)

		source.visit(opType, target)
	}

}
