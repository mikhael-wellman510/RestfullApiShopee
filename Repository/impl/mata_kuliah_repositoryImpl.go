package impl

import (
	"context"
	"database/sql"
	"errors"
	helper "golang-shopeeMart-api/Helper"
	"golang-shopeeMart-api/Model/domain"
)

type MataKuliahRepositoryImpl struct {
}

func (repository *MataKuliahRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mataKuliah domain.MataKuliah) domain.MataKuliah {
	SQL := "insert into matakuliah (kode_matakuliah , nama_matakuliah , sks , created_at) values ($1,$2,$3,$4) returning id"
	var id int

	err := tx.QueryRow(SQL, mataKuliah.KodeMataKuliah, mataKuliah.NamaMataKuliah, mataKuliah.Sks, mataKuliah.CreatedAt).Scan(&id)
	helper.PanicIfErr(err)

	// Jgn lupa pasang Close Connection
	mataKuliah.Id = id

	return mataKuliah

}

func (repository *MataKuliahRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, mataKuliah domain.MataKuliah) domain.MataKuliah {
	SQL := "update matakuliah set kode_matakuliah = $1 , nama_matakuliah = $2 , sks = $3 , updated_at = $4 where id = $5"

	_, err := tx.ExecContext(ctx, SQL, mataKuliah.KodeMataKuliah, mataKuliah.NamaMataKuliah, mataKuliah.Sks, mataKuliah.Id)

	// Jgn lupa close connection
	helper.PanicIfErr(err)

	return mataKuliah
}

func (repository *MataKuliahRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, matakuliahId int) {
	SQL := "delete from matakuliah where id=$1"

	_, err := tx.ExecContext(ctx, SQL, matakuliahId)

	helper.PanicIfErr(err)
	// Buat kondisi jikalau succes menghapus true , kalau gagal false
}

func (repository *MataKuliahRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, matakuliahId int) (domain.MataKuliah, error) {
	SQL := "select * from matakuliah where id=$1"

	rows, err := tx.QueryContext(ctx, SQL, matakuliahId)
	helper.PanicIfErr(err)

	// Close Connection
	defer rows.Close()

	mataKuliah := domain.MataKuliah{}

	// Cek Apakah data ada atau tidak
	if rows.Next() {
		// Rows mengisi data ke instance new mataKuliah
		err := rows.Scan(&mataKuliah.Id, &mataKuliah.KodeMataKuliah, &mataKuliah.NamaMataKuliah, &mataKuliah.Sks, &mataKuliah.CreatedAt, &mataKuliah.UpdatedAt)
		helper.PanicIfErr(err)

		// Return jika metemukan hasil nya
		return mataKuliah, nil
	} else {
		return mataKuliah, errors.New("category is not found")
	}

}

func (repository *MataKuliahRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.MataKuliah {
	SQL := "select * from matakuliah"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)

	// Close Connection
	defer rows.Close()

	var mataKuliah []domain.MataKuliah

	if rows.Next() {
		mk := domain.MataKuliah{}
		err := rows.Scan(mk.Id, mk.KodeMataKuliah, mk.NamaMataKuliah, mk.Sks, mk.CreatedAt, mk.UpdatedAt)
		helper.PanicIfErr(err)
		mataKuliah = append(mataKuliah, mk)
	}

	return mataKuliah

}
