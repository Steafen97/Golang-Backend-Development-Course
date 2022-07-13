package models

import "time"

type (
	Mahasiswa struct {
		ID          int       `json:"id"`
		Nama        string    `json:"nama"`
		MataKuliah  string    `json:"mata_kuliah"`
		IndeksNilai string    `json:"indeks_nilai"`
		Nilai       int       `json:"nilai"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
