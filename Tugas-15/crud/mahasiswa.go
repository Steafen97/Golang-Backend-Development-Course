package crud

import (
	"Tugas-15/models"
	"Tugas-15/repository_mysql/mahasiswa"
	"Tugas-15/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mahasiswas, err := mahasiswa.GetAll(ctx)
	if err != nil {
		fmt.Println(err.Error())
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

	if err = mahasiswa.Insert(ctx, siswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data created successfully!",
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
	if err = mahasiswa.Update(ctx, siswa, idSiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data updated successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteMahasiswa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
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
		"status": "Data deleted successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

