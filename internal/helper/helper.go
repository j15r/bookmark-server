package helper

import (
	"fmt"
	"net/http"
)

func SendError(w http.ResponseWriter, message string, code int) {
	// enableCors(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func SendResponse(w http.ResponseWriter, json []byte, code int) {
	// enableCors(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
