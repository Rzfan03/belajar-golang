package main

import (
	"fmt"
	"time"
	)

func main() {
	for {
		waktuSekarang := time.Now()
		formatWaktu := waktuSekarang.Format("15:04:05")

		fmt.Print("\r", formatWaktu)
		time.Sleep(1 * time.Second)
	}
}
