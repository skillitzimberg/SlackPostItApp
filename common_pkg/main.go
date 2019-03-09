package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(PostWebhook{})
	step.Register(GetWebhook{})
	step.Register(Filter{})
	step.Register(StringLengthCounter{})
<<<<<<< HEAD
	step.Register(SliceString{})
=======
	step.Register(FetchFile{})
	step.Register(ParseJsonObject{})
>>>>>>> 6d441fbc29150306581342754dd905bb97313ca5

	step.Run()
}
