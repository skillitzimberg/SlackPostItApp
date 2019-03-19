package main

import "github.com/apptreesoftware/go-workflow/pkg/step"

type MapRecords struct {

}

func (MapRecords) Name() string {
	return "map_records"
}

func (MapRecords) Version() string {
	return "1.0"
}

func (MapRecords) Execute(ctx step.Context) (interface{}, error) {
	inputs := MapRecordsInputs{}
	err := ctx.BindInputs(&inputs)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{}, 0)
	from := inputs.From
	for key, val := range inputs.MapValues {
		value := from[key]
		result[val] = value
	}
	return MapRecordsOutputs{To:result}, nil
}

type MapRecordsInputs struct {
	From map[string]interface{}
	MapValues map[string]string
}

type MapRecordsOutputs struct {
	To map[string]interface{}
}

