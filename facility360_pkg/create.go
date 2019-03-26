package main

import (
	"bytes"
	"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

const createRequestMethod = "POST"

type CreateRecord struct {
	Fetch
}

func (CreateRecord) Name() string {
	return "create_record"
}

func (CreateRecord) Version() string {
	return "1.0"
}

func (create CreateRecord) ExecuteJson(jsonString string) (interface{}, error) {
	input := &Facility360CreateIn{}
	err := json.Unmarshal([]byte(jsonString), input)
	if err != nil {
		return nil, err
	}
	return create.execute(input)
}

func (create CreateRecord) Execute(in step.Context) (interface{}, error) {
	input := &Facility360CreateIn{}
	err := in.BindInputs(input)
	if err != nil {
		return nil, err
	}
	return create.execute(input)
}

func (create CreateRecord) execute(input *Facility360CreateIn) (interface{}, error) {
	// get authenticated
	err := create.LogMeInFacility360(input.Facility360Input)
	if err != nil {
		return nil, err
	}

	// we have our create url
	createUrl, err := create.getUrl(input.Url, input.Endpoint)
	if err != nil {
		return nil, err
	}

	// get record bytes
	data, err := create.getRecordData(input)
	if err != nil {
		return nil, err
	}

	// build http request
	req, err := create.buildRequest(createRequestMethod, createUrl.String(), bytes.NewReader(data))

	if err != nil {
		return nil, err
	}
	// send request
	resp, err := create.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return create.handleFailedResponse(resp)
	}
	defer resp.Body.Close()
	return create.handleUpsertResponse(resp)
}

func (create CreateRecord) getRecordData(input *Facility360CreateIn) ([]byte, error) {
	return json.Marshal(input.Record)
}
