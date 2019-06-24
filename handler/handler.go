package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func sendJSON(payload interface{}, w http.ResponseWriter) {

	jstring, err := json.Marshal(payload)
	if err != nil {
		log.Printf(
			"error marshalling payload to json with error: %s",
			err.Error(),
		)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(``))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jstring)
}

func sendEmptyResponse(statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write([]byte(``))
}
