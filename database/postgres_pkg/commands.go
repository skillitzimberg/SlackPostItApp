package main

import (
	"database/sql"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	_ "github.com/lib/pq"
	"util/database/db_common"
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
	db, err := sql.Open("postgres", input.ConnectionString)
	step.ReportError(err)
	db_common.PerformQuery(db, input)
}

type QueryAndQueue struct {
}


func (QueryAndQueue) Name() string {
	return "query_and_queue"
}

func (QueryAndQueue) Version() string {
	return "1.0"
}

func (QueryAndQueue) Execute() {
	input := db_common.DatabaseCommandToQueue{}
	step.BindInputs(&input)
	db, err := sql.Open("postgres", input.ConnectionString)
	step.ReportError(err)
	db_common.PerformQueryAndQueue(db, input)
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
	db, err := sql.Open("postgres", input.ConnectionString)
	step.ReportError(err)
	err = db_common.PerformInsertAll(db, input)
	step.ReportError(err)
}
