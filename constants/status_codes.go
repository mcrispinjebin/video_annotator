package constants

import "net/http"

const (
	HttpStatusBadRequest = http.StatusBadRequest
	HttpUnauthorised     = http.StatusUnauthorized
	HttpResourceNotFound = http.StatusNotFound
	HttpResourceExists   = http.StatusConflict

	HttpInternalServerError = http.StatusInternalServerError

	HttpStatusOK        = http.StatusOK
	HttpStatusNoContent = http.StatusAccepted
)
