package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	// Menjalankan 1000 goroutine
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Race condition tetap ada di sini
		}()
	}

	// Menunggu semua goroutine selesai
	wg.Wait()

	fmt.Println("Nilai counter:", counter)
}
