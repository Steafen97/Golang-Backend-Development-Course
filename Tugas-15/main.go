package main

import (
	. "Tugas-15/crud"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// CRUD Mahasiswa
	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa", PostMahasiswa)
	router.PUT("/mahasiswa/:id", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", DeleteMahasiswa)

	// CRUD Mata Kuliah
	router.GET("/mata-kuliah", GetMataKuliah)
	router.POST("/mata-kuliah", PostMataKuliah)
	router.PUT("/mata-kuliah/:id", UpdateMataKuliah)
	router.DELETE("/mata-kuliah/:id", DeleteMataKuliah)

	// CRUD Nilai
	router.GET("/nilai", GetNilai)
	router.POST("/nilai", PostNilai)
	router.PUT("/nilai/:id", UpdateNilai)
	router.DELETE("/nilai/:id", DeleteNilai)

	fmt.Println("Server running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
