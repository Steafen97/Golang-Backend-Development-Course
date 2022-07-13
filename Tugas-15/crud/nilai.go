package crud

import (
	"Tugas-15/models"
	"Tugas-15/repository_mysql/nilai"
	"Tugas-15/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nilais, err := nilai.GetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilais, http.StatusOK)
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var skor models.Nilai

	if err = json.NewDecoder(r.Body).Decode(&skor); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	skor.IndeksNilai, err = ValidasiNilai(skor.Nilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err = nilai.Insert(ctx, skor); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data created successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var skor models.Nilai

	if err = json.NewDecoder(r.Body).Decode(&skor); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	skor.IndeksNilai, err = ValidasiNilai(skor.Nilai)
	if err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	idNilai := ps.ByName("id")
	if err = nilai.Update(ctx, skor, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data updated successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteNilai(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idNilai := ps.ByName("id")

	if err := nilai.Delete(ctx, idNilai); err != nil {
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

