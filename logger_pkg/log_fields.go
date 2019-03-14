package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type LogFields struct {
}

type LogFieldsInput struct {
	Record     map[string]interface{}
	FieldNames []string
}

func (LogFields) Name() string {
	return "log_fields"
}

func (LogFields) Version() string {
	return "1.0"
}

func (s LogFields) Execute() {
	input := LogFieldsInput{}
	step.BindInputs(&input)
	for _, key := range input.FieldNames {
		val := input.Record[key]
		if val == nil {
			fmt.Printf("%s : null\n", key)
		} else {
			fmt.Printf("%s : %v\n", key, val)
		}
	}
}
