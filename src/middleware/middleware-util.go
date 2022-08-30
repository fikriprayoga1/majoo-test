package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"src/model"
)

func ErrorHandler(err error, w http.ResponseWriter, responseCode int, responseMessage string) {
	var modelResponseError model.ModelResponseError
	var responseJson []byte

	log.Printf("logInfo : %v => %v", responseMessage, err)
	modelResponseError = model.ModelResponseError{ResponseMessage: responseMessage}
	responseJson, err = json.Marshal(modelResponseError)
	if err != nil {
		log.Println(err)
	}
	log.Printf("logInfo : Proccess failed\n\n")

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(responseJson)

}
