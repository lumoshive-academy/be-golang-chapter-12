package main

import (
	"fmt"
)

func main() {
	// Membuat buffered channel dengan kapasitas 3
	ch := make(chan int, 3)
	// Goroutine yang mengirim data ke channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Mengirim", i)
		}
		close(ch) // Menutup channel setelah selesai mengirim
	}()
	// Menggunakan range untuk menerima data dari channel sampai channel ditutup
	for value := range ch {
		fmt.Println("Menerima", value)
	}

	fmt.Println("Semua data telah diterima.")
}
