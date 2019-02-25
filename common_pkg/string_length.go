package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type StringLengthCounter struct {
}

func (StringLengthCounter) Name() string {
	return "string_length"
}

func (StringLengthCounter) Version() string {
	return "1.0"
}

func (StringLengthCounter) Execute() {
	input := StringLengthInput{}
	step.BindInputs(&input)
	output := StringLengthOutput{}
	output.Count = len(input.Text)
	step.SetOutput(output)
}

type StringLengthInput struct {
	Text string
}

type StringLengthOutput struct {
	Count int
}
