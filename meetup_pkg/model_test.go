package main

import (
	"github.com/json-iterator/go"
	"testing"
)

var eventJson = `{"created":1543692689000,"duration":5400000,"id":"256869518","name":"VM (Vicky) Brasseur—   Open Source: What even is? How even to?","date_in_series_pattern":false,"status":"upcoming","time":1555547400000,"local_date":"2019-04-17","local_time":"17:30","updated":1551891806000,"utc_offset":-25200000,"waitlist_count":0,"yes_rsvp_count":32,"venue":{"id":24631440,"name":"Vacasa","lat":45.529876708984375,"lon":-122.68421936035156,"repinned":true,"address_1":"926 NW 13th Ave","city":"Portland","country":"us","localized_country_name":"USA","zip":"","state":"OR"},"group":{"created":1438483775000,"name":"Portland JR DEVELOPER Meetup!","id":18793056,"join_mode":"open","lat":45.56999969482422,"lon":-122.63999938964844,"urlname":"Portland-JR-DEVELOPER-Meetup","who":"Junior Developers","localized_location":"Portland, OR","state":"OR","country":"us","region":"en_US","timezone":"US/Pacific"},"link":"https://www.meetup.com/Portland-JR-DEVELOPER-Meetup/events/256869518/","description":"<p>VM (Vicky) Brasseur, Director of Open Source Strategy for Juniper Networks, joins us to talk about all things open source. She'll start with a presentation that will cover:</p> <p>* What _is_ open source? There's a Definition for that!<br/>* What are licenses and why are they important?<br/>* How can you find a project to contribute to?</p> <p>After the presentation she'll take any and all questions related to free and open source software.</p> <p>One lucky attendee will go home with a free copy of the first and only book about how to contribute to open source! <a href=\"https://fossforge.com\" class=\"linkified\">https://fossforge.com</a></p> <p>Bio: <a href=\"https://gist.github.com/vmbrasseur/d51a02363e78a54657204c258cbef29e\" class=\"linkified\">https://gist.github.com/vmbrasseur/d51a02363e78a54657204c258cbef29e</a></p> ","visibility":"public"}`

func TestParseEvent(t *testing.T)  {
	e := Event{}

	err := jsoniter.Unmarshal([]byte(eventJson), &e)
	if err != nil {
		t.Fail()
	}

	if e.Venue.Name == "Vacasa" {
		t.Fail()
	}
}
