package main


import "fmt"

func main() {
	food := []string{"jeruk", "blueberry", "nanas", "apel", "mangga"}

	for a, i := range food {
		fmt.Println("perulangan ke", a, "dan nama buah : ", i)
	}
	
}
