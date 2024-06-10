package game

import (
	"mygame/internal/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	}
}

func (p *Player) UpdateFrame() {
	// Get the frames for the current state and orientation
	frames := p.animationFrames[p.state][p.orientation]

	// If the player is not idle, update the frame index based on elapsed time
	if p.state != Idle {
		elapsed := time.Since(p.lastFrameTime) // convert to milliseconds
		if elapsed > p.animationDuration {
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
