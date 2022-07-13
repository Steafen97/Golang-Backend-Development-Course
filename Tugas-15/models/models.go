package models

import "time"

type (
	Mahasiswa struct {
		ID        int       `json:"id"`
		Nama      string    `json:"nama"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	MataKuliah struct {
		ID        int       `json:"id"`
		Nama      string    `json:"nama"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Nilai struct {
		ID           int       `json:"id"`
		Nilai        int       `json:"nilai"`
		IndeksNilai  string    `json:"indeks_nilai"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		MahasiswaID  int       `json:"mahasiswa_id"`
		MataKuliahID int       `json:"mata_kuliah_id"`
	}
)
