package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/robertkrimen/otto"
)

type ScriptInput struct {
	Script string
	ScriptVars map[string]interface{}
}

type ScriptOutput struct {
	ReturnVal interface{}
}

type JavascriptRunner struct {

}

func (JavascriptRunner) Name() string {
	return "js"
}

func (JavascriptRunner) Version() string {
	return "1.0"
}

func (JavascriptRunner) Execute(ctx step.Context) (interface{}, error) {
	input := ScriptInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}

	vm := otto.New()

	for key, value := range input.ScriptVars {
		err = vm.Set(key, value)
		if err != nil {
			return nil, err
		}
	}

	val, err := vm.Run(input.Script)
	if err != nil {
		return nil, err
	}
	var returnVal interface{}
	if val.IsDefined() {
		if val.IsBoolean() {
			returnVal, err = val.ToBoolean()
		} else if val.IsNumber() {
			returnVal, err = val.ToInteger()
			if err != nil {
				returnVal, err = val.ToFloat()
			}
		} else if val.IsString() {
			returnVal, err = val.ToString()
		} else if val.IsObject() {
			returnVal = val.Object()
		}
	}
	return ScriptOutput{ReturnVal: returnVal}, err
}


