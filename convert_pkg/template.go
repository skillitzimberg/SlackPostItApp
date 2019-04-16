package main

import (
	//"encoding/json"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"log"
)

type Template struct {
}

type TemplateInput struct {
	Events map[string]interface{}
	SlackTemplate string
}

type TemplateOutput struct {
	TemplateJson string
}

func (t Template) Name() string {
	return "set_fields"
}

func (t Template) Version() string {
	return "1.0"
}

func (t Template) Execute(ctx step.Context) (interface{}, error) {
	input := TemplateInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		log.Fatal(err, "BindInputs: ")
		return nil, err
	}

	output, err := t.execute(input)
	if err != nil {
		log.Fatal(err, "Execute: ")
		return nil, err
	}

	return output, nil
}

func (t Template) execute(input TemplateInput) (
	*TemplateOutput, error) {
	var templateStr = `
	{{range input.Events}}
		{
			"type" : "mrkdwn",
			"text" : "*{{.Name}}*\n {{.Group.Name}}\n {{.Link}}\n {{.Date}}\n {{.Time}}\n *{{.Venue.Name}}*\n {{.Venue.Address}}\n {{.Venue.City}}, {{.Venue.State}} {{.Venue.Zip}}"
		},
	{{end}}
`
	return &TemplateOutput{TemplateJson: templateStr,}, nil
}

//type Response struct {
//	Events []Event
//}
//
//type Group struct {
//	Name string
//}
//
//type Event struct {
//	Name string
//	Group Group
//	Link string
//	Date string `json:"local_date"`
//	Time string `json:"local_time"`
//	Venue Venue
//}
//
//type Venue struct {
//	Name string
//	Address string `json:"address_1"`
//	City string
//	State string
//	Zip string
//}

/*
 events : [
    {"Name" : "Event 1",
        "Venue" : {
            "City" : "Portland"
        }
    },
    {"Name" : "Event 2"}
]

[
    {"title" : "Event 1", "fields" : [
        {"title" : "Location", "value" : "Portland", "short" : true}
    ]},,
    {"title" : "Event 2"}
]
 */


/*
   {{range .Events}}
		{"title" : {{.Name}},

          "fields" : [
				{{range .Attendees}}
					{ "title" : "Attendee", "text"  : {{.Name}}}
                {{end}
			]
		}
	{{end}}
 */
