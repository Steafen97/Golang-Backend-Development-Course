package main

import (
	"fmt"
	"strconv"
)

func main() {

	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)

	luasPP := panjang * lebar

	fmt.Printf("Luas Persegi Panjang dengan panjang %d dan lebar %d adalah : %d\n", panjang, lebar, luasPP)

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	luasSegitiga := (alas * tinggi) / 2

	fmt.Printf("Luas Segitiga dengan alas %d dan tinggi %d adalah : %d\n", alas, tinggi, luasSegitiga)

	// soal 2
	var nilaiJohn int = 80
	var nilaiDoe int = 50

	switch {
	case nilaiDoe >= 80:
		fmt.Println("Indeks Doe : A")
	case nilaiDoe >= 70 && nilaiDoe < 80:
		fmt.Println("Indeks Doe : B")
	case nilaiDoe >= 60 && nilaiDoe < 70:
		fmt.Println("Indeks Doe : C")
	case nilaiDoe >= 50 && nilaiDoe < 60:
		fmt.Println("Indeks Doe : D")
	case nilaiDoe < 50:
		fmt.Println("Indeks Doe : E")
	}

	if nilaiJohn >= 80 {
		fmt.Println("Indeks John : A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Indeks John : B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Indeks John : C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Indeks John : D")
	} else {
		fmt.Println("Indeks John : E")
	}

	// soal 3
	var tanggal int = 9
	var bulan int = 9
	var tahun int = 1997

	switch bulan {
	case 1:
		fmt.Printf("%d Januari %d\n", tanggal, tahun)
	case 2:
		fmt.Printf("%d Februari %d\n", tanggal, tahun)
	case 3:
		fmt.Printf("%d Maret %d\n", tanggal, tahun)
	case 4:
		fmt.Printf("%d April %d\n", tanggal, tahun)
	case 5:
		fmt.Printf("%d Mei %d\n", tanggal, tahun)
	case 6:
		fmt.Printf("%d Juni %d\n", tanggal, tahun)
	case 7:
		fmt.Printf("%d Juli %d\n", tanggal, tahun)
	case 8:
		fmt.Printf("%d Agustus %d\n", tanggal, tahun)
	case 9:
		fmt.Printf("%d September %d\n", tanggal, tahun)
	case 10:
		fmt.Printf("%d Oktober %d\n", tanggal, tahun)
	case 11:
		fmt.Printf("%d November %d\n", tanggal, tahun)
	case 12:
		fmt.Printf("%d Desember %d\n", tanggal, tahun)
	default:
		fmt.Println("Bulan tidak valid")
	}

	// soal 4
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Baby Boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generation X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generation Y")
	} else if tahun >= 1995 && tahun <= 2009 {
		fmt.Println("Generation Z")
	} else {
		fmt.Println("Masukkan tahun yang valid")
	}
}
