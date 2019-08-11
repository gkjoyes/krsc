package backend

import (
	"net/http"

	"github.com/george-kj/krsc/domain"
)

func newInternalServerError(err error, details ...string) *domain.ResponseError {
	return &domain.ResponseError{
		Code:       http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Details:    details,
		StackTrace: err.Error(),
	}
}
