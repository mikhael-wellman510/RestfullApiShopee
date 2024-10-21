package domain

import "time"

type MataKuliah struct {
	Id             int
	KodeMataKuliah string
	NamaMataKuliah string
	Sks            int
	CreatedAt      time.Time
	// agar nilai bisa di set Null , maka pakai pointer
	UpdatedAt *time.Time
}
