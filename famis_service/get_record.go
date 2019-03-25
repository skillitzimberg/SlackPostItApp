package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"net/http"
	"net/http/httputil"
)

type FetchSingle struct {
	Fetch
}

func (FetchSingle) Name() string {
	return "get_record"
}

func (FetchSingle) Version() string {
	return "1.0"
}

func (fetch FetchSingle) Execute(in step.Context) (interface{}, error) {
	input := FetchInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return fetch.execute(input)
}

func (fetch FetchSingle) execute(input FetchInput) (interface{}, error) {
	// set the username and pass on the fetcher
	fetch.username = input.Username
	fetch.password = input.Password
	fetch.url = input.Url
	input.Top = 1
	// first thing we need to do is login to the FAMIS services
	// we require Username, Password, and Url so we do not have to validate the values
	var err error
	fetch.authItem, err = fetch.Login(input.Username, input.Password, input.Url)
	if err != nil {
		return nil, err
	}

	uri, err := fetch.buildUrl(input)
	if err != nil {
		return nil, err
	}

	req, err := fetch.buildRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := fetch.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	b, err := httputil.DumpResponse(resp, true)
	if err == nil {
		println(string(b))
	}
	jsonResp, err := handleFetchResponse(resp)
	if err != nil {
		return nil, err
	}

	return handleSingleResponse(jsonResp)
}
