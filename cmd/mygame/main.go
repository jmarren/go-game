package main

import (
	_ "image/png"
	"log"
	"mygame/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Kicking Animation Example")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
