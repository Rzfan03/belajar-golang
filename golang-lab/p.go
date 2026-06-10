package main

import "fmt"

type Game struct {
	Name string
	Developer string
}

func (g *Game) GameInfo(l string) {
	fmt.Printf("Launcher : %s \nGame Name : %s \nDeveloper : %s\n\n",l, g.Name, g.Developer)
}

func main() {
	game := Game{
		Name: "Devil May Cry",
		Developer: "Capcom",
	}

	game2 := &game
	game2.Name = "Free Fire"
	game2.Developer = "Garena"

	game2.GameInfo("Play Store")
	game.GameInfo("Steam")
	fmt.Println(&game2)
}