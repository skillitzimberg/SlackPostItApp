package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type WriteToSheet struct {}

func (WriteToSheet) Name() string {
	return "write"
}

func (WriteToSheet) Version() string {
	return "1.0"
}

func (w WriteToSheet) Execute() {
	input := WriteToSheetInput {}
	step.BindInputs(&input)
	err := w.execute(input)
	step.ReportError(err)
}

func (WriteToSheet) execute(input WriteToSheetInput) error {
	conf, err := google.JWTConfigFromJSON([]byte(input.Credentials), spreadsheet.Scope)
	if err != nil {
		return err
	}
	client := conf.Client(context.Background())

	service := spreadsheet.NewServiceWithClient(client)
	googleSheet, err := service.FetchSpreadsheet(input.SpreadsheetId)
	if err != nil {
		return err
	}
	sheet, err := googleSheet.SheetByIndex(input.SheetIndex)
	if err != nil {
		return err
	}

	for k, v := range input.Fields {
		cellVal := input.Record[v]
		sheet.Update(0, k, cellVal)
	}
	err = sheet.Synchronize()

	return err
}

type WriteToSheetInput struct {
	SpreadsheetId string
	SheetIndex    uint
	Credentials   string
	Fields        []string
	Record map[string]string
}
