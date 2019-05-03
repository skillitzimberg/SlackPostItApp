package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/mongodb/mongo-go-driver/bson"
)

type ParseBsonObject struct {
}

func (ParseBsonObject) Name() string {
	return "parse_bson_object"
}

func (ParseBsonObject) Version() string {
	return "1.0"
}

func (ParseBsonObject) Execute(ctx step.Context) (interface{}, error) {
	input := parseBsonInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	record := map[string]interface{}{}
	err = bson.Unmarshal(input.Bson, &record)
	if err != nil {
		return nil, err
	}
	return parseBsonOutput{Record: record}, nil
}

type parseBsonInput struct {
	Bson []byte
}

type parseBsonOutput struct {
	Record map[string]interface{}
}


// RUNNING & TESTING THE STEP:
// STEP_NAME=parse_bson_object STEP_VERSION=1.0 go run convert_pkg/

//{
//	"record":"n\000\000\000\002id\0009\000\000\000\357\277\275?l8\357\277\275\357\277\275O\030\313\273\0210\357\277\275\024\357\277\275\357\277\275Y\357\277\275\357\277\275\020@s\357\277\275l\032y\357\277\275mm\357\277\275G\357\277\275\000\002category\000\n\000\000\000groceries\000\002note\000\006\000\000\000kale \000\000",
//		"metadata":"\005\000\000\000\000"
//}

// PUBLISHING THE STEP
// make publish-convert_pkg