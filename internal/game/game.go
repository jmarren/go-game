package game

import (
	"mygame/internal/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// "github.com/hajimehoshi/ebiten/inpututil"
// "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player     *Player
	background *Background
	keyActions []*KeyAction
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
	animationFrames   map[PlayerState]map[Orientation][]*ebiten.Image
	animationDuration time.Duration
	lastFrameTime     time.Time
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
		currentFrameIndex: 0,
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
				// p.animationDuration = 200 * time.Millisecond
			},
		},
		{
			key: ebiten.KeyA,
			callback: func(p *Player, b *Background) {
				p.state = Walking
				p.orientation = Left
				// p.frameDuration = 100 * time.Millisecond
				// p.animationDuration = 200 * time.Millisecond
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
				// p.currentFrameIndex = 0
				// p.frameDuration = 100 * time.Millisecond
				// p.animationDuration = 200 * time.Millisecond
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
				// p.currentFrameIndex = 0
				// p.frameDuration = 200 * time.Millisecond
				// p.animationDuration = 200 * time.Millisecond
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

	// Handle input and other game logic
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

func (p *Player) UpdateFrame() {
	// Get the frames for the current state and orientation
	frames := p.animationFrames[p.state][p.orientation]
	//
	// if p.state == Walking {
	// 	fmt.Printf("\nFrame Index: %d \nLast Frame Time: %s\nElapsed Time: %s\nAnimation Duration: %s\nElapsed > Animation Duration ?: %t", p.currentFrameIndex, p.lastFrameTime, time.Since(p.lastFrameTime), p.animationDuration, time.Since(p.lastFrameTime) > p.animationDuration)
	// }

	// If the player is not idle, update the frame index based on elapsed time
	if p.state != Idle {
		// now := time.Now()
		elapsed := time.Since(p.lastFrameTime) // convert to milliseconds
		if elapsed > p.animationDuration {
			// p.currentFrameIndex = (p.currentFrameIndex + 1) % len(frames)
			if p.currentFrameIndex == len(frames)-1 {
				p.currentFrameIndex = 0
			} else {
				p.currentFrameIndex++
			}
			p.lastFrameTime = time.Now()
			p.animationDuration = 200 * time.Millisecond
		}
		// If we reach the end of the frames, handle state transition if necessary
		if p.currentFrameIndex == len(frames)-1 && p.state != Walking {
			p.state = Idle // Transition to Idle or another state as needed
			p.currentFrameIndex = 0
		}

	} else {
		// Reset to the first frame if idle
		p.currentFrameIndex = 0
	}
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
