package main

import (
	"bufio"
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
)

type ReadLinesAndQueue struct{}

func (ReadLinesAndQueue) Name() string {
	return "read_lines_and_queue"
}

func (ReadLinesAndQueue) Version() string {
	return "1.0"
}

func (f ReadLinesAndQueue) Execute(ctx step.Context) (interface{}, error) {
	input := ReadLinesAndQueueInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	err = f.execute(input, ctx)
	return nil, err
}

func (ReadLinesAndQueue) execute(input ReadLinesAndQueueInput, ctx step.Context) error {
	file, err := os.Open(input.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	engine := ctx.Engine()

	for scanner.Scan() {
		lineRecord := scanner.Text()
		err := engine.AddToQueue(input.Workflow, lineRecord)
		if err != nil {
			return err
		}
		println("Queuing line", lineRecord)
	}
	return nil
}

type ReadLinesAndQueueInput struct {
	FilePath string
	Workflow string
}
