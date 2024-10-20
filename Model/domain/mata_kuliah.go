package domain

import "time"

type MataKuliah struct {
	Id             int
	KodeMataKuliah string
	NamaMataKuliah string
	Sks            int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
