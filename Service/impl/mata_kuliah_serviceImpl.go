package implService

import (
	"context"
	"database/sql"
	exception "golang-shopeeMart-api/Exception"
	helper "golang-shopeeMart-api/Helper"
	"golang-shopeeMart-api/Model/domain"
	"golang-shopeeMart-api/Model/web/requestDto"
	service "golang-shopeeMart-api/Service"

	"golang-shopeeMart-api/Model/web/responseDto"
	repository "golang-shopeeMart-api/Repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type MataKuliahServiceImpl struct {
	// Ini data nya dari contructor nya (NewMatakuliahService)
	MataKuliahRepository repository.MataKuliahRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

// Tipe Data Repository
// Return Implement nya
func NewMatakuliahService(mataKuliahRepository repository.MataKuliahRepository, Db *sql.DB, validate *validator.Validate) service.MataKuliahService {

	// Constructor
	// Masukan data ke Attributnya
	return &MataKuliahServiceImpl{
		MataKuliahRepository: mataKuliahRepository,
		DB:                   Db,
		Validate:             validate,
	}
}

func (service *MataKuliahServiceImpl) Create(ctx context.Context, request requestDto.MataKuliahRequest) responseDto.MataKuliahResponse {
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
		UpdatedAt:      nil,
	}

	mataKuliah = service.MataKuliahRepository.Save(ctx, tx, mataKuliah)

	return responseDto.MataKuliahResponse{
		Id:             mataKuliah.Id,
		KodeMataKuliah: mataKuliah.KodeMataKuliah,
		NamaMataKuliah: mataKuliah.NamaMataKuliah,
		Sks:            mataKuliah.Sks,
		CreatedAt:      mataKuliah.CreatedAt,
		UpdatedAt:      mataKuliah.UpdatedAt,
	}

}

// Update implements service.MataKuliahService.
func (service *MataKuliahServiceImpl) Update(ctx context.Context, request requestDto.MataKuliahRequest) responseDto.MataKuliahResponse {
	// Nnti akan di tangkap di err handling
	errValidation := service.Validate.Struct(request)
	helper.PanicIfErr(errValidation)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)

	// Find By id punya 2 keluaran (data dan error)
	mataKuliah, err := service.MataKuliahRepository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError("id not found"))
	}

	mataKuliah.Id = request.Id
	mataKuliah.KodeMataKuliah = request.KodeMataKuliah
	mataKuliah.NamaMataKuliah = request.NamaMataKuliah
	mataKuliah.Sks = request.Sks
	now := time.Now()
	mataKuliah.UpdatedAt = &now

	mataKuliah = service.MataKuliahRepository.Update(ctx, tx, mataKuliah)

	return responseDto.MataKuliahResponse{
		Id:             mataKuliah.Id,
		KodeMataKuliah: mataKuliah.KodeMataKuliah,
		NamaMataKuliah: mataKuliah.NamaMataKuliah,
		Sks:            mataKuliah.Sks,
		CreatedAt:      mataKuliah.CreatedAt,
		UpdatedAt:      mataKuliah.UpdatedAt,
	}

}

// FindById implements service.MataKuliahService.
func (service *MataKuliahServiceImpl) FindById(ctx context.Context, matapelajaranId int) responseDto.MataKuliahResponse {
	tx, err := service.DB.Begin()

	helper.PanicIfErr(err)

	defer helper.CommitOrRollback(tx)

	// Find with repository
	matakuliah, err := service.MataKuliahRepository.FindById(ctx, tx, matapelajaranId)

	if err != nil {
		panic(exception.NewNotFoundError("Data Not Found"))
	}

	return responseDto.MataKuliahResponse{
		Id:             matakuliah.Id,
		KodeMataKuliah: matakuliah.KodeMataKuliah,
		NamaMataKuliah: matakuliah.NamaMataKuliah,
		Sks:            matakuliah.Sks,
		CreatedAt:      matakuliah.CreatedAt,
		UpdatedAt:      matakuliah.UpdatedAt,
	}

}

// FindAll implements service.MataKuliahService.
func (service *MataKuliahServiceImpl) FindAll(ctx context.Context) []responseDto.MataKuliahResponse {
	panic("unimplemented")
}

// Delete implements service.MataKuliahService.
func (service *MataKuliahServiceImpl) Delete(ctx context.Context, matapelajaranId int) {
	panic("unimplemented")
}
