package game

import (
	"fmt"
	"mygame/internal/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimoehoshi/ebiten/v2/inpututil"
)

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
	isJumpEnabled     bool
}

func NewPlayer(x, y int, orientation Orientation) *Player {
	// Load all frames and organize them by state and orientation
	frames := map[PlayerState]map[Orientation][]*ebiten.Image{
		Walking: {
			Right: assets.LoadFrames([]string{"OldMan-stride-right-1.png", "OldMan-stride-right-2.png"}),
			Left:  assets.LoadFrames([]string{"OldMan-stride-left-1.png", "OldMan-stride-left-2.png"}),
		},
		Kicking: {
			Right: assets.LoadFrames([]string{"OldMan-kick-right-1.png", "OldMan-kick-right-2.png"}),
			Left:  assets.LoadFrames([]string{"OldMan-kick-left-1.png", "OldMan-kick-left-2.png"}),
		},
		Jumping: {
			Right:  assets.LoadFrames([]string{"OldMan-center-jump.png"}),
			Left:   assets.LoadFrames([]string{"OldMan-center-jump.png"}),
			Center: assets.LoadFrames([]string{"OldMan-center-jump.png"}),
		},
		Idle: {
			Right: assets.LoadFrames([]string{"OldMan-facing-right.png"}),
			Left:  assets.LoadFrames([]string{"OldMan-facing-left.png"}),
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
		isJumpEnabled:     true,
	}
}

func (p *Player) UpdateFrame() {
	switch p.state {
	case Walking:
		p.walk()
	case Jumping:
		p.jump()
	case Idle:
		p.idle()
	case Kicking:
		p.kick()
	default:
		p.idle()
	}
}

func (p *Player) walk() {
	elapsed := time.Since(p.lastFrameTime)
	frames := p.animationFrames[Walking][p.orientation]

	if elapsed > p.animationDuration {
		if p.currentFrameIndex == len(frames)-1 {
			p.currentFrameIndex = 0
		} else {
			p.currentFrameIndex++
		}
		p.lastFrameTime = time.Now()
		p.animationDuration = 200 * time.Millisecond
	}
}

func (p *Player) jump() {
	elapsed := time.Since(p.lastFrameTime)
	fmt.Println("\nelapsed: ", elapsed, "\nanimationDuration: ", p.animationDuration, "\ncurrentFrameIndex: ", p.currentFrameIndex, "\nstate: ", p.state)

	if elapsed < 250*time.Millisecond {
		p.y -= p.speed
	} else if elapsed < 500*time.Millisecond {
		p.y += p.speed
	} else if elapsed >= p.animationDuration {
		p.isJumpEnabled = false
		p.state = Idle
		p.currentFrameIndex = 0
		p.lastFrameTime = time.Now()
	}
}

func (p *Player) kick() {
	elapsed := time.Since(p.lastFrameTime)
	frames := p.animationFrames[Jumping][p.orientation]

	if elapsed > p.animationDuration {
		if p.currentFrameIndex == len(frames)-1 {
			p.currentFrameIndex = 0
			p.state = Idle
		} else {
			p.currentFrameIndex++
		}
		// p.lastFrameTime = time.Now()
		p.animationDuration = 100 * time.Millisecond
	}
}

func (p *Player) idle() {
	p.currentFrameIndex = 0
}
