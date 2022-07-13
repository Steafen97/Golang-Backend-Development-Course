package bangun_datar

import (
	"strconv"
	"sync"
)

func ValidasiHitung(ch chan string, hitung string, wg *sync.WaitGroup, counter *int) {
	if hitung != "luas" && hitung != "keliling" {
		ch <- " - params hitung seharusnya luas atau keliling\n"
		*counter++
	} else {
		ch <- hitung
	}

	wg.Done()
}

func ValidasiAlas(ch chan interface{}, alasString string, wg *sync.WaitGroup, counter *int) {
	alas, err := strconv.ParseInt(alasString, 10, 64)
	if err != nil {
		*counter++
		ch <- " - params alas seharusnya angka\n"
	} else {
		ch <- alas
	}

	wg.Done()
}

func ValidasiTinggi(ch chan interface{}, tinggiString string, wg *sync.WaitGroup, counter *int) {
	tinggi, err := strconv.Atoi(tinggiString)
	if err != nil {
		*counter++
		ch <- " - params tinggi seharusnya angka\n"
	} else {
		ch <- int64(tinggi)
	}

	wg.Done()
}

func ValidasiSisi(ch chan interface{}, sisiString string, wg *sync.WaitGroup, counter *int) {
	sisi, err := strconv.ParseInt(sisiString, 10, 64)
	if err != nil {
		*counter++
		ch <- " - params sisi seharusnya angka\n"
	} else {
		ch <- int64(sisi)
	}

	wg.Done()
}

func ValidasiPanjang(ch chan interface{}, panjangString string, wg *sync.WaitGroup, counter *int) {
	panjang, err := strconv.Atoi(panjangString)
	if err != nil {
		*counter++
		ch <- " - params panjang seharusnya angka\n"
	} else {
		ch <- int64(panjang)
	}

	wg.Done()
}

func ValidasiLebar(ch chan interface{}, lebarString string, wg *sync.WaitGroup, counter *int) {
	lebar, err := strconv.ParseInt(lebarString, 10, 64)
	if err != nil {
		*counter++
		ch <- " - params lebar seharusnya angka\n"
	} else {
		ch <- int64(lebar)
	}

	wg.Done()
}

func ValidasiJariJari(ch chan interface{}, jariJariString string, wg *sync.WaitGroup, counter *int) {
	jariJari, err := strconv.Atoi(jariJariString)
	if err != nil {
		*counter++
		ch <- " - params jariJari seharusnya angka\n"
	} else {
		ch <- int64(jariJari)
	}

	wg.Done()
}

func BadInputParams(hitung string, satuan ...interface{}) string {
	errors := "Terdapat kesalahan pada params:\n"

	if hitung != "keliling" && hitung != "luas" {
		errors += hitung
	}

	for _, satu := range satuan {
		if w, ok := satu.(string); ok {
			errors += w
		}
	}

	return errors
}