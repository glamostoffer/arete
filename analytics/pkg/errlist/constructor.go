package errlist

import "fmt"

type customErr struct {
	message string
	code    int
}

func (e *customErr) Error() string {
	return fmt.Sprintf("STATUS_CODE: %d; ERROR: %s", e.code, e.message)
}

func New(code int, message string) error {
	return &customErr{
		message: message,
		code:    code,
	}
}
