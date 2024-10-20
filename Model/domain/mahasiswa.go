package domain

import (
	"time"
)

// Class
type Mahasiswa struct {
	Nim          string
	Nama         string
	TanggalLahir time.Time
	JenisKelamin string
	Alamat       string
	Email        string
	Jurusan      Jurusan
}
