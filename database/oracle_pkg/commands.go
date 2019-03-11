package main

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"github.com/apptreesoftware/step_library_go/database/db_common"
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
	input := db_common.DatabaseCommand{}
	step.BindInputs(&input)
	db, err := sql.Open("goracle", input.ConnectionString)
	step.ReportError(err)
	db_common.PerformQuery(db, input)
}

type InsertBatch struct {
}

func (InsertBatch) Name() string {
	return "insert_batch"
}

func (InsertBatch) Version() string {
	return "1.0"
}

func (InsertBatch) Execute() {
	input := &db_common.InsertCommand{}
	step.BindInputs(input)
	db, err := sql.Open("goracle", input.ConnectionString)
	step.ReportError(err)
	err = db_common.PerformInsertAll(db, input)
	step.ReportError(err)
}
