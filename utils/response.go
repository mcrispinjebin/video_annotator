package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"video_annotator/constants"
	"video_annotator/models"
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

func ErrorResponse(w http.ResponseWriter, customErr *models.CustomErr) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := customErr.StatusCode
	if statusCode == 0 {
		statusCode = constants.HttpInternalServerError
	}
	var buf = new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	log.Printf("%s - %q", customErr.Message, customErr.Err)
	_ = encoder.Encode(errResponse{Message: customErr.Message})
	w.WriteHeader(statusCode)
	_, _ = w.Write(buf.Bytes())
}
