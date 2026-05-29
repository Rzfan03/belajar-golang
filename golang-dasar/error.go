package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("kek mana lahh, mana bisa pake 0 wee!")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(50, 0)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}
