package response

import "time"

type MataKuliahResponse struct {
	Id             int       `json:"id"`
	KodeMataKuliah string    `json:"kodeMataKuliah"`
	NamaMataKuliah string    `json:"namaMataKuliah"`
	Sks            int       `json:"sks"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAts"`
}
