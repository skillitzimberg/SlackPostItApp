package main

import (
	"context"
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type WriteToSheet struct{}

func (WriteToSheet) Name() string {
	return "write"
}

func (WriteToSheet) Version() string {
	return "1.0"
}

func (w WriteToSheet) Execute(ctx step.Context) (interface{}, error) {
	input := WriteToSheetInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	err = w.execute(input)
	return nil, err
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
	newRow := data.RowCount
	for k, v := range fields {
		cellVal := input.Record[v]
		sheet.Update(int(newRow), k, fmt.Sprintf("%v", cellVal))
	}

	err = sheet.Synchronize()
	return err
}

type WriteToSheetInput struct {
	SpreadsheetId string
	SheetIndex    uint
	Credentials   string
	Fields        []string
	Record        map[string]interface{}
}
