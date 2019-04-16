package main

import (
	"os"
	"testing"
	"text/template"
)

func TestResponse(t *testing.T)  {

	testGroup := Group{
		Name: "Portland JR DEVELOPER Meetup!",
	}

	testVenue := Venue{
		Name: "Vacasa",
		Address: "926 NW 13th Ave",
		City: "Portland",
		State: "OR",
		Zip: "",
	}

	testEvent := Event{
		Name: "VM (Vicky) Brasseur—   Open Source: What even is? How even to?",
		Group: testGroup,
		Link: "https://www.meetup.com/Portland-JR-DEVELOPER-Meetup/events/256869518/",
		Date: "2019-04-17",
		Time: "17:30",
		Venue: testVenue,
	}

	var testEvents []Event
	testEvents = append(testEvents, testEvent)


	testVal := Response{
		Events: testEvents,
	}

	tmpl := template.New("test")
	tmpl, err := tmpl.Parse(templateStr)
	if err != nil {
		t.Error(err)
		return
	}
	err = tmpl.Execute(os.Stdout, testVal)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestExecute(t *testing.T) {
	// input is map[string]interface{}

	//output is
}

type Response struct {
	Events []Event
}

type Group struct {
	Name string
}

type Event struct {
	Name string
	Group Group
	Link string
	Date string `json:"local_date"`
	Time string `json:"local_time"`
	Venue Venue
}

type Venue struct {
	Name string
	Address string `json:"address_1"`
	City string
	State string
	Zip string
}

var templateStr = `
	{{range .Events}}
		{
			"type" : "mrkdwn",
			"text" : "*{{.Name}}*\n {{.Group.Name}}\n {{.Link}}\n {{.Date}}\n {{.Time}}\n *{{.Venue.Name}}*\n {{.Venue.Address}}\n {{.Venue.City}}, {{.Venue.State}} {{.Venue.Zip}}"
		},
	{{end}}
`