package repository

import (
	"context"
	"database/sql"
	"golang-shopeeMart-api/Model/domain"
)

type MataKuliahRepository interface {
	Save(ctx context.Context, tx *sql.Tx, mataKuliah domain.MataKuliah) domain.MataKuliah
	Update(ctx context.Context, tx *sql.Tx, mataKuliah domain.MataKuliah) domain.MataKuliah
	Delete(ctx context.Context, tx *sql.Tx, mataKuliah domain.MataKuliah)
	FindById(ctx context.Context, tx *sql.Tx, mataKuliah int) (domain.MataKuliah, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MataKuliah
}
