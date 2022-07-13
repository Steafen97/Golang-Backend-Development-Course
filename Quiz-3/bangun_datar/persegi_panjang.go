package bangun_datar

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func PersegiPanjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.URL.Query().Get("hitung")
	panjangString := r.URL.Query().Get("panjang")
	lebarString := r.URL.Query().Get("lebar")

	if hitung == "" || panjangString == "" || lebarString == "" {
		fmt.Fprintf(w, "In this path, you are expected to give params of:\n - panjang\n - lebar\n - hitung\n\n")
		return
	}

	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan interface{}, 2)
	var panjang, lebar interface{}
	counter := 0

	wg.Add(1)
	go ValidasiHitung(ch1, hitung, &wg, &counter)

	wg.Add(1)
	go ValidasiPanjang(ch2, panjangString, &wg, &counter)

	wg.Add(1)
	go ValidasiLebar(ch2, lebarString, &wg, &counter)

	hitung = <-ch1
	panjang = <-ch2
	lebar = <-ch2
	wg.Wait()

	if counter != 0 {
		print := BadInputParams(hitung, panjang, lebar)
		fmt.Fprintln(w, print)
		return
	}

	var hasil int64
	if hitung == "keliling" {
		hasil = kelilingPersegiPanjang(panjang, lebar)
	} else {
		hasil = luasPersegiPanjang(panjang, lebar)
	}

	fmt.Fprintf(w, "%v Persegi Panjang adalah %v", strings.Title(hitung), hasil)
}

func kelilingPersegiPanjang(panjang, lebar interface{}) int64 {
	return 2 * (panjang.(int64) * lebar.(int64))
}

func luasPersegiPanjang(panjang, lebar interface{}) int64 {
	return panjang.(int64) * lebar.(int64)
}