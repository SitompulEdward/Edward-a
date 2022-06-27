package models

type Buku struct {
	Id           int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Judul        string `json:"judul"`
	Pengarang    string `json:"pengarang"`
	Penerbit     string `json:"penerbit"`
	Tahun_Terbit string `json:"tahun_terbit"`
	No_Rak       string `json:"no_rak"`
}
