package crud

import (
	"Quiz-3/models"
	call "Quiz-3/repository_mysql/category"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categories, err := call.GetAll(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	utils.ResponseJSON(w, categories, http.StatusOK)
}

func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var err error
	if res := Middleware(r); res != "" {
		utils.ResponseJSON(w, res, http.StatusUnauthorized)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var category models.Category

	if err = json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err = call.Insert(ctx, category); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data created successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	if res := Middleware(r); res != "" {
		utils.ResponseJSON(w, res, http.StatusUnauthorized)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var category models.Category

	if err = json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var categoryID = ps.ByName("id")
	if err = call.Update(ctx, category, categoryID); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data updated successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	if res := Middleware(r); res != "" {
		utils.ResponseJSON(w, res, http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var categoryID = ps.ByName("id")

	if err = call.Delete(ctx, categoryID); err != nil {
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

func ShowBooksByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categoryID := ps.ByName("id")
	books, err := call.BooksByForeignKey(ctx, categoryID)
	if err != nil {
		fmt.Println(err.Error())
	}

	books = BooksFilter(books, r)

	utils.ResponseJSON(w, books, http.StatusOK)
}