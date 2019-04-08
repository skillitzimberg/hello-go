package main

import (
	"fmt"
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

type Greeting interface {
	PrintGreeting() string

}

type Hello struct {

}


func (speak Hello) Name() string {
	return  "hello"
}

func (speak Hello) Version() string {
	return  "1.0"
}


func (speak Hello) Execute(in step.Context) (interface{}, error) {
	input := HelloInput{}
	step.BindInputs(&input)
	output := HelloOutput{}

	output.Success = true
	step.SetOutput(output)
}

func (Hello) PrintGreeting() string {
	fmt.Println("Hello, Scott")
	return "Hello, Scott"
}

type HelloInput struct {
	Text string
}

type HelloOutput struct {
	Success bool
}