package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	//"github.com/mongodb/mongo-go-driver/bson"
)

type CachePush struct {
}

func (CachePush) Name() string {
	return "push"
}

func (CachePush) Version() string {
	return "1.0"
}

func (CachePush) Execute(in step.Context) (interface{}, error) {
	input := pushInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	engine := in.Engine()
	err = engine.PutRecord(input.Id, input.Record, map[string]interface{}{}, input.CacheName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type pushInput struct {
	Record    map[string]interface{}
	Id        string
	CacheName string
}


type CachePull struct {
}

func (CachePull) Name() string {
	return "pull"
}

func (CachePull) Version() string {
	return "1.0"
}

func (CachePull) Execute(in step.Context) (interface{}, error) {
	input := pullInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	engine := in.Engine()
	rec := map[string]interface{}{}
	found, _, err := engine.PullRecord(input.Id, &rec, input.CacheName)
	fmt.Println(found)
	if err != nil {
		return nil, err
	}
	if !found {
		return pullOutput{Found: false}, nil
	}
	return pullOutput{
		Record: rec,
		Found:  true,
	}, nil
}

type pullInput struct {
	Id        string
	CacheName string
}

type pullOutput struct {
	Record map[string]interface{}
	Found  bool
}


type CachePullBulk struct {
}

func (CachePullBulk) Name() string {
	return "pull_bulk"
}

func (CachePullBulk) Version() string {
	return "1.0"
}

func (CachePullBulk) Execute(in step.Context) (interface{}, error) {
	input := pullBulkInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	engine := in.Engine()
	//records := make([]Records, 0)
	filter := &MyFilter{category: input.Category}
	found, err := engine.Find(filter, input.CacheName, input.Limit)
	if err != nil {
		return nil, err
	}

	if len(found) == 0 {
		return pullBulkOutput{Found: false}, nil
	}

	fmt.Printf("found: %v\n", found)
	bsonString := fmt.Sprintf("%b", found)
	fmt.Printf("found: %v\n", bsonString)

	//err = bson.Unmarshal(found[0].Record, &records)
	//if err != nil {
	//	return nil, err
	//}
	//
	//fmt.Printf("found: %v\n", records)

	return pullBulkOutput{
		BsonString: bsonString,
	}, nil
}

type pullBulkInput struct {
	Category  string
	CacheName string
	Limit     int64
}

type pullBulkOutput struct {
	BsonString string
	Found  bool
}

type MyFilter struct {
	category string `bson:"Category" json:"Category"`
}

type Record struct {
	Id string
	Category string
	Note string
}

type Records struct {
	Records []Record
}