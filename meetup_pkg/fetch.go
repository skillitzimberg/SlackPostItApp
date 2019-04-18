package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/joho/godotenv"
	"github.com/json-iterator/go"
	"log"
	"net/http"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}


type FetchMeetup struct {
}

func (f FetchMeetup) Name() string {
	return "fetchMeetup"
}

func (f FetchMeetup) Version() string {
	return "1.0"
}

type FetchMeetupInput struct {
	ApiKey string
	TopicCategory string
	EndDate string
	Radius string
	Lat string
	Lon string
	Page string
}

type FetchMeetupOutput struct {
	Events []Event
}

func (f FetchMeetup) Execute(in step.Context) (interface{}, error) {

	input := FetchMeetupInput{}
	err := in.BindInputs(&input)
	if err != nil {
		log.Fatal(err, "BindInputs: ")
		return nil, err
	}

	output, err := f.execute(input)
	if err != nil {
		log.Fatal(err, "Execute: ")
		return nil, err
	}

	fmt.Println(output)
	return output, nil
}

func (f FetchMeetup) execute(input FetchMeetupInput) (
	*FetchMeetupOutput, error) {
		fmt.Println(input.ApiKey)
	requestUrl := fmt.Sprintf("https://api.meetup." +
		"com/find/upcoming_events?&key=%s&lon=%s" +
		"&end_date=%s&topic_category=%s&page=%s&radius=%s&lat=%s",
		input.ApiKey, input.Lon, input.EndDate, input.TopicCategory,
		input.Page, input.Radius, input.Lat)
	response, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal(err, "Request: ")
		return nil, err
	}
	defer response.Body.Close()

	dec := jsoniter.NewDecoder(response.Body)
	resp := Response{}

	err = dec.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &FetchMeetupOutput{
		Events: resp.Events,
	}, nil
}