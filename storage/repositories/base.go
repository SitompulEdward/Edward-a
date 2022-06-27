package repositories

import (
	"perpustakaan2/storage/models"

	"gorm.io/gorm"
)

type repo struct {
	apps *gorm.DB
}

type Repo interface {
	SelectAllBuku(i interface{}) error
	InsertDataBuku(i interface{}) error
	SelectBukuById(id int) (data models.Buku, err error)
	DeleteDataBuku(i models.Buku) error
	UpdateDataBuku(i models.Buku) error

	SelectAllRak(i interface{}) error
	InsertDataRak(i interface{}) error
	SelectRakById(id int) (data models.Rak, err error)
	DeleteDataRak(i models.Rak) error
	UpdateDataRak(i models.Rak) error
}

func NewRepo(apps *gorm.DB) Repo {
	return &repo{apps}
}
