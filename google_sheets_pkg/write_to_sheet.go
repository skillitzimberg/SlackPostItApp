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

	// Row names defauls to Record name is user doesn't pass in fields
	var fields []string
	if input.Fields != nil && len(input.Fields) != 0 {
		fields = input.Fields
	} else {
		for k := range input.Record {
			fields = append(fields, k)
		}
	}
	data := sheet.Properties.GridProperties
	newRow := data.RowCount + 1
	for k, v := range fields {
		cellVal := input.Record[v]
		sheet.Update(int(newRow),k,cellVal)
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
