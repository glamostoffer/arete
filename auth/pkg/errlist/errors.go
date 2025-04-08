package errlist

import (
	"net/http"
)

var (
	ErrInvalidAffectedRows     = New(http.StatusInternalServerError, "INVALID_AFFECTED_ROWS_COUNT")
	ErrInvalidConfirmationCode = New(http.StatusBadRequest, "INVALID_CONFIRMATION_CODE")
	ErrResendCooldown          = New(http.StatusBadRequest, "RESEND_COOLDOWN_DIDNT_PASSED_YET")
	ErrUserAlreadyExists       = New(http.StatusBadRequest, "USER_WITH_GIVEN_LOGIN_OR_PASSWORD_ALREADY_EXISTS")
)
