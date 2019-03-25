package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

func main() {
	step.Register(Fetch{})
	step.Register(GetRecord{})
	step.Register(FetchAndQueue{})
	step.Run()
}
