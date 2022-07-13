package main

import (
	bangun "Quiz-3/bangun_datar"
	. "Quiz-3/crud"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// Soal 2
	router.GET("/bangun-datar/segitiga-sama-sisi", bangun.SegitigaSamaSisi)
	router.GET("/bangun-datar/persegi", bangun.Persegi)
	router.GET("/bangun-datar/persegi-panjang", bangun.PersegiPanjang)
	router.GET("/bangun-datar/lingkaran", bangun.Lingkaran)
	router.GET("/bangun-datar/jajar-genjang", bangun.JajarGenjang)

	// CRUD Categories
	router.GET("/categories", GetCategories)
	router.POST("/categories", PostCategory)
	router.PUT("/categories/:id", UpdateCategory)
	router.DELETE("/categories/:id", DeleteCategory)
	router.GET("/categories/:id/books", ShowBooksByCategory)

	// CRUD Books
	router.GET("/books", GetBooks)
	router.POST("/books", PostBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)

	fmt.Println("Server running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}