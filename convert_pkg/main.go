package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(ParseJsonObject{})
	step.Register(MapRecords{})
	step.Register(SetFields{})
	step.Register(ExtractFields{})

	step.Run()
}
