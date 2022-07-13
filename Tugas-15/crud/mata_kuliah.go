package crud

import (
	"Tugas-15/models"
	"Tugas-15/repository_mysql/mata_kuliah"
	"Tugas-15/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetMataKuliah(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	matkuls, err := mata_kuliah.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, matkuls, http.StatusOK)
}

func PostMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var matkul models.MataKuliah

	if err = json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err = mata_kuliah.Insert(ctx, matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data created successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var matkul models.MataKuliah

	if err = json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idMatkul := ps.ByName("id")
	if err = mata_kuliah.Update(ctx, matkul, idMatkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data updated successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteMataKuliah(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idMatkul := ps.ByName("id")

	if err := mata_kuliah.Delete(ctx, idMatkul); err != nil {
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
