package game

import (
	"mygame/internal/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type BackgroundImage struct {
	Image *ebiten.Image
	X, Y  float64
}

type Background struct {
	images               []*BackgroundImage
	bgOffsetX, bgOffsetY float64
}

func NewBackground(images []*BackgroundImage, initialX float64, initialY float64) (*Background, error) {
	return &Background{
		images:    images,
		bgOffsetX: initialX,
		bgOffsetY: initialY,
	}, nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	for _, img := range b.images {
		// Draw the current frame of the background
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(2.0, 2.0)
		opts.GeoM.Translate(b.bgOffsetX+img.X, b.bgOffsetY+img.Y)
		screen.DrawImage(img.Image, opts)
	}
}

func CreateBackgroundImage(path string, initialX float64, initialY float64) (*BackgroundImage, error) {
	image, loadErr := assets.LoadImage(path)
	if loadErr != nil {
		return nil, loadErr
	}
	return &BackgroundImage{
		Image: image,
		X:     initialX,
		Y:     initialY,
	}, nil
}

func (b *Background) MoveBackground(deltaX float64, deltaY float64) {
	b.bgOffsetX += deltaX
	b.bgOffsetY += deltaY
}
