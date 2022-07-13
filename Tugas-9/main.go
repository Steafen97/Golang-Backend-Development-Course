package main

import (
	"Tugas-9/library"
	"fmt"
	"strings"
)

func main() {
	// Soal 1
	var bangunDatar library.HitungBangunDatar
	var bangunRuang library.HitungBangunRuang

	bangunDatar = library.SegitigaSamaSisi{Alas: 6, Tinggi: 10}
	fmt.Println("==== Segitiga Sama Sisi ===")
	fmt.Println("Luas     :", bangunDatar.Luas())
	fmt.Printf("Keliling : %d\n\n", bangunDatar.Keliling())

	bangunDatar = library.PersegiPanjang{Panjang: 10, Lebar: 8}
	fmt.Println("====== Persegi Panjang ====")
	fmt.Println("Luas     :", bangunDatar.Luas())
	fmt.Printf("Keliling : %d\n\n", bangunDatar.Keliling())

	bangunRuang = library.Tabung{JariJari: 5, Tinggi: 12}
	fmt.Println("========== Tabung =========")
	fmt.Println("Volume         :", bangunRuang.Volume())
	fmt.Printf("Luas Permukaan : %g\n\n", bangunRuang.LuasPermukaan())

	bangunRuang = library.Balok{Panjang: 6, Lebar: 7, Tinggi: 8}
	fmt.Println("========== Balok ==========")
	fmt.Println("Volume         :", bangunRuang.Volume())
	fmt.Printf("Luas Permukaan : %g\n\n", bangunRuang.LuasPermukaan())

	// Soal 2
	var hp library.OutputPhone

	colors := []string{"Red", "Grey", "Green", "Yellow", "Blue", "Purple"}
	hp = library.Phone{Name: "Nokia Zenfone", Brand: "Nokia", Year: 2020, Colors: colors}
	hp.PrintOutput()

	// Soal 3
	fmt.Println(library.LuasPersegi(4, true))
	fmt.Println(library.LuasPersegi(8, false))
	fmt.Println(library.LuasPersegi(0, true))
	fmt.Println(library.LuasPersegi(0, false))
	fmt.Printf("\n")

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	pertama := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(kumpulanAngkaPertama.([]int))), " + "), "[]")
	kedua := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(kumpulanAngkaKedua.([]int))), " + "), "[]")
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)
	total := angkaPertama[0] + angkaPertama[1] + angkaKedua[0] + angkaKedua[1]
	fmt.Printf("%s%s + %s = %d\n", prefix.(string), pertama, kedua, total)
}
