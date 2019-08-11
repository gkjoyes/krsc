package handler

import (
	"net/http"

	"github.com/george-kj/krsc/domain"
)

// Status give application status.
func Status(w http.ResponseWriter, r *http.Request) (interface{}, *domain.ResponseError) {
	return struct {
		Response string `json:"status"`
	}{"I am alive:)"}, nil
}
