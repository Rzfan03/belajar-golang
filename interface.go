package main

import "fmt"

type Speaker interface {
	getSound() string
}

func SoundCheck(s Speaker) {
	fmt.Println(s.getSound())
}

type speakers struct {
	sound string
}

func (s speakers) getSound() string {
	return "Bunyi Speaker : " + s.sound 
}

func main() {
	toshiba := speakers{sound: "kaboom!"}
	lg := speakers{sound: "hee heee..."}
	SoundCheck(toshiba)
	SoundCheck(lg)
}