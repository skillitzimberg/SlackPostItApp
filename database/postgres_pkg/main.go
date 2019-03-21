package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(Query{})
	step.Register(QueryAndQueue{})
	step.Register(InsertBatch{})
	step.Register(Execute{})
	step.Run()
}
