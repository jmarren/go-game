package game

import (
	"errors"
	"log"
	"mygame/internal/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Orientation int

const (
	left Orientation = iota
	right
	center
)

type PlayerState int

const (
	idle PlayerState = iota
	kick1
	kick2
)

type Player struct {
	frames        map[Orientation]map[PlayerState][]*ebiten.Image
	currentFrame  int // current frame index
	frameDuration time.Duration
	lastFrameTime time.Time
	X, Y          float64
	speed         float64
	orientation   Orientation
	state         PlayerState
	stateTimer    time.Time
	stateDuration map[PlayerState]time.Duration
}

// Example function to change player orientation
func (p *Player) SetOrientation(o Orientation) error {
	if o > 2 || o < 0 {
		log.Fatal("Invalid orientation")
		return errors.New("Invalid orientation")
	}
	p.orientation = o
	return nil
}

func (p *Player) SetState(s PlayerState) error {
	if s > 2 || s < 0 {
		log.Fatal("Invalid state")
		return errors.New("Invalid state")
	}
	p.state = s
	return nil
}

func NewPlayer(frameFiles []string) (*Player, error) {
	frames := make(map[Orientation]map[PlayerState][]*ebiten.Image)

	for _, orientation := range []Orientation{left, right, center} {
		frames[orientation] = make(map[PlayerState][]*ebiten.Image)
		for _, state := range []PlayerState{idle, kick1, kick2} {
			frames[orientation][state] = []*ebiten.Image{}
		}
	}

	// Load images for each state and orientation
	images := map[Orientation]map[PlayerState][]string{
		left: {
			idle:  {"../../assets/OldMan-facing-left.png"},
			kick1: {"../../assets/OldMan-kick-left-1.png"},
			kick2: {"../../assets/OldMan-kick-left-2.png"},
		},
		right: {
			idle:  {"../../assets/OldMan-facing-right.png"},
			kick1: {"../../assets/OldMan-kick-right-1.png"},
			kick2: {"../../assets/OldMan-kick-right-2.png"},
		},
		// center: {
		// 	idle: {"../../assets/OldMan-center.png"},
		// },
	}

	// Load images into frames map
	for orientation, states := range images {
		for state, files := range states {
			for _, file := range files {
				img, err := assets.LoadImage(file)
				if err != nil {
					return nil, err
				}
				frames[orientation][state] = append(frames[orientation][state], img)
			}
		}
	}
	return &Player{
		frames:        frames,
		currentFrame:  0,
		frameDuration: 100 * time.Millisecond,
		lastFrameTime: time.Now(),
		X:             100,
		Y:             100,
		orientation:   left,
		state:         idle,
		speed:         4,
		stateTimer:    time.Now(),
		stateDuration: map[PlayerState]time.Duration{
			idle:  1 * time.Second,        // Duration for idle state
			kick1: 500 * time.Millisecond, // Duration for kick1 state
			kick2: 500 * time.Millisecond, // Duration for kick2 state
		},
	}, nil
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		err := p.SetOrientation(right)
		if err != nil {
			log.Fatal(err)
		}
		p.X += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		err := p.SetOrientation(left)
		if err != nil {
			log.Fatal(err)
		}
		p.X -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		err := p.SetState(kick1)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Handle state transitions based on state timer
	if time.Since(p.stateTimer) > p.stateDuration[p.state] {
		switch p.state {
		case kick1:
			err := p.SetState(kick2)
			if err != nil {
				log.Fatal(err)
			}
		case kick2:
			err := p.SetState(idle)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Update animation frame based on time and current state
	if time.Since(p.lastFrameTime) > p.frameDuration {
		p.currentFrame++
		if p.currentFrame >= len(p.frames[p.orientation][p.state]) {
			p.currentFrame = 0
		}
		p.lastFrameTime = time.Now()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	img := p.frames[p.orientation][p.state][p.currentFrame]

	// Draw the current frame of the player
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2.0, 2.0)
	opts.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(img, opts)
}
