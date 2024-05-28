package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Image represents an image and its drawing options
type Image struct {
	Path     string
	Img      *ebiten.Image
	X, Y     float64
	ScaleX   float64
	ScaleY   float64
	Rotation float64
	ResX     float64
	ResY     float64
}

type Game struct {
	Images     []Image
	Background Image
}

func init() {
	// No need to load images here; loading will be done in Game initialization
}

func (g *Game) Update() error {
	// Update the character's position based on input (if needed)
	updateCharacterPositions(g)
	return nil
}

func updateCharacterPositions(g *Game) {
	// Example of moving the first image with 'A' and 'D' keys
	if len(g.Images) > 0 {
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.Images[1].X -= 2
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.Images[1].X += 2
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	numberOfTilesX := 1280 / 16
	numberOfTilesY := 960 / 16
	for i := 0; i < numberOfTilesX; i++ {
		for j := 0; j < numberOfTilesY; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*16), float64(j*16))
			op.GeoM.Scale(4, 4)
			screen.DrawImage(g.Background.Img, op)
		}
	}
	for _, img := range g.Images {
		if img.Img == nil {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(img.X, img.Y)
		op.GeoM.Scale(img.ScaleX, img.ScaleY)
		op.GeoM.Rotate(img.Rotation)
		screen.DrawImage(img.Img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	g := &Game{}

	// Define the images with their paths and options
	g.Images = []Image{
		{Path: "../assets/OldMan1.png", X: 40, Y: 20, ScaleX: 4, ScaleY: 4, Rotation: 0, ResX: 32, ResY: 32},
		{Path: "../New Piskel-1.png.png", X: 20, Y: 20, ScaleX: 4, ScaleY: 4, Rotation: 0, ResX: 32, ResY: 32},
	}

	// Define the background
	g.Background = Image{
		Path: "../assets/GrayTile1.png", X: 0, Y: 0, ScaleX: 4, ScaleY: 4, Rotation: 0, ResX: 16, ResY: 16,
	}

	// Load images
	for i := range g.Images {
		img, _, err := ebitenutil.NewImageFromFile(g.Images[i].Path)
		if err != nil {
			log.Fatal(err)
		}
		g.Images[i].Img = img
	}

	tile, _, err := ebitenutil.NewImageFromFile(g.Background.Path)
	if err != nil {
		log.Fatal(err)
	}
	g.Background.Img = tile

	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Render multiple images with different options")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
