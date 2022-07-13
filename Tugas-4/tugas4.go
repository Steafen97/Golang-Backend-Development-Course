package main

import (
	"fmt"
)

func main() {
	// Soal 1
	for i := 1; i <= 20; i++ {
		if i%2 == 1 && i%3 == 0 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d - Santai\n", i)
		} else {
			fmt.Printf("%d - Berkualitas\n", i)
		}
	}
	fmt.Printf("\n")

	// Soal 2
	height := 0
	width := 0
	for height < 7 {
		for width <= height {
			fmt.Print("#")
			width++
		}
		fmt.Printf("\n")
		height++
		width = 0
	}
	fmt.Printf("\n")

	// Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var olah = kalimat[2:]
	fmt.Println(olah)
	fmt.Println(kalimat[2] + " " + kalimat[3] + " " + kalimat[4] + " " + kalimat[5] + " " + kalimat[6] + "\n")

	// Soal 4
	var sayuran = []string{}
	sayurans := append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, sayur := range sayurans {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}
	fmt.Printf("\n")

	// Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volume = 1

	for i, satu := range satuan {
		fmt.Printf("%s = %d\n", i, satu)
		volume = volume * satu
	}
	fmt.Printf("Volume Balok = %d", volume)
	// fmt.Printf("Volume Balok = %d", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])
}
