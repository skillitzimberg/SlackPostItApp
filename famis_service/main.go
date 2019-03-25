package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

func main() {
	step.Register(Fetch{})
	step.Register(FetchSingle{})
	step.Register(FetchAndQueue{})
	step.Run()
}
