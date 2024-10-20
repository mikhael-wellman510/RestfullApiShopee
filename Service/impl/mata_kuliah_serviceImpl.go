package impl

import (
	"context"
	"database/sql"
	helper "golang-shopeeMart-api/Helper"
	"golang-shopeeMart-api/Model/domain"
	"golang-shopeeMart-api/Model/web/request"
	"golang-shopeeMart-api/Model/web/response"
	repository "golang-shopeeMart-api/Repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type MataKuliahServiceImpl struct {
	MataKuliahRepository repository.MataKuliahRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func (service *MataKuliahServiceImpl) Create(ctx context.Context, request request.MataKuliahRequest) response.MataKuliahResponse {
	// Validasi Request masuk
	errValidation := service.Validate.Struct(request)

	helper.PanicIfErr(errValidation)
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)

	//Validate transaction
	defer helper.CommitOrRollback(tx)

	mataKuliah := domain.MataKuliah{
		KodeMataKuliah: request.KodeMataKuliah,
		NamaMataKuliah: request.NamaMataKuliah,
		Sks:            request.Sks,
		CreatedAt:      time.Now(),
	}

	mataKuliah = service.MataKuliahRepository.Save(ctx, tx, mataKuliah)

	return response.MataKuliahResponse{
		Id:             mataKuliah.Id,
		KodeMataKuliah: mataKuliah.KodeMataKuliah,
		NamaMataKuliah: mataKuliah.NamaMataKuliah,
		Sks:            mataKuliah.Sks,
		CreatedAt:      mataKuliah.CreatedAt,
	}

}
