package main

import "fmt"

func main() {
	var nama string
	fmt.Println("Enter ur name: ")
	fmt.Scan(&nama)

	switch nama {
		case "Rzfan03":
		fmt.Println("OK!")
		case "aiman": 
		fmt.Println("who?")
	}
}