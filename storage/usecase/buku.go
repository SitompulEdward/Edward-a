package usecase

import (
	"errors"
	"perpustakaan2/storage/models"
)

func (r *UC) GetDataBuku() ([]models.Buku, error) {
	var modelbuku []models.Buku

	err := r.queryrepo.SelectAllBuku(modelbuku)
	if err != nil {
		return modelbuku, err
	}

	return modelbuku, nil
}

func (r *UC) InsertBuku(data models.Buku) error {
	err := r.queryrepo.InsertDataBuku(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) SelectDetailBuku(id int) (data models.Buku, err error) {
	data, err = r.queryrepo.SelectBukuById(id)

	if err != nil {
		return data, errors.New("data tidak ada")
	}

	return data, nil
}

func (r *UC) DeleteBuku(data models.Buku) error {
	buku, err := r.queryrepo.SelectBukuById(data.Id)
	if err != nil {
		return errors.New("Buku tidak ditemukan")
	}

	buku.Id = data.Id

	err = r.queryrepo.DeleteDataBuku(buku)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) UpdateDataBuku(data models.Buku) error {
	buku, err := r.queryrepo.SelectBukuById(data.Id)
	if err != nil {
		return errors.New("Buku tidak ditemukan")
	}

	buku.Judul = data.Judul
	buku.Pengarang = data.Pengarang
	buku.Penerbit = data.Penerbit
	buku.Tahun_Terbit = data.Tahun_Terbit
	buku.No_Rak = data.No_Rak

	err = r.queryrepo.UpdateDataBuku(data)
	if err != nil {
		return err
	}

	return nil
}
