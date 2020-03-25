package client

import "fmt"

var (
	ErrNotFound     = fmt.Errorf("flag not found")
	ErrUnauthorized = fmt.Errorf("unauthorized")
)

type ServerError struct {
	code     int
	response string
}

func (err *ServerError) Error() string {
	return fmt.Sprintf("unexpected status code %v: %v", err.code, err.response)
}
