package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(PostWebhook{})
	step.Register(GetWebhook{})
	step.Register(Filter{})
	step.Register(StringLengthCounter{})

	step.Run()
}
