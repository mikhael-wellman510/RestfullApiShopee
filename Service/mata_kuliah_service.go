package service

import (
	"context"
	"golang-shopeeMart-api/Model/web/request"
	"golang-shopeeMart-api/Model/web/response"
)

type MataKuliahService interface {
	Create(ctx context.Context, request request.MataKuliahRequest) response.MataKuliahResponse
	Update(ctx context.Context, request request.MataKuliahRequest) response.MataKuliahResponse
	Delete(ctx context.Context, matapelajaranId int)
	FindById(ctx context.Context, matapelajaranId int) response.MataKuliahResponse
	FindAll(ctx context.Context) []response.MataKuliahResponse
}
