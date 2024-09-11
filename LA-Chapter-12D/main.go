package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Menandakan bahwa goroutine selesai saat keluar
	fmt.Printf("Worker %d sedang bekerja\n", id)
	// Melakukan pekerjaan yang sesuai di sini
}

func main() {
	// Mendapatkan jumlah CPU yang tersedia
	numCPU := runtime.NumCPU()
	fmt.Printf("Jumlah CPU: %d\n", numCPU)

	// Mengatur GOMAXPROCS menjadi 2
	runtime.GOMAXPROCS(4)
	// Contoh penggunaan goroutine
	for i := 0; i < 5; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("Selesai menjalankan goroutine")
		}()
	}

	// Menunggu goroutine selesai
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai menjalankan program")

	// wait group
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Menambahkan jumlah goroutine yang sedang berjalan
		go worker(i, &wg)
	}

	wg.Wait() // Menunggu semua goroutine selesai
	fmt.Println("Semua pekerjaan selesai")

}
