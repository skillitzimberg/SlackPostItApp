package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/pkg/errors"
)

type FetchAndQueue struct {
	Fetcher
}

type FetchAndQueueInputs struct {
	FetchListInputs
	Workflow string
}

func (FetchAndQueue) Name() string {
	return "fetch_and_queue"
}

func (FetchAndQueue) Version() string {
	return "1.0"
}
func (fetch FetchAndQueue) Execute(in step.Context) (interface{}, error) {
	input := FetchAndQueueInputs{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return fetch.executeQueue(input, in)
}

func (fetch FetchAndQueue) executeQueue(input FetchAndQueueInputs, in step.Context) (interface{}, error) {
	output, err := fetch.execute(input.FetchListInputs)
	if err != nil {
		return nil, err
	}
	outputs, ok := output.(FetchListOutputs)
	if !ok {
		return nil, errors.New("invalid output from fetch list")
	}

	engine := in.Engine()
	for _, record := range outputs.Records {
		err := engine.AddToQueue(input.Workflow, record)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
