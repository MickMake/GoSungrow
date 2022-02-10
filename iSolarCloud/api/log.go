package api

import "fmt"

type ErrorArg string

const (
	ErrorType ErrorArg = "error_type"
	Interface ErrorArg = "interface"
	Instance  ErrorArg = "instance"
)

type ErrorArgs map[ErrorArg]interface{}

func LogError(args ErrorArgs) {
	fmt.Println(args)
}

type errorType string

const (
	DoesNotImplement errorType = "does_not_implement"
)
