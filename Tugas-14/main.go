package main

import (
	"Tugas-14/mahasiswa"
	"Tugas-14/models"
	"Tugas-14/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa", PostMahasiswa)
	router.PUT("/mahasiswa/:id", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", DeleteMahasiswa)

	fmt.Println("Server running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func validasiNilai(nilai int) (string, error) {
	if nilai > 100 && nilai < 0 {
		return "fail", errors.New("Nilai tidak boleh lebih dari 100 dan kurang dari 0")
	}

	var index string

	if nilai >= 80 {
		index = "A"
	} else if nilai >= 70 {
		index = "B"
	} else if nilai >= 60 {
		index = "C"
	} else if nilai >= 50 {
		index = "D"
	} else {
		index = "E"
	}

	return index, nil
}

func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mahasiswas, err := mahasiswa.GetAll(ctx)

	if err != nil {
		fmt.Println()
	}

	utils.ResponseJSON(w, mahasiswas, http.StatusOK)
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var siswa models.Mahasiswa

	if err = json.NewDecoder(r.Body).Decode(&siswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	siswa.IndeksNilai, err = validasiNilai(siswa.Nilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err = mahasiswa.Insert(ctx, siswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var siswa models.Mahasiswa

	if err = json.NewDecoder(r.Body).Decode(&siswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idSiswa = ps.ByName("id")
	siswa.IndeksNilai, err = validasiNilai(siswa.Nilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err = mahasiswa.Update(ctx, siswa, idSiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idSiswa = ps.ByName("id")

	if err := mahasiswa.Delete(ctx, idSiswa); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
