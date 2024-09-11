package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock() // Mengunci mutex sebelum mengakses counter
			counter++
			mutex.Unlock() // Membuka kunci mutex setelah selesai mengakses counter
			wg.Done()
		}()
	}

	wg.Wait() // Menunggu semua goroutine selesai
	fmt.Println("Nilai counter:", counter)
}
