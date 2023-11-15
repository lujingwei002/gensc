package gensc

import "errors"

var (
	ErrGenNotChange   = errors.New("doc file not changed")
	ErrTODO           = errors.New("TODO")
	ErrDbNotSupport   = errors.New("db not support")
	ErrNullPointer    = errors.New("null pointer")
	ErrDataNotExist   = errors.New("data not exist")
	ErrDataDeleteFail = errors.New("data delete fail")
	ErrDataNotFound   = errors.New("data not found")
)
