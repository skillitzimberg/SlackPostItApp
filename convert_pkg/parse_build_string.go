package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"strings"
)

type ParseBuildString struct {
}

func (ParseBuildString) Name() string {
	return "parse_build_string"
}

func (ParseBuildString) Version() string {
	return "1.0"
}

func (ParseBuildString) Execute(in step.Context) (interface{}, error) {
	input := buildStringInput{}

	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	components := strings.Split(input.String, input.Delimiter)
	var buildString strings.Builder
	for _, letter := range components {
		fmt.Fprintf(&buildString, "%v ", letter)
	}
	newString := buildString.String()

	return buildStringOutput{
		NewString: newString,
	}, nil
}

type buildStringInput struct {
	String       string
	Delimiter 	 string
}

type buildStringOutput struct {
	NewString string
}

// RUNNING & TESTING THE STEP:
// STEP_NAME=parse_build_string STEP_VERSION=1.0 go run convert_pkg/
//{
//	"String": "post+todo+say+hi",
//	"Delimiter": "+"
//}