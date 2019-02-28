package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type SliceString struct {
}

func (SliceString) Name() string {
	return "slice"
}

func (SliceString) Version() string {
	return "1.0"
}

func (SliceString) Execute() {
	input := SliceInput{}
	step.BindInputs(&input)
	output := SliceOutput{}

	if len(input.Text) < input.EndIndex {
		panic("Out of bounds! The end index must be less than the text's" +
			" length")
	}

	if input.StartIndex < 0 {
		panic("Out of bounds! The start index must be 0 or greater.")
	}

	message := input.Text[input.StartIndex:input.EndIndex]

	output.Text = message
	step.SetOutput(output)
}

type SliceInput struct {
	Text string
	StartIndex int
	EndIndex int
}

type SliceOutput struct {
	Text string
}
