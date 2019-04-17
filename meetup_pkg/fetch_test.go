package main

import "testing"

func TestFetcMeetup_Execute(t *testing.T) {
	input := FetchMeetupInput{
		TopicCategory: "techcodedevelopercodinghack",
		EndDate: "2019-04-13T22:00:00",
		Radius: "10",
		Lat: "45.531577",
		Lon: "-122.680746",
		Page: "10",
	}

	out, err := FetchMeetup{}.execute(input)
	if err != nil {
		t.Fail()
	}

	if out.Events[0].Venue.Name == "Vacasa" {
		t.Fail()
	}
}