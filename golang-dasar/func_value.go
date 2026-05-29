package main


import "fmt"


func sayHello(say string) string {
	return "Hello " + say
}

func main() {
	halo := sayHello
	fmt.Println(halo("rizfan"))
}
