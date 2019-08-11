package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/george-kj/krsc/domain"
)

// Response represent final response struct.
type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result,omitempty"`
}

// Fail with some error.
func fail(w http.ResponseWriter, errRes *domain.ResponseError) {

	// Log error details.
	errDetails, err := json.Marshal(errRes)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Println(string(errDetails))

	// Return only necessary information as response of the API.
	r := Response{
		Status: "nok",
		Result: struct {
			Message string
			Details []string
		}{errRes.Message, errRes.Details},
	}

	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8;")
	w.WriteHeader(errRes.Code)
	w.Write(j)
}

// Success with response body.
func succeed(w http.ResponseWriter, result interface{}) {
	r := Response{
		Status: "ok",
		Result: result,
	}

	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8;")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// ResponseHandler is Generic handler for writing response header and body all handler functions
func ResponseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, *domain.ResponseError)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Execute application handler
		res, err := h(w, r)
		if err != nil {
			fail(w, err)
			return
		}
		succeed(w, res)
	})
}
