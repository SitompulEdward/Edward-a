package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"perpustakaan2/storage/connection"
	"perpustakaan2/storage/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func init() {
	DB = connection.ConnectToDb()
}

func (c *ctrl) GetAllRak(w http.ResponseWriter, r *http.Request) {
	var listdatrak []models.Rak

	DB.Find(&listdatrak)
	datajson, err := json.Marshal(listdatrak)

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

func (c *ctrl) PostDataRak(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var datarequest models.Rak

	err := decoder.Decode(&datarequest)

	if err != nil {
		w.Write([]byte("Error Decode JSON Payload"))
		w.WriteHeader(500)
		return
	}

	err = c.us.InsertRak(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Insert Data")
	return
}

func (c *ctrl) GetDetailRak(w http.ResponseWriter, r *http.Request) {
	idrak := chi.URLParam(r, "id")

	idConvert, err := strconv.Atoi(idrak)

	if err != nil {
		ResponseApi(w, 500, nil, "Internal Server Error")
		return
	}

	log.Println(idConvert)
	data, err := c.us.SelectDetailRak(idConvert)
	if err != nil {
		ResponseApi(w, 500, nil, err.Error())
	}

	ResponseApi(w, 200, data, "Sukses Get Data")
}

func (c *ctrl) DeleteDataRak(w http.ResponseWriter, r *http.Request) {
	idrak := chi.URLParam(r, "id")

	if idrak == "" {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	var datarequest models.Rak
	idConvert, err := strconv.Atoi(idrak)

	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idConvert
	err = c.us.DeleteRak(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses delete")
	return
}

func (c *ctrl) UpdateDataRak(w http.ResponseWriter, r *http.Request) {
	idrak := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)
	var datarequest models.Rak
	err := decoder.Decode(&datarequest)

	if err != nil {
		ResponseApi(w, 500, nil, "INTERNAL SERVER ERROR")
		return
	}

	idconvert, err := strconv.Atoi(idrak)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	datarequest.Id = idconvert
	err = c.us.UpdateRak(datarequest)
	if err != nil {
		ResponseApi(w, 500, nil, "Invalid Request")
		return
	}

	ResponseApi(w, 200, nil, "Sukses Update Data")
	return
}
