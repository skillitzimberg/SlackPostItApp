package main

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	_ "gopkg.in/goracle.v2"
)

type Query struct {
}

func (Query) Name() string {
	return "query"
}

func (Query) Version() string {
	return "1.0"
}

func (Query) Execute() {
	input := SelectInput{}
	step.BindInputs(&input)
	db, err := sql.Open("goracle", input.ConnectionString)
	step.ReportError(err)
	rows, err := db.Query(input.Query)
	step.ReportError(err)
	cols, err := rows.Columns()
	step.ReportError(err)

	defer rows.Close()
	var results []map[string]interface{}
	for rows.Next() {
		rowMap, err := scanIntoMap(rows, cols)
		step.ReportError(err)
		results = append(results, rowMap)
	}
	output := &RowOutput{
		Results: results,
	}
	step.SetOutput(output)
}

type SelectInput struct {
	ConnectionString string
	Query            string
}

type RowOutput struct {
	Results []map[string]interface{}
}

type Requestor struct {
	Id             string
	LastName       string
	FirstName      string
	MiddleInitial  string
	Department     string
	DepartmentCode string
}

func scanIntoMap(rows *sql.Rows, cols []string) (map[string]interface{}, error) {
	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i, _ := range columns {
		columnPointers[i] = &columns[i]
	}

	// Scan the result into the column pointers...
	if err := rows.Scan(columnPointers...); err != nil {
		return nil, err
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	m := make(map[string]interface{})
	for i, colName := range cols {
		val := columnPointers[i].(*interface{})
		m[colName] = *val
	}
	return m, nil
}
