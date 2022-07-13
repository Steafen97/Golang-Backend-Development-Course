package crud

import (
	"Quiz-3/models"
	call "Quiz-3/repository_mysql/book"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	books, err := call.GetAll(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	books = BooksFilter(books, r)

	utils.ResponseJSON(w, books, http.StatusOK)
}

func PostBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	var book models.Book

	if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	invalids := InputValidation(book.ImageURL, book.ReleaseYear)
	if len(invalids) != 0 {
		utils.ResponseJSON(w, invalids, http.StatusBadRequest)
		return
	}

	book.Thickness = ThicknessConvertion(book.TotalPage)
	if err = call.Insert(ctx, book); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data created successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	var book models.Book

	if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	invalids := InputValidation(book.ImageURL, book.ReleaseYear)
	if len(invalids) != 0 {
		utils.ResponseJSON(w, invalids, http.StatusBadRequest)
		return
	}

	book.Thickness = ThicknessConvertion(book.TotalPage)
	var bookID = ps.ByName("id")
	if err = call.Update(ctx, book, bookID); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Data updated successfully!",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	if res := Middleware(r); res != "" {
		utils.ResponseJSON(w, res, http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var bookID = ps.ByName("id")

	if err = call.Delete(ctx, bookID); err != nil {
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