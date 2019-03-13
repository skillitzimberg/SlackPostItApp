package main

import (
	"bufio"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
	"strings"
)

//read_lines_and_queue
//oracle/postgres insert_one


type FetchFileInput struct {
	FilePath              string
	FieldNames            []string
	UseHeaderAsFieldNames bool
	FieldDelimiter        string
}

type FetchFileOutput struct {
	Records []map[string]string
}

type FetchFile struct {
}

func (FetchFile) Name() string {
	return "file_read_lines"
}

func (FetchFile) Description() string {
	return "Parses a file into a list of records"
}

func (FetchFile) Version() string {
	return "1.0"
}

func (f FetchFile) Execute() {
	input := FetchFileInput{}
	step.BindInputs(&input)
	out, err := f.execute(input)
	step.ReportError(err)
	step.SetOutput(out)
}

func (FetchFile) execute(input FetchFileInput) (*FetchFileOutput, error) {
	file, err := os.Open(input.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	records := make([]map[string]string, 0)
	for i, line := range lines {
		fieldSlice := strings.Split(line, input.FieldDelimiter)
		if i == 0 && input.UseHeaderAsFieldNames {
			input.FieldNames = fieldSlice
		} else {
			records = append(records, ConvertToMap(fieldSlice, input.FieldNames))
		}
	}
	return &FetchFileOutput{Records: records}, nil
}

func ConvertToMap(fields, fieldNames []string) map[string]string {
	fieldMap := make(map[string]string)
	for i, fieldName := range fieldNames {
		if i >= len(fields) {
			break
		}
		fieldMap[fieldName] = fields[i]
	}
	return fieldMap
}
