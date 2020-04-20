package errors

type Op string

type Error struct {
	Op      Op
	Err     error
	Payload map[string]interface{}
}

func (e *Error) Error() string {
	return string(e.Op) + " " + e.Err.Error()
}

func (e *Error) With(payload map[string]interface{}) *Error {
	for k, v := range payload {
		e.Payload[k] = v
	}

	return e
}

func E(op Op, err error) *Error {
	return &Error{
		Op:  op,
		Err: err,
	}
}
