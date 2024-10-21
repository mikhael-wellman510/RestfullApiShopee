package helper

import (
	"encoding/json"
	"fmt"

	"net/http"
)

// Baca Request
func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	if request.Body == nil || request.ContentLength == 0 {
		fmt.Println("Request body is empty")

	}
	// Kirim data ke object Result
	err := decoder.Decode(result)
	PanicIfErr(err)
}

// Kirim Response

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	// Kirim data ke Response
	err := encoder.Encode(response)
	PanicIfErr(err)
}
