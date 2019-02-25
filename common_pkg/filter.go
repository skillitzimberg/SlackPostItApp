package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/robertkrimen/otto"
)

type Filter struct {
}

func (Filter) Name() string {
	return "filter"
}

func (Filter) Version() string {
	return "1.0"
}

func (Filter) Execute() {

	input := FilterInput{}
	step.BindInputs(&input)

	script := fmt.Sprintf("out = records.filter(function(record) { return %s });", input.Filter)
	vm := otto.New()

	err := vm.Set("records", input.Records)
	step.ReportError(err)
	_, err = vm.Run(script)
	step.ReportError(err)

	outArr, err := vm.Get("out")
	step.ReportError(err)

	output := FilterOutput{}
	exportedRecords, err := outArr.Export()
	step.ReportError(err)

	if outputRecords, ok := exportedRecords.([]map[string]interface{}); ok {
		output.Records = outputRecords
	}
	step.SetOutput(output)
}

type FilterInput struct {
	Records []map[string]interface{}
	Filter  string
}

type FilterOutput struct {
	Records []map[string]interface{}
}
