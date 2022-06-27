package usecase

import (
	"errors"
	"perpustakaan2/storage/models"
)

func (r *UC) GetDataRak() ([]models.Rak, error) {
	var modelrak []models.Rak

	err := r.queryrepo.SelectAllRak(modelrak)
	if err != nil {
		return modelrak, err
	}

	return modelrak, nil
}

func (r *UC) InsertRak(data models.Rak) error {
	err := r.queryrepo.InsertDataRak(&data)
	if err != nil {
		return err
	}
	return nil
}

func (r *UC) SelectDetailRak(id int) (data models.Rak, err error) {
	data, err = r.queryrepo.SelectRakById(id)

	if err != nil {
		return data, errors.New("data tidak ada")
	}

	return data, nil
}

func (r *UC) DeleteRak(data models.Rak) error {
	rak, err := r.queryrepo.SelectRakById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	rak.Id = data.Id
	rak.Lokasi = data.Lokasi

	err = r.queryrepo.DeleteDataRak(rak)
	if err != nil {
		return err
	}

	return nil
}

func (r *UC) UpdateRak(data models.Rak) error {
	rak, err := r.queryrepo.SelectRakById(data.Id)
	if err != nil {
		return errors.New("admin tidak ditemukan")
	}

	rak.Id = data.Id
	rak.Lokasi = data.Lokasi

	err = r.queryrepo.UpdateDataRak(rak)
	if err != nil {
		return err
	}

	return nil
}
