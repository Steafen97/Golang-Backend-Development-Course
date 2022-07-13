package bangun_datar

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.URL.Query().Get("hitung")
	alasString := r.URL.Query().Get("alas")
	tinggiString := r.URL.Query().Get("tinggi")

	if hitung == "" || alasString == "" || tinggiString == "" {
		fmt.Fprintf(w, "In this path, you are expected to give params of:\n - alas\n - tinggi\n - hitung\n\n")
		return
	}

	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan interface{}, 2)
	var alas, tinggi interface{}
	counter := 0

	wg.Add(1)
	go ValidasiHitung(ch1, hitung, &wg, &counter)

	wg.Add(1)
	go ValidasiAlas(ch2, alasString, &wg, &counter)

	wg.Add(1)
	go ValidasiTinggi(ch2, tinggiString, &wg, &counter)

	hitung = <-ch1
	alas = <-ch2
	tinggi = <-ch2
	wg.Wait()

	if counter != 0 {
		print := BadInputParams(hitung, alas, tinggi)
		fmt.Fprintln(w, print)
		return
	}

	var hasil float64
	if hitung == "keliling" {
		hasil = kelilingSegitigaSamaSisi(alas, tinggi)
	} else {
		hasil = luasSegitigaSamaSisi(alas, tinggi)
	}

	fmt.Fprintf(w, "%v Segitiga Sama Sisi adalah %v", strings.Title(hitung), hasil)
}

func kelilingSegitigaSamaSisi(alas, tinggi interface{}) float64 {
	return float64(3 * alas.(int64))
}

func luasSegitigaSamaSisi(alas, tinggi interface{}) float64 {
	return float64(alas.(int64) * tinggi.(int64) / 2)
}
