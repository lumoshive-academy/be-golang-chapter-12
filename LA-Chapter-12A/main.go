package main

import (
	"fmt"
)

func main() {
	// Membuat Buffer channel dengan kapasitas 2
	ch := make(chan int, 2)
	// Goroutine untuk mengirim data
	go func() {
		fmt.Println("Mengirim data pertama ke channel...")
		ch <- 1 // Mengirim data ke channel
		fmt.Println("Data pertama telah dikirim.")

		fmt.Println("Mengirim data kedua ke channel...")
		ch <- 2 // Mengirim data ke channel
		fmt.Println("Data kedua telah dikirim.")
	}()
	// Menerima data dari channel
	fmt.Println("Menunggu menerima data pertama...")
	value1 := <-ch // Menerima data dari channel
	fmt.Println("Data pertama diterima:", value1)

	fmt.Println("Menunggu menerima data kedua...")
	value2 := <-ch // Menerima data dari channel
	fmt.Println("Data kedua diterima:", value2)
}
