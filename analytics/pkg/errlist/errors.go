package errlist

import (
	"net/http"
)

var (
	ErrInvalidAffectedRows = New(http.StatusInternalServerError, "INVALID_AFFECTED_ROWS_COUNT")
)
