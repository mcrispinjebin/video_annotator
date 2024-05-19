package utils

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetURLParam(r *http.Request, paramName string) (string, error) {
	params := mux.Vars(r)

	param, ok := params[paramName]
	if !ok {
		return "", fmt.Errorf("%s %q", "URL param not found", paramName)
	}

	return param, nil
}
