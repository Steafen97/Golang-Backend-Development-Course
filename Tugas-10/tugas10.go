package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	// Soal 3
	angka := 1

	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 1
	defer deferFunc("Golang Backend Development", 2021)

	// Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
	fmt.Printf("\n")

	// Soal 4
	var phones = []string{}
	tambahPhones(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	sortPhones(phones)
	fmt.Printf("\n")

	// Soal 5
	fmt.Println("Lingkaran Jari-jari 7 dengan pembulatan")
	fmt.Printf("Luas Lingkaran\t\t: %g\nKeliling Lingkaran\t: %g\n\n", luasLingkaran(7), kelilingLingkaran(7))
	fmt.Println("Lingkaran Jari-jari 10 dengan pembulatan")
	fmt.Printf("Luas Lingkaran\t\t: %g\nKeliling Lingkaran\t: %g\n\n", luasLingkaran(10), kelilingLingkaran(10))
	fmt.Println("Lingkaran Jari-jari 15 dengan pembulatan")
	fmt.Printf("Luas Lingkaran\t\t: %g\nKeliling Lingkaran\t: %g\n\n", luasLingkaran(15), kelilingLingkaran(15))

	// Soal 6
	panjang := flag.Int64("panjang", 6, "Enter the length")
	lebar := flag.Int64("lebar", 4, "Enter the width")

	flag.Parse()
	fmt.Printf("Persegi Panjang dengan Panjang %d dan Lebar %d\n", *panjang, *lebar)
	fmt.Printf("Luas Persegi Panjang\t\t: %d\n", luasPersegiPanjang(*panjang, *lebar))
	fmt.Printf("Keliling Persegi Panjang\t: %d\n\n", kelilingPersegiPanjang(*panjang, *lebar))
}

// Soal 1
func deferFunc(kalimat string, tahun int) {
	fmt.Printf("Eksekusi %s pada tahun %d\n\n", kalimat, tahun)
}

// Soal 2
func kelilingSegitigaSamaSisi(sisi int, kondisi bool) (output interface{}) {
	if sisi == 0 {
		if kondisi {
			output = errors.New("Maaf anda belum menginput sisi dari Segitiga Sama Sisi")
		} else {
			defer func() {
				output = recover()
			}()
			panic("Maaf anda belum menginput sisi dari Segitiga Sama Sisi")
		}
	} else {
		if kondisi {
			output = fmt.Sprintf("Keliling Segitiga Sama Sisi dengan sisi %d cm adalah %d", sisi, 3*sisi)
		} else {
			output = 3 * sisi
		}
	}
	return
}

// Soal 3
func tambahAngka(angkaBaru int, angka *int) {
	*angka += angkaBaru
}

func cetakAngka(angka *int) {
	fmt.Printf("Total angka : %d\n", *angka)
}

// Soal 4
func tambahPhones(phonesLama *[]string, phones ...string) {
	*phonesLama = append(*phonesLama, phones...)
}

func sortPhones(phones []string) {
	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}
}

// Soal 5
func luasLingkaran(jariJari float64) float64 {
	return math.Round(math.Pi * math.Pow(jariJari, 2))
}

func kelilingLingkaran(jariJari float64) float64 {
	return math.Round(2 * math.Pi * jariJari)
}

// Soal 6
func luasPersegiPanjang(panjang, lebar int64) int64 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int64) int64 {
	return 2 * (panjang + lebar)
}
