package rest

import "fmt"

type ErrHttp struct {
	StatusCode int
	Status     string
}

func (err ErrHttp) Error() string {
	return fmt.Sprintf("HTTP error: status=%d, reason=%s", err.StatusCode, err.Status)
}
