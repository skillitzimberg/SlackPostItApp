package main

import (
	"bufio"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
	"strings"
)

type ReadLinesAndQueue struct{}

func (ReadLinesAndQueue) Name() string {
	return "read_lines_and_queue"
}

func (ReadLinesAndQueue) Version() string {
	return "1.0"
}

func (f ReadLinesAndQueue) Execute(ctx step.Context) (interface{}, error) {
	input := ReadLinesAndQueueInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	err = f.execute(input)
	return nil, err
}

func (ReadLinesAndQueue) execute(input ReadLinesAndQueueInput) error {
	file, err := os.Open(input.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	engine := step.GetEngine()

	lines := make([]string, 0)
	for scanner.Scan() {
		lineRecord := scanner.Text()
		lines = append(lines, lineRecord)
	}

	records := make([]map[string]string, 0)
	for i, line := range lines {
		fieldSlice := strings.Split(line, input.FieldDelimiter)
		if i == 0 && input.UseHeaderAsFieldNames {
			input.FieldNames = fieldSlice
		} else {
			recordToQueue := ConvertToMap(fieldSlice, input.FieldNames)
			records = append(records, recordToQueue)
			err := engine.AddToQueue(input.Workflow, recordToQueue)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ReadLinesAndQueueInput struct {
	FilePath              string
	FieldNames            []string
	UseHeaderAsFieldNames bool
	FieldDelimiter        string
	Workflow              string
}
