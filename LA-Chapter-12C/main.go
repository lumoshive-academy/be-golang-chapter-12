package main

import (
	"fmt"
	"time"
)

func fetchFromAPI1(ch chan string) {
	time.Sleep(2 * time.Second) // Simulasi API lambat
	ch <- "Data dari API 1"
}

func fetchFromAPI2(ch chan string) {
	time.Sleep(1 * time.Second) // Simulasi API lebih cepat
	ch <- "Data dari API 2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go fetchFromAPI1(ch1)
	go fetchFromAPI2(ch2)

	// select statement
	select {
	case data := <-ch1:
		fmt.Println("Menerima:", data)
	case data := <-ch2:
		fmt.Println("Menerima:", data)
	}

	ch3 := make(chan string)
	ch4 := make(chan string)

	// Goroutine untuk mengirim data ke channel 1 setelah 2 detik
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "Data dari channel 1"
	}()
	// Goroutine untuk mengirim data ke channel 2 setelah 1 detik
	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "Data dari channel 2"
	}()

	// Menggunakan select untuk menangani kedua channel
	// default statement
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch3:
			fmt.Println("Menerima dari channel 1:", msg1)
		case msg2 := <-ch4:
			fmt.Println("Menerima dari channel 2:", msg2)
		default:
			fmt.Println("Tidak ada data yang siap untuk diterima.")
			time.Sleep(500 * time.Millisecond) // Simulasi operasi lain
		}
	}

}
