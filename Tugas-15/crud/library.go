package crud

import (
	"errors"
)

func ValidasiNilai(nilai int) (string, error) {
	if nilai > 100 && nilai < 0 {
		return "fail", errors.New("Nilai tidak boleh lebih dari 100 dan kurang dari 0")
	}

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

	return index, nil
}

