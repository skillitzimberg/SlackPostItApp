package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"reflect"
)

type JsonDiff struct {
}

type JsonDiffInput struct {
	Left            map[string]interface{}
	Right           map[string]interface{}
	FieldsToCompare []string
}

type JsonDiffOutput struct {
	Different          bool
	FieldsThatDiffered []string
}

func (JsonDiff) Name() string {
	return "json_diff"
}

func (JsonDiff) Version() string {
	return "1.0"
}

// This step takes two json objects `Left` and `Right` and a list of
// fields or `FieldsToCompare` you wish this step to compare
//
// if `FieldsToCompare` is not provided this step will compare every single
// field in the left object and compare it to the right
//
func (diff JsonDiff) Execute(in step.Context) (interface{}, error) {
	jsonDiffIn := &JsonDiffInput{}
	err := in.BindInputs(jsonDiffIn)
	if err != nil {
		return nil, err
	}

	left := jsonDiffIn.Left
	right := jsonDiffIn.Right
	// did we get an fields to compare?
	if fields := jsonDiffIn.FieldsToCompare; fields != nil && len(fields) > 0 {
		return diff.diffFields(fields, left, right), nil
	} else {
		// we will always get the fields to check from the Left map
		mapFields := diff.getStringKeysFromMap(left)
		return diff.diffFields(mapFields, left, right), nil
	}
}

// I know this is causing me to make two passes over that map keys
// I am ok with that because of the simplicity on `diffing` the fields
func (diff JsonDiff) getStringKeysFromMap(data map[string]interface{}) []string {
	if data == nil {
		return make([]string, 0)
	}
	keys := reflect.ValueOf(data).MapKeys()
	if len(keys) < 1 {
		return make([]string, 0)
	}
	result := make([]string, 0, len(keys))
	// iterate keys and put strings into result
	for idx, key := range keys {
		result[idx] = key.String()
	}
	return result
}

func (diff JsonDiff) diffFields(fields []string, left map[string]interface{}, right map[string]interface{}) JsonDiffOutput {
	// are the two objs different
	isDifferent := false
	// a collection of the fields that are different
	differentFields := make([]string, 0)
	// iterate  given fields and check if the there are differences
	for _, field := range fields {
		// the fields values differ
		if diff.fieldsDiffer(field, left, right) {
			isDifferent = true
			differentFields = append(differentFields, field)
		}
	}
	return JsonDiffOutput{Different: isDifferent, FieldsThatDiffered: differentFields}

}

func (diff JsonDiff) fieldsDiffer(field string, left map[string]interface{}, right map[string]interface{}) bool {
	leftData := left[field]
	rightData := right[field]
	return leftData == rightData
}
