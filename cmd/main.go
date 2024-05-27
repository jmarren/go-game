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
}

type Game struct {
	Images []Image
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
			g.Images[0].X -= 2
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.Images[0].X += 2
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
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
		{Path: "../background.png", X: 0, Y: 0, ScaleX: 6.5, ScaleY: 5, Rotation: 0},
		{Path: "../New Piskel-1.png.png", X: 20, Y: 20, ScaleX: 4, ScaleY: 4, Rotation: 0},
		// Add more images as needed
	}

	// Load images
	for i := range g.Images {
		img, _, err := ebitenutil.NewImageFromFile(g.Images[i].Path)
		if err != nil {
			log.Fatal(err)
		}
		g.Images[i].Img = img
	}

	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Render multiple images with different options")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// var (
// 	background   *ebiten.Image
// 	img          *ebiten.Image
// 	charX, charY float64 = 0, 0 // Character's position
// )
//
// type Image struct {
// 	Img      *ebiten.Image
// 	X, Y     float64
// 	ScaleX   float64
// 	ScaleY   float64
// 	Rotation float64
// }
//
// func init() {
// }
//
// type Game struct {
// 	Images []Image
// }
//
// func updateCharacterPosition() {
// 	if ebiten.IsKeyPressed(ebiten.KeyA) {
// 		charX -= 2
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyD) {
// 		charX += 2
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyW) {
// 		charY -= 2
// 	}
// }
//
// func (g *Game) Draw(screen *ebiten.Image) {
// for _, img := range g.Images {
// 		if img.Img == nil {
// 			continue
// 		}
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(img.X, img.Y)
// 		op.GeoM.Scale(img.ScaleX, img.ScaleY)
// 		op.GeoM.Rotate(img.Rotation)
// 		screen.DrawImage(img.Img, op)
// 	}
// }
//
// func (g *Game) Update() error {
// 	updateCharacterPosition()
// 	return nil
// }
//
// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 640, 480
// }
//
// func main() {
// 	ebiten.SetWindowSize(1280, 960)
// 	ebiten.SetWindowTitle("Render an image")
// 	g := &Game{}
// 	g.Images = loadImages()
// 	if err := ebiten.RunGame(g); err != nil {
// 		log.Fatal(err)
// 	}
// }
//
// func loadImages() []Image {
// 	var images []Image
//
// 	// Load first image
// 	img1, _, err := ebitenutil.NewImageFromFile("../New Piskel-1.png.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	images = append(images, Image{Img: img1, X: 0, Y: 0})
//
// 	// Load second image
// 	img2, _, err := ebitenutil.NewImageFromFile("../background.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	images = append(images, Image{Img: img2, X: 100, Y: 100})
//
// 	// Add more images as needed
// 	// ...
//
// 	return images
// }
