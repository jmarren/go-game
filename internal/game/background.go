package game

import (
	"mygame/internal/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	objects []*BackgroundObject
	xOffset float64
	yOffset float64
}

type BackgroundObject struct {
	x     float64
	y     float64
	image *ebiten.Image
	scale float64
}

type BackgroundConfig struct {
	path     string
	initialX float64
	initialY float64
	scale    float64
}

func NewBackground(backgroundConfigs []BackgroundConfig, xOffset float64, yOffset float64) *Background {
	background := Background{}
	for _, config := range backgroundConfigs {
		img, _ := assets.LoadImage(config.path)
		backgroundObject := BackgroundObject{
			x:     config.initialX,
			y:     config.initialY,
			image: img,
			scale: config.scale,
		}
		background.objects = append(background.objects, &backgroundObject)
	}
	background.xOffset = xOffset
	background.yOffset = yOffset
	return &background
}
