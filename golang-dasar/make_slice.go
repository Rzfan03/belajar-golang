package main


import "fmt"


func main() {
	var NewSlice []string = make([]string, 2, 5)
	NewSlice[0] = "Hello"
	NewSlice[1] = "World!"

	fmt.Println(NewSlice)
	fmt.Println(len(NewSlice))
	fmt.Println(cap(NewSlice))


	NewSlice2 := append(NewSlice, "HAi")
	fmt.Println(NewSlice2)
	fmt.Println(len(NewSlice2))
	fmt.Println(cap(NewSlice2))	


	fromSlice := NewSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}