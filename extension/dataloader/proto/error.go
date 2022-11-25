package proto

import "errors"

var (
	ErrUnsupportedMethod = errors.New("unsupported method called")
	ErrFileCannotLoad    = errors.New("file cannot load")
)
