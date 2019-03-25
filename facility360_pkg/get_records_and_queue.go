package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type FetchAndQueue struct {
	Fetch
}

func (FetchAndQueue) Name() string {
	return "get_records_and_queue"
}

func (FetchAndQueue) Version() string {
	return "1.0"
}
func (fetch FetchAndQueue) Execute(in step.Context) (interface{}, error) {
	input := FetchAndQueueInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return fetch.executeQueue(input, in)
}

func (fetch FetchAndQueue) executeQueue(input FetchAndQueueInput, in step.Context) (interface{}, error) {
	// set the username and pass on the fetcher
	fetch.username = input.Username
	fetch.password = input.Password
	fetch.url = input.Url
	// first thing we need to do is login to the FAMIS services
	// we require Username, Password, and Url so we do not have to validate the values
	var err error
	fetch.authItem, err = fetch.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}

	uri, err := fetch.buildUrl(input.FetchInput)
	if err != nil {
		return nil, err
	}

	engine := in.Engine()
	err = fetch.performPagedFetch(uri, input.Endpoint, func(messages []jsoniter.RawMessage) {
		for _, msg := range messages {
			err := engine.AddToQueue(input.Workflow, msg)
			if err != nil {
				log.Fatalf("Unable to add item to engine queue: %s", err.Error())
			}
		}
	})
	return nil, err
}
