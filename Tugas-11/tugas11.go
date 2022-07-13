package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	wg.Add(1)
	go printPhones(phones, &wg)
	wg.Wait()
	fmt.Println()

	// Soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
	fmt.Println()

	// Soal 3
	jariJari := []int{8, 14, 20}
	luasCh := make(chan string, 3)
	kelilingCh := make(chan string, 3)
	volumeCh := make(chan string, 3)

	go luasLingkaran(jariJari, luasCh)
	go kelilingLingkaran(jariJari, kelilingCh)
	go volumeTabung(jariJari, 10, volumeCh)

	for i := 0; i < 3; i++ {
		fmt.Println(<-luasCh)
		fmt.Println(<-kelilingCh)
		fmt.Println(<-volumeCh)
	}

	// Soal 4
	luasCh = make(chan string)
	kelilingCh = make(chan string)
	volumeCh = make(chan string)
	panjang := 5
	lebar := 10
	tinggi := 15

	go luasPersegiPanjang(panjang, lebar, luasCh)
	go kelilingPersegiPanjang(panjang, lebar, kelilingCh)
	go volumeBalok(panjang, lebar, tinggi, volumeCh)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-luasCh:
			fmt.Println(luas)
		case keliling := <-kelilingCh:
			fmt.Println(keliling)
		case volume := <-volumeCh:
			fmt.Println(volume)
		}
	}
}

// Soal 1
func printPhones(phones []string, wg *sync.WaitGroup) {
	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}

	wg.Done()
}

// Soal 2
func getMovies(moviesChannel chan string, movies ...string) {
	fmt.Println("List Movies:")

	for i, movie := range movies {
		movie = fmt.Sprintf("%d. %s", i+1, movie)
		moviesChannel <- movie
	}

	close(moviesChannel)
}

// Soal 3
func luasLingkaran(jariJari []int, luasCh chan string) {
	for _, r := range jariJari {
		luasCh <- fmt.Sprintf("Perhitungan jari-jari %d:\nLuas Lingkaran = %.2f", r, float64(math.Pi)*float64(r)*float64(r))
	}

}

func kelilingLingkaran(jariJari []int, kelilingCh chan string) {
	for _, r := range jariJari {
		kelilingCh <- fmt.Sprintf("Keliling Lingkarang = %.2f", float64(math.Pi)*float64(2)*float64(r))
	}

}

func volumeTabung(jariJari []int, tinggi int, volumeCh chan string) {
	for _, r := range jariJari {
		volumeCh <- fmt.Sprintf("Volume Tabung = %.2f\n", float64(math.Pi)*float64(r)*float64(r)*float64(tinggi))
	}

}

// Soal 4
func luasPersegiPanjang(panjang, lebar int, luasCh chan string) {
	luasCh <- fmt.Sprintf("Luas Persegi Panjang = %d", panjang*lebar)
}

func kelilingPersegiPanjang(panjang, lebar int, kelilingCh chan string) {
	kelilingCh <- fmt.Sprintf("Kelililng Persegi Panjang = %d", 2*(panjang+lebar))
}

func volumeBalok(panjang, lebar, tinggi int, volumeCh chan string) {
	volumeCh <- fmt.Sprintf("Volume Balok = %d", panjang*lebar*tinggi)
}
