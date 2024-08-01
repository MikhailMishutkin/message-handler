package httphandler

import "net/http"

func FailOnErrorsHttp(w http.ResponseWriter, err error, message string, status int) {
	if err != nil {
		w.WriteHeader(status)
		w.Write([]byte(message + err.Error()))
	}
}

type errorResponse struct {
	Message string `json:"message"`
}
