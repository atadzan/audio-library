package repository

import "fmt"

const (
	errNotFound           = "NOT_FOUND"
	errOperationFailedMsg = "OPERATION_FAILED"
)

var (
	ErrNotFound        = fmt.Errorf(errNotFound)
	ErrOperationFailed = fmt.Errorf(errOperationFailedMsg)
)
