package main

import (
	"crypto/sha256"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type HashString struct {

}

func (HashString) Name() string {
	return "hash_string"
}

func (HashString) Version() string {
	return "1.0"
}

func (HashString) Execute(in step.Context) (interface{}, error) {
	input := hashInput{}
	err := in.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	noteToHash := input.Note

	h := sha256.New()
	h.Write([]byte(noteToHash))
	hash := h.Sum(nil)
	hashToString := string(hash)

	hashedString := hashToString
	record := map[string]interface{}{"id": hashedString, "category": input.Category, "note": noteToHash}

	return hashOutput{
		Record: record,
		Hash: hashedString,
	}, nil
}

type hashInput struct {
	Note string
	Category string
}

type hashOutput struct {
	Record  interface{}
	Hash    string
}

// RUNNING & TESTING THE STEP:
// STEP_NAME=hash_string STEP_VERSION=1.0 go run convert_pkg/
//{
//	"String": "banana"
//}