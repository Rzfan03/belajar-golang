package main

import "fmt"

func random() any {
	return 1200 // "rzfan"
}

func main() {
	result := random()


	// pake switch
	switch value := result.(type) {
		case string: 
		fmt.Println("string", value)
		case int:
		fmt.Println("Int", value)
		default:
		fmt.Println("Au ah!")
	} 

	
	// fmt.Println(result)
	// fmt.Println(resultString)
}