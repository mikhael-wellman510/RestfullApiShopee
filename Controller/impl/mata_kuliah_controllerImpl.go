package implController

import (
	controller "golang-shopeeMart-api/Controller"
	helper "golang-shopeeMart-api/Helper"
	"golang-shopeeMart-api/Model/web/requestDto"
	"golang-shopeeMart-api/Model/web/responseDto"
	service "golang-shopeeMart-api/Service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MataKuliahControllerImpl struct {
	// Injection service nya
	MatakuliahService service.MataKuliahService
}

func NewMatakuliahController(matakuliahService service.MataKuliahService) controller.MataKuliahController {
	return &MataKuliahControllerImpl{
		MatakuliahService: matakuliahService,
	}
}

func (controller *MataKuliahControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	matakuliahRequest := requestDto.MataKuliahRequest{}

	helper.ReadFromRequestBody(request, &matakuliahRequest)

	mataKuliahResponse := controller.MatakuliahService.Create(request.Context(), matakuliahRequest)

	webResponse := responseDto.CommonResponse{
		Code:   200,
		Status: "OK",
		Data:   mataKuliahResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements controller.MataKuliahController.
func (controller *MataKuliahControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	matakuliahRequest := requestDto.MataKuliahRequest{}

	helper.ReadFromRequestBody(request, &matakuliahRequest)

	// Ambil data dari params
	matakuliahId := params.ByName("matakuliahId")
	id, err := strconv.Atoi(matakuliahId)
	helper.PanicIfErr(err)

	matakuliahRequest.Id = id

	matakuliahResponse := controller.MatakuliahService.Update(request.Context(), matakuliahRequest)

	webResponse := responseDto.CommonResponse{
		Code:   200,
		Status: "OK",
		Data:   matakuliahResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *MataKuliahControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Wajib di convert . karena setiap yang di kirim dari params berupa string
	matapelajaranId := params.ByName("matakuliahId")

	// Convert ke int
	id, err := strconv.Atoi(matapelajaranId)

	helper.PanicIfErr(err)

	matakuliahResponse := controller.MatakuliahService.FindById(request.Context(), id)

	webResponse := responseDto.CommonResponse{
		Code:   200,
		Status: "OK",
		Data:   matakuliahResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

// Delete implements controller.MataKuliahController.
func (controller *MataKuliahControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindAll implements controller.MataKuliahController.
func (controller *MataKuliahControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

// FindById implements controller.MataKuliahController.
