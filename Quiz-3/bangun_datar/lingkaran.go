package bangun_datar

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
)

func Lingkaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hitung := r.URL.Query().Get("hitung")
	jariJariString := r.URL.Query().Get("jariJari")

	if hitung == "" || jariJariString == "" {
		fmt.Fprintf(w, "In this path, you are expected to give params of:\n - jariJari\n - hitung\n\n")
		return
	}

	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan interface{})
	var jariJari interface{}
	counter := 0

	wg.Add(1)
	go ValidasiHitung(ch1, hitung, &wg, &counter)

	wg.Add(1)
	go ValidasiJariJari(ch2, jariJariString, &wg, &counter)

	hitung = <-ch1
	jariJari = <-ch2
	wg.Wait()

	if counter != 0 {
		print := BadInputParams(hitung, jariJari)
		fmt.Fprintln(w, print)
		return
	}

	var hasil float64
	if hitung == "keliling" {
		hasil = kelilingLingkaran(jariJari)
	} else {
		hasil = luasLingkaran(jariJari)
	}

	fmt.Fprintf(w, "%v Lingkaran adalah %v", strings.Title(hitung), hasil)
}

func kelilingLingkaran(jariJari interface{}) float64 {
	return math.Pi * float64(2*jariJari.(int64))
}

func luasLingkaran(jariJari interface{}) float64 {
	return math.Pi * float64(jariJari.(int64)*jariJari.(int64))
}