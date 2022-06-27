package controllers

import (
	"encoding/json"
	"log"
	"perpustakaan2/storage/connection"
	"perpustakaan2/storage/models"
	"strconv"

	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func (c *ctrl) GetAllBuku(w http.ResponseWriter, r *http.Request) {
	var listbuku []models.Buku

	DB.Find(&listbuku)
	datajson, err := json.Marshal(listbuku)

	if err != nil {
		w.Write([]byte("Error Convert TO JSON"))
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(datajson)
	w.WriteHeader(200)
	return
}

func (c *ctrl) PostDataBuku(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Buku

	err := decoder.Decode(&datarequest)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertBuku(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
}

func (c *ctrl) GetDataDetailBuku(w http.ResponseWriter, r *http.Request) {
	idbuku := chi.URLParam(r, "id")

	idConvert, err := strconv.Atoi(idbuku)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}
	log.Println(idConvert)
	data, err := c.us.SelectDetailBuku(idConvert)
	if err != nil {
		ResponseApi(w, 500, nil, err.Error())
	}

	ResponseApi(w, 200, data, "Sukses Get Data")
}

func (c *ctrl) DeleteDataBuku(w http.ResponseWriter, r *http.Request) {
	idbuku := chi.URLParam(r, "id")

	if idbuku == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	var datarequest models.Buku
	idConvert, err := strconv.Atoi(idbuku)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idConvert
	err = c.us.DeleteBuku(datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func (c *ctrl) UpdateBuku(w http.ResponseWriter, r *http.Request) {
	idbuku := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Buku
	err := decoder.Decode(&datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	idconvert, err := strconv.Atoi(idbuku)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idconvert
	err = c.us.UpdateDataBuku(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}

func ResponseApi(w http.ResponseWriter, code int, data interface{}, msg string) {

	resevice := models.Response{}
	resevice.Code = code
	resevice.Message = msg
	resevice.Data = data

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resevice)

}
