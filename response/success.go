package response

import (
	"net/http"
)

type Success struct {
	Success interface{} `json:"success"`
}

func ServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write([]byte(`{"error": "unexpected error"}`))
}

func ClientError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write([]byte(`{"error": "invalid request"}`))
}
