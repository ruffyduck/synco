package main

func run(operations []operation) {
	jobs := make(chan func(), 20)

	for i := 0; i < 10; i++ {
		go worker(jobs)
	}

	for _, op := range operations {
		opType := getOperationType(op.Operation)

		source := makeRootDirectory(op.Source)
		target := makeRootDirectory(op.Target)

		source.visit(jobs, opType, target)
	}
}

func worker(jobs <-chan func()) {
	for n := range jobs {
		n()
	}
}
