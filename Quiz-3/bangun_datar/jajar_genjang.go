package bangun_datar

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func JajarGenjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.URL.Query().Get("hitung")
	sisiString := r.URL.Query().Get("sisi")
	alasString := r.URL.Query().Get("alas")
	tinggiString := r.URL.Query().Get("tinggi")

	if hitung == "" || sisiString == "" || alasString == "" || tinggiString == "" {
		fmt.Fprintf(w, "In this path, you are expected to give params of:\n - sisi\n - alas\n - tinggi\n - hitung\n\n")
		return
	}

	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan interface{}, 3)
	var sisi, alas, tinggi interface{}
	counter := 0

	wg.Add(1)
	go ValidasiHitung(ch1, hitung, &wg, &counter)

	wg.Add(1)
	go ValidasiSisi(ch2, sisiString, &wg, &counter)

	wg.Add(1)
	go ValidasiAlas(ch2, alasString, &wg, &counter)

	wg.Add(1)
	go ValidasiTinggi(ch2, tinggiString, &wg, &counter)

	hitung = <-ch1
	sisi = <-ch2
	alas = <-ch2
	tinggi = <-ch2
	wg.Wait()

	if counter != 0 {
		print := BadInputParams(hitung, sisi, alas, tinggi)
		fmt.Fprintln(w, print)
		return
	}

	var hasil int64
	if hitung == "keliling" {
		hasil = kelilingJajarGenjang(sisi, alas, tinggi)
	} else {
		hasil = luasJajarGenjang(sisi, alas, tinggi)
	}

	fmt.Fprintf(w, "%v Jajar Genjang adalah %v", strings.Title(hitung), hasil)
}

func kelilingJajarGenjang(sisi, alas, tinggi interface{}) int64 {
	return 2 * (alas.(int64) + sisi.(int64))
}

func luasJajarGenjang(sisi, alas, tinggi interface{}) int64 {
	return alas.(int64) * tinggi.(int64)
}