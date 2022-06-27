package controllers

import (
	"net/http"
	"perpustakaan2/storage/usecase"
)

type ctrl struct {
	us usecase.Usecase
}

type ControllerInterface interface {
	GetAllBuku(w http.ResponseWriter, r *http.Request)
	PostDataBuku(w http.ResponseWriter, r *http.Request)
	GetDataDetailBuku(w http.ResponseWriter, r *http.Request)
	DeleteDataBuku(w http.ResponseWriter, r *http.Request)
	UpdateBuku(w http.ResponseWriter, r *http.Request)

	GetAllRak(w http.ResponseWriter, r *http.Request)
	PostDataRak(w http.ResponseWriter, r *http.Request)
	GetDetailRak(w http.ResponseWriter, r *http.Request)
	DeleteDataRak(w http.ResponseWriter, r *http.Request)
	UpdateDataRak(w http.ResponseWriter, r *http.Request)
}

func NewController(us usecase.Usecase) ControllerInterface {
	return &ctrl{us: us}
}
