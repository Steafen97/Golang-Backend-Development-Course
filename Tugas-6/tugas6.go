package main

import (
	"fmt"
)

func main() {
	// Soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	inputJarijari(&luasLigkaran, &kelilingLingkaran, 15)
	fmt.Printf("Luas Lingkaran : %g\n", luasLigkaran)
	fmt.Printf("Address : %p\n", &luasLigkaran)
	fmt.Printf("Keliling Lingkaran : %g\n", kelilingLingkaran)
	fmt.Printf("Address : %p\n\n", &kelilingLingkaran)

	//Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	fmt.Printf("\n")

	// Soal 3
	var buah = []string{}

	tambahkanBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, satuan := range buah {
		fmt.Printf("%d. %s\n", i+1, satuan)
	}
	fmt.Printf("\n")

	// Soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, items := range dataFilm {
		fmt.Printf("%d.", i+1)

		for j, item := range items {
			if j == "duration" {
				fmt.Printf("\t%s : %s\n", j, item)
				continue
			}

			fmt.Printf("\t%s\t : %s\n", j, item)
		}

		fmt.Printf("\n")
	}
}

// Soal 1
func inputJarijari(luas *float64, keliling *float64, r float64) {
	*luas = 3.14 * r * r
	*keliling = 3.14 * 2 * r
}

// Soal 2
func introduce(kalimat *string, karakter ...string) {
	if karakter[1] == "perempuan" {
		*kalimat = "Bu " + karakter[0]
	} else {
		*kalimat = "Pak " + karakter[0]
	}

	*kalimat += " adalah seorang " + karakter[2] + " yang berusia " + karakter[3] + " tahun"
}

// Soal 3
func tambahkanBuah(buah *[]string, fruits ...string) {
	for _, fruit := range fruits {
		*buah = append(*buah, fruit)
	}
}

// Soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	data := map[string]string{}

	data["title"] = title
	data["duration"] = duration
	data["genre"] = genre
	data["year"] = year

	*dataFilm = append(*dataFilm, data)
}
