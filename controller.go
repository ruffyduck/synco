package main

func run(operations []Operation) {
	jobs := make(chan func(), 20)

	for i := 0; i < 10; i++ {
		go worker(jobs)
	}

	for _, op := range operations {
		opType := GetOperationType(op.Operation)

		source := MakeRootDirectory(op.Source)
		target := MakeRootDirectory(op.Target)

		source.visit(jobs, opType, target)
	}
}

func worker(jobs <-chan func()) {
	for n := range jobs {
		n()
	}
}
