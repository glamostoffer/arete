package errlist

import (
	"errors"
	"fmt"
	"net/http"
)

type customErr struct {
	message string
	code    int
}

func (e *customErr) Error() string {
	return fmt.Sprintf("STATUS_CODE: %d; ERROR: %s", e.code, e.message)
}

func (e *customErr) Code() int {
	return e.code
}

func New(code int, message string) error {
	return &customErr{
		message: message,
		code:    code,
	}
}

func GetErrStatus(err error) int {
	var cerr *customErr
	if errors.As(err, &cerr) {
		return cerr.Code()
	}
	return http.StatusInternalServerError
}
