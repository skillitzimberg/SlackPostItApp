package db_common

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func PerformQuery(db *sql.DB, command DatabaseCommand) {
	rows, err := db.Query(command.Sql)
	step.ReportError(err)
	cols, err := rows.Columns()
	step.ReportError(err)

	defer rows.Close()
	results := []map[string]interface{}{}
	for rows.Next() {
		rowMap, err := ScanIntoMap(rows, cols)
		step.ReportError(err)
		results = append(results, rowMap)
	}
	output := &RowOutput{
		Results: results,
	}
	step.SetOutput(output)
}

func PerformQueryAndQueue(db *sql.DB, command DatabaseCommandToQueue) {
	rows, err := db.Query(command.Sql)
	step.ReportError(err)
	cols, err := rows.Columns()
	step.ReportError(err)

	engine := step.GetEngine()

	defer rows.Close()
	for rows.Next() {
		rowMap, err := ScanIntoMap(rows, cols)
		step.ReportError(err)
		err = engine.AddToQueue(command.Workflow, rowMap)
		step.ReportError(err)
	}
}

func PerformInsertAll(db *sql.DB, command *InsertCommand) error {
	if len(command.Records) == 0 {
		return nil
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, rec := range command.Records {
		var rowValues []interface{}
		for _, fieldName := range command.ValueFields {
			value := rec[fieldName]
			rowValues = append(rowValues, value)
		}
		_, err := tx.Exec(command.Sql, rowValues...)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
