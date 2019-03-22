package main

import (
	"context"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func main() {
	step.Register(ReadSheet{})
	step.Register(WriteToSheet{})
	step.Run()
}

func GetGoogleSheet(input ReadSheetInput) (*spreadsheet.Sheet, error) {
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
	return sheet, nil
}
