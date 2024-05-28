package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	players []*Player
}

func NewGame() (*Game, error) {
	oldManFrames := []string{"OldMan-facing-left.png", "OldMan-facing-right.png", "OldMan-kick-left-1.png", "OldMan-kick-right-1.png"}
	players := []*Player{}

	oldMan, err := NewPlayer(oldManFrames)
	players = append(players, oldMan)
	if err != nil {
		log.Fatal(err)
	}
	return &Game{
		players: players,
	}, nil
}

func (g *Game) Update() error {
	// Update player
	for _, player := range g.players {
		player.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// clear screen
	screen.Fill(color.RGBA{255, 255, 255, 255})

	// Draw players
	for _, player := range g.players {
		player.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}