package request

type MataKuliahRequest struct {
	Id             int
	KodeMataKuliah string `validate:"required , min=0 , max=100" json:"kodeMataKuliah"`
	NamaMataKuliah string `validate:"required , min=0 , max=100" json:"namaMataKuliah"`
	Sks            int    `validate:"required , gte=1 , lte=5" json:"sks"`
}
