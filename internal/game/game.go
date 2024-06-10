package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimoehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player     *Player
	background *Background
	keyActions []*KeyAction
}

type KeyAction struct {
	key      ebiten.Key
	callback func(*Player, *Background)
}

func NewKeyActions(player *Player, background *Background) []*KeyAction {
	return []*KeyAction{
		{
			key: ebiten.KeyW,
			callback: func(p *Player, b *Background) {
				p.state = Jumping
				elapsed := time.Since(p.lastFrameTime)
				if elapsed < p.animationDuration/2 {
					p.y += p.speed
				} else {
					p.y -= p.speed
				}
			},
		},
		{
			key: ebiten.KeyA,
			callback: func(p *Player, b *Background) {
				p.state = Walking
				p.orientation = Left
				if p.x <= 10 {
					b.xOffset += p.speed
				} else {
					p.x -= p.speed
				}
			},
		},
		{
			key: ebiten.KeyD,
			callback: func(p *Player, b *Background) {
				p.state = Walking
				p.orientation = Right
				if p.x >= 120 {
					b.xOffset -= p.speed
				} else {
					p.x += p.speed
				}
			},
		},
		{
			key: ebiten.KeySpace,
			callback: func(p *Player, b *Background) {
				p.state = Kicking
			},
		},
	}
}

func NewGame() (*Game, error) {
	g := Game{}
	player := NewPlayer(0, 160.0, Right)
	g.player = player

	backgroundConfig := []BackgroundConfig{
		{
			path:     "../../assets/tree-1.png",
			initialX: 15.0,
			initialY: 0.0,
			scale:    2.0,
		},
		{
			path:     "../../assets/trailer-1.png",
			initialX: 25.0,
			initialY: 40.0,
			scale:    2.0,
		},
	}

	g.background = NewBackground(backgroundConfig, 0.0, 0.0)

	g.keyActions = NewKeyActions(g.player, g.background)

	return &g, nil
}

func (g *Game) Update() error {
	// Call player's update function to manage frame updates

	g.player.UpdateFrame()

	keyPressed := false

	for _, keyAction := range g.keyActions {
		if ebiten.IsKeyPressed(keyAction.key) {
			keyAction.callback(g.player, g.background)
			keyPressed = true
		}
	}

	if !keyPressed {
		g.player.state = Idle
		g.player.currentFrameIndex = 0
	}

	return nil
}

// Draw method for Game struct
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background objects
	for _, obj := range g.background.objects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(obj.x+g.background.xOffset, obj.y+g.background.yOffset)
		op.GeoM.Scale(obj.scale, obj.scale)
		screen.DrawImage(obj.image, op)
	}

	// Draw the player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.x, g.player.y)
	currentFrames := g.player.animationFrames[g.player.state][g.player.orientation]
	if len(currentFrames) > 0 {
		screen.DrawImage(currentFrames[g.player.currentFrameIndex], op)
	}
}

// Layout method for Game struct
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}
