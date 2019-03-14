package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(FetchFile{})
	step.Register(ListDirectory{})
	step.Register(ReadLinesAndQueue{})

	step.Run()
}
