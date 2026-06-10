package main 


import "fmt"

type Speaker struct {
	Brand string
}

func (s Speaker) CheckSound() {
	fmt.Println("Brand Speaker adalah : ", s.Brand)
}

func main() {
	result := Speaker{
		Brand: "Yamaha",
	}
	result.CheckSound()
	fmt.Println(result)
}
