package main

import (
	"fmt"
)

// Soal 1
type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (bangun segitiga) luas() {
	fmt.Printf("Luas Segitiga : %d\n", bangun.alas*bangun.tinggi/2)
}

func (bangun persegi) luas() {
	fmt.Printf("Luas Persegi : %d\n", bangun.sisi*bangun.sisi)
}

func (bangun persegiPanjang) luas() {
	fmt.Printf("Luas Persegi Panjang : %d\n", bangun.panjang*bangun.lebar)
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (hp phone) tambahWarna(warna *[]string, colors ...string) {
	for _, color := range colors {
		*warna = append(*warna, color)
	}
}

// Soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	data := movie{title, genre, duration, year}
	*dataFilm = append(*dataFilm, data)
}

func main() {
	// Soal 1
	fruits := []buah{
		{"Nanas", "Kuning", false, 9000},
		{"Jeruk", "Oranye", true, 8000},
		{"Semangka", "Hijau & Merah", true, 10000},
		{"Pisang", "Kuning", false, 5000},
	}

	fmt.Printf("Nama\t\t|   Warna\t\t|   Ada Bijinya\t|   Harga\n")
	var strng string
	for _, fruit := range fruits {
		if fruit.adaBijinya == false {
			strng = "Tidak\t"
		} else {
			strng = "Ada\t\t"
		}
		if len(fruit.nama) >= 8 {
			fmt.Printf("%s\t|   %s\t|   %s|   %d\n", fruit.nama, fruit.warna, strng, fruit.harga)
			continue
		}
		fmt.Printf("%s\t\t|   %s\t\t|   %s|   %d\n", fruit.nama, fruit.warna, strng, fruit.harga)
	}
	fmt.Printf("\n")

	// Soal 2
	segitiga := segitiga{6, 4}
	persegi := persegi{10}
	persegiPanjang := persegiPanjang{10, 5}
	segitiga.luas()
	persegi.luas()
	persegiPanjang.luas()
	fmt.Printf("\n")

	// Soal 3
	objek := phone{}
	objek.name = "Nokia Zenfone"
	objek.brand = "HP Jadul"
	objek.year = 2012
	objek.tambahWarna(&objek.colors, "Merah", "Jingga", "Kuning", "Hijau", "Biru")
	fmt.Println(objek)
	fmt.Printf("\n")

	// Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("%d.", i+1)
		fmt.Printf("\tTitle\t : %s\n", item.title)
		fmt.Printf("\tGenre\t : %s\n", item.genre)
		fmt.Printf("\tYear\t : %d\n", item.year)
		fmt.Printf("\tDuration : %d jam\n", item.duration/60)
		if i < 3 {
			fmt.Printf("\n")
		}
	}
}
