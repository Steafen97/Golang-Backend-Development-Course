package main

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return math.Pi * 2 * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang + b.lebar + b.tinggi))
}

// Soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type outputPhone interface {
	printOutput()
}

func (p phone) printOutput() {
	fmt.Printf("Name\t: %s\nBrand\t: %s\nYear\t: %d\n", p.name, p.brand, p.year)
	fmt.Printf("Colors\t: %s\n\n", strings.Join(p.colors, ", "))
}

// Soal 3
func luasPersegi(sisi int, kondisi bool) interface{} {
	var kosong interface{}
	if sisi == 0 {
		if kondisi {
			kosong = "Maaf anda belum menginput sisi dari persegi"
		}
	} else {
		if kondisi {
			kosong = fmt.Sprintf("Luas persegi dengan sisi %d cm adalah %d", sisi, sisi*sisi)
		} else {
			kosong = sisi * sisi
		}
	}
	return kosong
}

func main() {
	// Soal 1
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{6, 10}
	fmt.Println("==== Segitiga Sama Sisi ===")
	fmt.Println("Luas     :", bangunDatar.luas())
	fmt.Printf("Keliling : %d\n\n", bangunDatar.keliling())

	bangunDatar = persegiPanjang{10, 8}
	fmt.Println("====== Persegi Panjang ====")
	fmt.Println("Luas     :", bangunDatar.luas())
	fmt.Printf("Keliling : %d\n\n", bangunDatar.keliling())

	bangunRuang = tabung{5, 12}
	fmt.Println("========== Tabung =========")
	fmt.Println("Volume         :", bangunRuang.volume())
	fmt.Printf("Luas Permukaan : %g\n\n", bangunRuang.luasPermukaan())

	bangunRuang = balok{6, 7, 8}
	fmt.Println("========== Balok ==========")
	fmt.Println("Volume         :", bangunRuang.volume())
	fmt.Printf("Luas Permukaan : %g\n\n", bangunRuang.luasPermukaan())

	// Soal 2
	var hp outputPhone

	colors := []string{"Red", "Grey", "Green", "Yellow", "Blue", "Purple"}
	hp = phone{"Nokia Zenfone", "Nokia", 2020, colors}
	hp.printOutput()

	// Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))
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
