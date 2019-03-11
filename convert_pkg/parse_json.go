package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/json-iterator/go"
)

type ParseJsonObject struct {
}

func (ParseJsonObject) Name() string {
	return "parse_json_object"
}

func (ParseJsonObject) Version() string {
	return "1.0"
}

func (ParseJsonObject) Execute() {
	input := parseJsonInput{}
	step.BindInputs(&input)

	rec := map[string]interface{}{}
	err := jsoniter.UnmarshalFromString(input.String, &rec)
	if err != nil {
		step.ReportError(err)
		return
	}
	step.SetOutput(&parseJsonOutput{Record: rec})
}

type parseJsonInput struct {
	String string
}

type parseJsonOutput struct {
	Record map[string]interface{}
}
