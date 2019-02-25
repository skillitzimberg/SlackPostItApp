package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"io/ioutil"
	"net/http"
	"strings"
)

type PostWebhookInput struct {
	Url         string
	Body        string
	ContentType string
}

type GetWebhookInput struct {
	Url string
}

type WebhookOutput struct {
	ResponseBody string
}

type PostWebhook struct {
}

func (PostWebhook) Name() string {
	return "webhook_post"
}

func (PostWebhook) Description() string {
	return "Posts a webhook"
}

func (PostWebhook) Version() string {
	return "1.0.0"
}

func (PostWebhook) Execute() {
	input := PostWebhookInput{}
	step.BindInputs(&input)
	reader := strings.NewReader(input.Body)
	handleResponse(http.Post(input.Url, input.ContentType, reader))
}

type GetWebhook struct {
}

func (GetWebhook) Name() string {
	return "webhook_get"
}

func (GetWebhook) Description() string {
	return "Performs a GET webhook"
}

func (GetWebhook) Version() string {
	return "1.0.0"
}

func (GetWebhook) Execute() {
	input := GetWebhookInput{}
	step.BindInputs(&input)
	handleResponse(http.Get(input.Url))
}

func handleResponse(resp *http.Response, err error) {
	if err != nil {
		step.ReportError(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			step.ReportError(err)
		}
		step.SetOutput(&WebhookOutput{ResponseBody: string(body)})
	} else {
		step.ReportError(fmt.Errorf("invalid response code %d", resp.StatusCode))
	}
}
