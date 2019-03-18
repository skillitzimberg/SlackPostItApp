package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
	"os"
)

type ListDirectory struct{}

func (ListDirectory) Name() string {
	return "list_directory_contents"
}

func (ListDirectory) Version() string {
	return "1.0"
}

func (f ListDirectory) Execute(ctx step.Context) (interface{}, error) {
	input := ListDirectoryInput{}
	err := ctx.BindInputs(&input)
	if err != nil {
		return nil, err
	}
	return f.execute(input)
}

func (ListDirectory) execute(input ListDirectoryInput) (*ListDirectoryOutput, error) {
	dir, err := os.Open(input.DirectoryPath)
	if err != nil {
		return nil, err
	}
	files, err := dir.Readdir(-1)
	dir.Close()
	if err != nil {
		return nil, err
	}

	var output []string
	for _, file := range files {
		output = append(output, file.Name())
	}

	return &ListDirectoryOutput{
		Files: output,
	}, nil
}

type ListDirectoryInput struct {
	DirectoryPath string
}

type ListDirectoryOutput struct {
	Files []string
}
