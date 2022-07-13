package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/tambah-mahasiswa", Auth(http.HandlerFunc(PostMahasiswa)))
	http.HandleFunc("/ambil-mahasiswa", GetMahasiswa)

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nilaiMahasiswa NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&nilaiMahasiswa); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			Nama := r.PostFormValue("Nama")
			MataKuliah := r.PostFormValue("MataKuliah")
			getNilai := r.PostFormValue("Nilai")
			Nilai, _ := strconv.Atoi(getNilai)
			nilaiMahasiswa = NilaiMahasiswa{
				Nama:       Nama,
				MataKuliah: MataKuliah,
				Nilai:      uint(Nilai),
			}
		}
	}

	if nilaiMahasiswa.Nilai > 100 && nilaiMahasiswa.Nilai < 0 {
		w.Write([]byte("Nilai tidak boleh lebih dari 100 dan kurang dari 0"))
		return
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
	i := len(nilaiNilaiMahasiswa)
	nilaiNilaiMahasiswa[i-1].ID = uint(i)

	indeksNilai := IndeksNilai(nilaiNilaiMahasiswa[i-1].Nilai)
	nilaiNilaiMahasiswa[i-1].IndeksNilai = indeksNilai

	dataMahasiswa, _ := json.Marshal(nilaiNilaiMahasiswa)
	w.Write(dataMahasiswa)
	return
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func IndeksNilai(nilai uint) string {
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

	return index
}

