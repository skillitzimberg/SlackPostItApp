package main

import (
	"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type FetchSingle struct {
	Fetcher
}

type FetchSingleInputs struct {
	Username string
	Password string
	Url      string
	Endpoint string
	Select   string
	Id       int
	IdField  string
}

type FetchSingleOutput struct {
	Found  bool
	Record JsonMap
}

func (FetchSingle) Name() string {
	return "fetch_single"
}

func (FetchSingle) Version() string {
	return "1.0"
}

func (fetch FetchSingle) Execute(in step.Context) (interface{}, error) {
	input := FetchSingleInputs{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return fetch.execute(input)
}

func (fetch FetchSingle) ExecuteJson(js string) (interface{}, error) {
	input := FetchSingleInputs{}
	err := json.Unmarshal([]byte(js), &input)
	if err != nil {
		return nil, err
	}
	return fetch.execute(input)
}

func (fetch FetchSingle) execute(input FetchSingleInputs) (interface{}, error) {
	// set the username and pass on the fetcher
	fetch.username = input.Username
	fetch.password = input.Password
	fetch.url = input.Url
	// first thing we need to do is login to the FAMIS services
	// we require Username, Password, and Url so we do not have to validate the values
	var err error
	fetch.authToken, fetch.refreshToken, fetch.expiration, err = fetch.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}

	// add "Id" filter for fetch record
	// feilds defaults to Id if not overwritten
	field := "Id"
	if len(input.IdField) > 0 {
		field = input.IdField
	}
	return fetch.FetchSingle(input.Url, input.Endpoint, input.Select, field, input.Id)
}
