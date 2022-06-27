package main

import (
	"log"
	"net/http"
	"perpustakaan2/storage/connection"
	"perpustakaan2/storage/controllers"
	"perpustakaan2/storage/repositories"
	"perpustakaan2/storage/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config File !")
	}
}

func main() {

	koneksidb := connection.ConnectToDb()
	repo := repositories.NewRepo(koneksidb)
	usecase := usecase.NewUsecase(repo)
	ctrl := controllers.NewController(usecase)

	r := chi.NewRouter()

	r.Get("/get-data-buku", ctrl.GetAllBuku)
	r.Post("/post-data-buku", ctrl.PostDataBuku)
	r.Get("/get-detail-buku/{id}", ctrl.GetDataDetailBuku)
	r.Delete("/delete-data-buku/{id}", ctrl.DeleteDataBuku)
	r.Put("/edit-data-buku/{id}", ctrl.UpdateBuku)

	r.Get("/get-data-rak", ctrl.GetAllRak)
	r.Post("/post-data-rak", ctrl.PostDataRak)
	r.Get("/get-detail-rak/{id}", ctrl.GetDetailRak)
	r.Delete("/delete-data-rak/{id}", ctrl.DeleteDataRak)
	r.Put("/update-data-rak/{id}", ctrl.UpdateDataRak)

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Println("Error Starting Service !")
	}

}
