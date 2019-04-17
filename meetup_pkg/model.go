package main


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