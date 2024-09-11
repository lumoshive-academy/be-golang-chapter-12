package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Println("Current time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())
	fmt.Println("Location:", now.Location())

	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted time:", formattedTime)

	// time parser
	layout := "2006-01-02 15:04:05"
	str := "2023-06-23 14:30:00"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", parsedTime)

	// Format tanggal dan waktu lengkap
	layout1 := "2006-01-02 15:04:05"
	formattedTime1 := now.Format(layout1)
	fmt.Println("Formatted time 1:", formattedTime1)
	// Format hanya tanggal
	layout2 := "02-Jan-2006"
	formattedTime2 := now.Format(layout2)
	fmt.Println("Formatted time 2:", formattedTime2)

	// Format dengan nama hari dan bulan
	layout3 := "Monday, 02-Jan-2006 03:04:05 PM"
	formattedTime3 := now.Format(layout3)
	fmt.Println("Formatted time 3:", formattedTime3)

	// Format dengan zona waktu
	layout4 := "2006-01-02 15:04:05 MST"
	formattedTime4 := now.Format(layout4)
	fmt.Println("Formatted time 4:", formattedTime4)

	// start time sleep
	fmt.Println("Start")

	// Menunda eksekusi selama 2 detik
	time.Sleep(2 * time.Second)

	fmt.Println("2 seconds later")

	// Menunda eksekusi selama 500 milidetik (0.5 detik)
	time.Sleep(500 * time.Millisecond)

	fmt.Println("0.5 seconds later")

	// time ticker
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("Main function exiting")

	// time after
	fmt.Println("Start")

	// Membuat channel yang mengirimkan waktu tunggal setelah 2 detik
	timeout := time.After(2 * time.Second)

	// Menunggu nilai dari channel timeout atau done
	select {
	case <-timeout:
		fmt.Println("Timeout occurred")
	case <-time.After(3 * time.Second):
		fmt.Println("Operation took too long")
	}

	fmt.Println("End")

	// time after function
	fmt.Println("Start")

	// Menjadwalkan eksekusi fungsi callback setelah 2 detik
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Callback executed after 2 seconds")
	})

	// Menunggu agar callback dieksekusi sebelum keluar
	time.Sleep(3 * time.Second)
	fmt.Println("End")

	// time new timer
	fmt.Println("Start")

	// Membuat timer yang mengirimkan waktu setelah 2 detik
	timer := time.NewTimer(2 * time.Second)

	// Menunggu nilai dari channel timer.C
	<-timer.C
	fmt.Println("Timer expired")

	fmt.Println("End")
}
