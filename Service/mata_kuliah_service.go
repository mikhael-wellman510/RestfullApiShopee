package service

import (
	"context"
	"golang-shopeeMart-api/Model/web/requestDto"
	"golang-shopeeMart-api/Model/web/responseDto"
)

type MataKuliahService interface {
	Create(ctx context.Context, request requestDto.MataKuliahRequest) responseDto.MataKuliahResponse
	Update(ctx context.Context, request requestDto.MataKuliahRequest) responseDto.MataKuliahResponse
	Delete(ctx context.Context, matapelajaranId int)
	FindById(ctx context.Context, matapelajaranId int) responseDto.MataKuliahResponse
	FindAll(ctx context.Context) []responseDto.MataKuliahResponse
}
