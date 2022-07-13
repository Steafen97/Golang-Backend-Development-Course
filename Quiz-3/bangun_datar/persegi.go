package bangun_datar

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func Persegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.URL.Query().Get("hitung")
	sisiString := r.URL.Query().Get("sisi")

	if hitung == "" || sisiString == "" {
		fmt.Fprintf(w, "In this path, you are expected to give params of:\n - sisi\n - hitung\n\n")
		return
	}

	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan interface{})
	var sisi interface{}
	counter := 0

	wg.Add(1)
	go ValidasiHitung(ch1, hitung, &wg, &counter)

	wg.Add(1)
	go ValidasiSisi(ch2, sisiString, &wg, &counter)

	hitung = <-ch1
	sisi = <-ch2
	wg.Wait()

	if counter != 0 {
		print := BadInputParams(hitung, sisi)
		fmt.Fprintln(w, print)
		return
	}

	var hasil int64
	if hitung == "keliling" {
		hasil = kelilingPersegi(sisi)
	} else {
		hasil = luasPersegi(sisi)
	}

	fmt.Fprintf(w, "%v Persegi adalah %v", strings.Title(hitung), hasil)
}

func kelilingPersegi(sisi interface{}) int64 {
	return 2 * (sisi.(int64) * sisi.(int64))
}

func luasPersegi(sisi interface{}) int64 {
	return sisi.(int64) * sisi.(int64)
}
