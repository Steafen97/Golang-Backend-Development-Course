package main

import (
	"fmt"
)

func main() {
	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	fmt.Printf("\n")

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	fmt.Printf("\n")

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
	fmt.Printf("\n")

	// Soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm := func(variables ...string) {
		data := map[string]string{}
		var key string

		for i, variable := range variables {
			switch i {
			case 0:
				key = "title"
			case 1:
				key = "jam"
			case 2:
				key = "genre"
			case 3:
				key = "tahun"
			}
			data[key] = variable
		}

		dataFilm = append(dataFilm, data)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// Soal 1
func luasPersegiPanjang(panjang int, lebar int) (luas int) {
	luas = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (keliling int) {
	keliling = 2 * (panjang + lebar)
	return
}

func volumeBalok(satuan ...int) int {
	volume := 1
	for _, satu := range satuan {
		volume *= satu
	}
	return volume
}

// Soal 2
func introduce(variable ...string) (kalimat string) {
	if variable[1] == "perempuan" {
		kalimat = "Bu " + variable[0]
	} else {
		kalimat = "Pak " + variable[0]
	}

	kalimat += " adalah seorang " + variable[2] + " yang berusia " + variable[3] + " tahun"
	return
}

// Soal 3
func buahFavorit(nama string, buahs ...string) (kalimat string) {
	kalimat = "halo nama saya " + nama + " dan buah favorit saya adalah "
	for i, buah := range buahs {
		if len(buahs) == i+1 {
			kalimat += `"` + buah + `"`
			break
		}
		kalimat += `"` + buah + `", `
	}
	return
}
