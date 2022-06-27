package repositories

import "perpustakaan2/storage/models"

func (r *repo) SelectAllBuku(i interface{}) error {
	result := r.apps.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) InsertDataBuku(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) SelectBukuById(id int) (data models.Buku, err error) {
	err = r.apps.First(&data, id).Error
	return
}

func (r *repo) DeleteDataBuku(i models.Buku) error {
	result := r.apps.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) UpdateDataBuku(i models.Buku) error {
	result := r.apps.Save(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// __RAK__

func (r *repo) SelectAllRak(i interface{}) error {
	result := r.apps.Find(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) InsertDataRak(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) SelectRakById(id int) (data models.Rak, err error) {
	err = r.apps.First(&data, id).Error
	return
}

func (r *repo) DeleteDataRak(i models.Rak) error {
	result := r.apps.Delete(i)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) UpdateDataRak(i models.Rak) error {
	result := r.apps.Save(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
