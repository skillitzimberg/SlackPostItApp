package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"strconv"
	"strings"
)

type TargetedParseLine struct {
}

func (TargetedParseLine) Name() string {
	return "targeted_line_parse"
}

func (TargetedParseLine) Version() string {
	return "1.0"
}

func (TargetedParseLine) Execute(in step.Context) (interface{}, error) {
	input := targetedSplitInput{}

	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	components := strings.SplitN(input.String, input.Delimiter, input.Indices)
	record := map[string]interface{}{}

	for k, v := range input.StringFields {
		if len(components) <= v {
			continue
		}
		val := components[v]
		record[k] = val
	}
	for k, v := range input.IntFields {
		if len(components) <= v {
			continue
		}
		val := components[v]
		if val == "" {
			continue
		}
		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse index %d (%s) into an int value", v, val)
		}
		record[k] = intVal
	}

	for k, v := range input.FloatFields {
		if len(components) <= v {
			continue
		}
		val := components[v]
		if val == "" {
			continue
		}
		floatVal, err := strconv.ParseFloat(val, 10)
		if err != nil {
			return nil, fmt.Errorf("unable to parse index %d (%s) into an int value", v, val)
		}
		record[k] = floatVal
	}

	return targetedSplitOutput{
		Record: record,
	}, nil
}

type targetedSplitInput struct {
	String       string
	Delimiter    string
	Indices		 int
	StringFields map[string]int
	IntFields    map[string]int
	FloatFields  map[string]int
}

type targetedSplitOutput struct {
	Record map[string]interface{}
}

// RUNNING & TESTING THE STEP:
// STEP_NAME=targeted_line_parse STEP_VERSION=1.0 go run convert_pkg/
//{
//	"String": "post+todo+say+hi",
//	"Delimiter": "+",
//	"Indices": 3,
//	"StringFields": {"action": 0, "category": 1, "unparsedNote": 2}
//}