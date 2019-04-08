package main

import (
	"github.com/apptreesoftware/go-workflow/pkg/step"
)

func main() {
	step.Register(Hello{}) // tell the workflow engine that this step exists
	step.Run()
}