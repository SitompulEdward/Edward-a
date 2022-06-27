package models

type Rak struct {
	Id     int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Lokasi string `json:"lokasi"`
}
