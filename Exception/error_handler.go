package exception

import (
	helper "golang-shopeeMart-api/Helper"
	"golang-shopeeMart-api/Model/web/responseDto"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Dipanggil Dari MAIN
func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundErr(writer, request, err) {
		return
	}

	if validationErr(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)

}

func notFoundErr(writer http.ResponseWriter, request *http.Request, err interface{}) bool {

	exception, ok := err.(NotFoundError)

	if ok {

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := responseDto.CommonResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func validationErr(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	// Taro nama error nya apa
	// Jika error nya berasal dari validasi .  maka ok akan terisi
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := responseDto.CommonResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

// Contoh internal error jika salah query database
// kesalahan mengambil params
func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := responseDto.CommonResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
