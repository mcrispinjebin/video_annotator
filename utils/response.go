package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type errResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func ReturnResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	en := json.NewEncoder(w)
	_ = en.Encode(data)
}

func ErrorResponse(w http.ResponseWriter, responseErrorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	var buf = new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	_ = encoder.Encode(errResponse{Message: responseErrorMessage})
	w.WriteHeader(statusCode)
	_, _ = w.Write(buf.Bytes())
}
