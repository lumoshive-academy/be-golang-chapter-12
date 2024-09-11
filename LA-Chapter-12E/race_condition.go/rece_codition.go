package main

import (
	"fmt"
)

// example race condition
func main() {
	var counter int

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	fmt.Println("Nilai counter:", counter)
}
