package http

import "fmt"

type HandlerError struct {
	Previous error
	Message  string
}

func (e HandlerError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.Previous)
}
