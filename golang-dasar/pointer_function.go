package main

import "fmt"

type Motor struct {
	motor string
}

func ChangeMotoToKawasaki(moto *Motor) {
	moto.motor = "Kawasaki"
}

func main() {
	moto := Motor{}
	ChangeMotoToKawasaki(&moto)
	fmt.Println(moto)
}