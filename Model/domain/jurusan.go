package domain

import "time"

// Class
type Jurusan struct {
	// Attribute
	Id          int
	NamaJurusan string
	Fakultas    string
	CreatedAt   time.Time
}
