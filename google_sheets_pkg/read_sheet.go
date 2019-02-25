package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type ReadSheet struct {
}

func (ReadSheet) Name() string {
	return "read"
}

func (ReadSheet) Version() string {
	return "1.0"
}

func (s ReadSheet) Execute() {
	input := ReadSheetInput{}
	step.BindInputs(&input)
	out, err := s.execute(input)
	step.ReportError(err)
	step.SetOutput(out)
}

func (ReadSheet) execute(input ReadSheetInput) (*ReadSheetOutput, error) {
	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), spreadsheet.Scope)
	if err != nil {
		return nil, err
	}
	client := conf.Client(context.Background())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(input.SpreadsheetId)
	if err != nil {
		return nil, err
	}
	sheet, err := spreadsheet.SheetByIndex(input.SheetIndex)
	if err != nil {
		return nil, err
	}

	var rows []map[string]interface{}
	for _, row := range sheet.Rows {
		rowMap := map[string]interface{}{}
		for colIndex, name := range input.Fields {
			rowMap[name] = row[colIndex].Value
		}
		rows = append(rows, rowMap)
	}
	return &ReadSheetOutput{
		Rows: rows,
	}, nil
}

type ReadSheetInput struct {
	SpreadsheetId string
	SheetIndex    uint
	Credentials   string
	Fields        []string
}

type ReadSheetOutput struct {
	Rows []map[string]interface{}
}
