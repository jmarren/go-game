package game

import (
	"mygame/internal/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

// "github.com/hajimehoshi/ebiten/inpututil"
// "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player     *Player
	background *Background
	keyActions []*KeyAction
	bgColor    colorm.ColorM
}

type PlayerState int

const (
	Idle PlayerState = iota
	Walking
	Kicking
	Jumping
)

type Orientation int

const (
	Right Orientation = iota
	Left
	Center
)

type Player struct {
	x, y              float64
	state             PlayerState
	stateTimer        time.Time
	animationFrames   map[PlayerState]map[Orientation][]*ebiten.Image
	animationDuration time.Duration
	currentFrameIndex int
	orientation       Orientation
	speed             float64
}

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

// type Animation struct {
// 	frames        []*ebiten.Image
// 	frameDuration time.Duration
// }

type KeyAction struct {
	key      ebiten.Key
	frames   []*ebiten.Image
	callback func(*Player, *Background)
}

func NewPlayer(x, y int, orientation Orientation) *Player {
	// Load all frames and organize them by state and orientation
	frames := map[PlayerState]map[Orientation][]*ebiten.Image{
		Walking: {
			Right: loadFrames([]string{"OldMan-stride-right-1.png", "OldMan-stride-right-2.png"}),
			Left:  loadFrames([]string{"OldMan-stride-left-1.png", "OldMan-stride-left-2.png"}),
		},
		Kicking: {
			Right: loadFrames([]string{"OldMan-kick-right-1.png", "OldMan-kick-right-2.png"}),
			Left:  loadFrames([]string{"OldMan-kick-left-1.png", "OldMan-kick-left-2.png"}),
		},
		Jumping: {
			Right:  loadFrames([]string{"OldMan-center-jump.png"}),
			Left:   loadFrames([]string{"OldMan-center-jump.png"}),
			Center: loadFrames([]string{"OldMan-center-jump.png"}),
		},
		Idle: {
			Right: loadFrames([]string{"OldMan-facing-right.png"}),
			Left:  loadFrames([]string{"OldMan-facing-left.png"}),
		},
		// Add more states and orientations as needed
	}

	return &Player{
		x:                 float64(x),
		y:                 float64(y),
		state:             Idle,
		orientation:       orientation,
		animationFrames:   frames,
		animationDuration: 100 * time.Millisecond,
		speed:             2.0,
	}
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

func NewKeyActions(player *Player, background *Background) []*KeyAction {
	return []*KeyAction{
		{
			key: ebiten.KeyW,
			callback: func(p *Player, b *Background) {
				p.state = Jumping
				p.currentFrameIndex = 0
			},
		},
		{
			key: ebiten.KeyA,
			callback: func(p *Player, b *Background) {
				if p.x <= 10 {
					b.xOffset += p.speed
				} else {
					p.x -= p.speed
				}
				p.state = Walking
				p.orientation = Left
				p.currentFrameIndex = 0
			},
		},
		{
			key: ebiten.KeyD,
			callback: func(p *Player, b *Background) {
				p.state = Walking
				p.orientation = Right
				p.currentFrameIndex = 0
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
				p.currentFrameIndex = 0
			},
		},
	}
}

func loadFrames(paths []string) []*ebiten.Image {
	frames := []*ebiten.Image{}
	for _, path := range paths {
		img, error := assets.LoadImage("../../assets/" + path)
		if error != nil {
			panic(error)
		}
		frames = append(frames, img)
	}
	return frames
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
			initialX: 75.0,
			initialY: 40.0,
			scale:    2.0,
		},
	}

	g.background = NewBackground(backgroundConfig, 0.0, 0.0)

	g.keyActions = NewKeyActions(g.player, g.background)

	return &g, nil
}

// Update method for Game struct
func (g *Game) Update() error {
	for _, keyAction := range g.keyActions {
		if ebiten.IsKeyPressed(keyAction.key) {
			keyAction.callback(g.player, g.background)
			// Handle frame update
			if time.Since(g.player.stateTimer) > g.player.animationDuration {
				frames := g.player.animationFrames[g.player.state][g.player.orientation]
				g.player.currentFrameIndex = (g.player.currentFrameIndex + 1) % len(frames)
				g.player.stateTimer = time.Now()
			}
		} else if time.Since(g.player.stateTimer) > g.player.animationDuration {
			g.player.state = Idle
			g.player.currentFrameIndex = 0
			g.player.stateTimer = time.Time{}
		}
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
