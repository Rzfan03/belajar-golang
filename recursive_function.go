package main

import "fmt"

func factoricalLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func 

func main() {
	result := 10 * 9 * 8 * 7 * 6 *5 * 4 * 3 * 2 * 1
	fmt.Println(result)
	fmt.Println(factoricalLoop(10))
}