package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	players    []*Player
	background *Background
}

type Location struct {
	X float64
	Y float32
}

type BackgroundImageInit struct {
	path     string
	initialX float64
	initialY float64
}

func NewGame() (*Game, error) {
	oldManFrames := []string{"OldMan-facing-left.png", "OldMan-facing-right.png", "OldMan-kick-left-1.png", "OldMan-kick-right-1.png"}
	players := []*Player{}

	backgroundInit := []BackgroundImageInit{
		{
			path:     "../../assets/tree-1.png",
			initialX: 30.0,
			initialY: -30.0,
		},
		{
			path:     "../../assets/trailer-1.png",
			initialX: 130.0,
			initialY: 40.0,
		},
	}

	backgroundImages := []*BackgroundImage{}

	for _, image := range backgroundInit {
		image, err := CreateBackgroundImage(image.path, image.initialX, image.initialY)
		if err != nil {
			return nil, err
		}
		backgroundImages = append(backgroundImages, image)

	}

	background, backgroundErr := NewBackground(backgroundImages, 30, 30)
	if backgroundErr != nil {
		return nil, backgroundErr
	}
	oldMan, err := NewPlayer(oldManFrames)
	players = append(players, oldMan)
	if err != nil {
		log.Fatal(err)
	}
	return &Game{
		players:    players,
		background: background,
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

	g.background.Draw(screen)
	// Draw players
	for _, player := range g.players {
		player.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
