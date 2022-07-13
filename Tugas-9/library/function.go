package library

import (
	"fmt"
	"math"
	"strings"
)

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return math.Pi * 2 * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang + b.Lebar + b.Tinggi))
}

func (p Phone) PrintOutput() {
	fmt.Printf("Name\t: %s\nBrand\t: %s\nYear\t: %d\n", p.Name, p.Brand, p.Year)
	fmt.Printf("Colors\t: %s\n\n", strings.Join(p.Colors, ", "))
}

func LuasPersegi(sisi int, kondisi bool) interface{} {
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
