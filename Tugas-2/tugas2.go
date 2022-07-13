package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	kata1 := "Bootcamp"
	kata2 := " Digital"
	kata3 := " Skill"
	kata4 := " Sanbercode"
	kata5 := " Golang"
	fmt.Println(kata1 + kata2 + kata3 + kata4 + kata5)

	// Soal 2
	halo := "Halo Dunia"
	find := "Dunia"
	replace := strings.Replace(halo, find, "Golang", 1)
	fmt.Println(replace)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	gabungan := kataPertama + " " + strings.Replace(kataKedua, "s", "S", 1) + " " + strings.Replace(kataKetiga, "r", "R", 1) + " " + strings.ToUpper(kataKeempat)
	fmt.Println(gabungan)

	// Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	jadiAngkaPertama, _ := strconv.ParseInt(angkaPertama, 10, 8)
	jadiAngkaKedua, _ := strconv.ParseInt(angkaKedua, 10, 4)
	jadiAngkaKetiga, _ := strconv.ParseInt(angkaKetiga, 10, 4)
	jadiAngkaKeempat, _ := strconv.ParseInt(angkaKeempat, 10, 4)
	jumlah := jadiAngkaPertama + jadiAngkaKedua + jadiAngkaKetiga + jadiAngkaKeempat
	fmt.Println(jumlah)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	olahKalimat1 := strings.Replace(kalimat, "halo", "Hi", 2)
	olahKalimat2 := strings.Replace(olahKalimat1, "H", `"H`, 1)
	olahKalimat3 := strings.Replace(olahKalimat2, "g", `g"`, 1)
	olahAngka := strconv.Itoa(angka)
	hasil := olahKalimat3 + " - " + olahAngka
	fmt.Println(hasil)
}
