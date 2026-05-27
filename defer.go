package main 

import "fmt"


func Logging() {
	fmt.Println("Selesai Memanggil Function!")
}

func RunApps() {
	defer Logging()
	fmt.Println("Run Application!")
}

func main() {
	RunApps()
}