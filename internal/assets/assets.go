package assets // import "github.com/kevinburke/assets"

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadFrames(paths []string) []*ebiten.Image {
	frames := []*ebiten.Image{}
	for _, path := range paths {
		img, error := LoadImage("../../assets/" + path)
		if error != nil {
			panic(error)
		}
		frames = append(frames, img)
	}
	return frames
}

func LoadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}
	return img, nil
}
