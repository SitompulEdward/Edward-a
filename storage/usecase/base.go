package usecase

import (
	"perpustakaan2/storage/models"
	"perpustakaan2/storage/repositories"
)

type UC struct {
	queryrepo repositories.Repo
}

type Usecase interface {
	GetDataBuku() ([]models.Buku, error)
	InsertBuku(models.Buku) error
	SelectDetailBuku(id int) (data models.Buku, err error)
	DeleteBuku(models.Buku) error
	UpdateDataBuku(models.Buku) error

	GetDataRak() ([]models.Rak, error)
	InsertRak(models.Rak) error
	SelectDetailRak(id int) (data models.Rak, err error)
	DeleteRak(models.Rak) error
	UpdateRak(models.Rak) error
}

func NewUsecase(r repositories.Repo) Usecase {
	return &UC{queryrepo: r}
}
