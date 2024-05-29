package game

import (
	"mygame/internal/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	images []*ebiten.Image
	X, Y   float64
}

func NewBackground() (*Background, error) {
	images := []*ebiten.Image{}
	image, err := assets.LoadImage("../../assets/trailer-1.png")
	if err != nil {
		return nil, err
	}
	images = append(images, image)
	return &Background{
		images: images,
		X:      30,
		Y:      30,
	}, nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	for _, img := range b.images {
		// Draw the current frame of the background
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2.0, 2.0)
		opts.GeoM.Translate(b.X, b.Y)
		screen.DrawImage(img, opts)
	}
}
