package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// players    []*Player
	player     *Player
	background *Background
	keyActions []KeyAction
}

type Location struct {
	X float64
	Y float32
}

type KeyAction struct {
	key         ebiten.Key
	handler     func(*Player, *Background)
	actionState ActionState
	orientation Orientation
}

type BackgroundObjectInit struct {
	path     string
	initialX float64
	initialY float64
}

type BackgroundObject struct {
	Image *ebiten.Image
	X     float64
	Y     float64
}

type Background struct {
	Objects []*BackgroundObject
	OffsetX float64
	OffsetY float64
}
type Frame struct {
	path          string
	frameDuration int
}

type PlayerInit struct {
	PlayerStates []PlayerState
	// frames   []Frame
	initialX float64
	initialY float64
}

type Player struct {
	// frames        map[Orientation]map[PlayerState][]*ebiten.Image
	actionState       ActionState
	currentFrame      int // current frame index
	animationDuration time.Duration
	lastFrameTime     time.Time
	X, Y              float64
	speed             float64
	orientation       Orientation
	state             PlayerState
	stateTimer        time.Time
	// stateDuration map[PlayerState]time.Duration
}

type Orientation int

const (
	left Orientation = iota
	right
	center
)

type ActionState int

const (
	idle ActionState = iota
	kick
	walk
	jump
)

// const (
// 	idle ActionState = iota
// 	kick1
// 	kick2
// 	stride1
// 	stride2
// 	jump
// )

// type AnimationName int

// const AnimationNames (
//
//	  kick1 AnimationName = iota
//	  kick2
//	  walk1
//	)
type PlayerState struct {
	ActionState     ActionState
	Orientation     Orientation
	AnimationFrames []Frame
	speed           float64
	duration        int
}

type Animation struct {
	name        ActionState
	orientation Orientation
	frames      []Frame
	duration    int
}

type Animations struct {
	name       Animation
	animations []Animation
}

func (a *Animations) getAnimation(name ActionState, orientation Orientation, animations []*Animations) Animation {
	for _, animation := range animations {
		if animation.name == name && animation.orientation == orientation {
			return animation
		}
	}
	return nil
}

func NewGame() (*Game, error) {
	// Initialize Background
	backgroundInit := []BackgroundObjectInit{
		{
			path:     "../../assets/tree-1.png",
			initialX: 30.0,
			initialY: -30.0,
		},
		{
			path:     "../../assets/trailer-1.png",
			initialX: 130.0,
			initialY: 40.0,
		},
	}

	backgroundObjects := []*BackgroundObject{}

	animations := []Animation{
		{
			name:        kick,
			orientation: left,
			frames: []Frame{
				{
					path:          "OldMan-kick-left-1.png",
					frameDuration: 1,
				},
				{
					path:          "OldMan-kick-left-2.png",
					frameDuration: 1,
				},
			},
			duration: 2,
		},
		{
			name:        kick,
			orientation: right,
			frames: []Frame{
				{
					path:          "OldMan-kick-right-1.png",
					frameDuration: 1,
				},
				{
					path:          "OldMan-kick-right-2.png",
					frameDuration: 1,
				},
			},
		},
		{
			name:        walk,
			orientation: left,
			frames: []Frame{
				{
					path:          "OldMan-stride-left-1.png",
					frameDuration: 1,
				},
				{
					path:          "OldMan-stride-left-2.png",
					frameDuration: 1,
				},
			},
			duration: 2,
		},
		{
			name:        walk,
			orientation: right,
			frames: []Frame{
				{
					path:          "OldMan-stride-right-1.png",
					frameDuration: 1,
				},
				{
					path:          "OldMan-stride-right-2.png",
					frameDuration: 1,
				},
			},
			duration: 2,
		},
	}

	keyActions := []KeyAction{
		{
			key: ebiten.KeyA,
			handler: func(p *Player, b *Background) {
				if p.X >= 20 {
					p.X -= p.speed
				} else {
					b.OffsetX -= p.speed
				}
			},
			actionState: walk,
			orientation: left,
		},
		{
			key: ebiten.KeyD,
			handler: func(p *Player, b *Background) {
				if p.X <= 275 {
					p.X += p.speed
				} else {
					b.OffsetX += p.speed
				}
			},
			actionState: walk,
			orientation: right,
		},
		{
			key: ebiten.KeyW,
			handler: func(p *Player, b *Background) {
			},
			actionState: jump,
		},
		{
			key: ebiten.KeySpace,
			handler: func(p *Player, b *Background) {
			},
			actionState: kick,
		},
	}

	// return &Game{
	// background: &Background{
}

func (g *Game) Update() error {
	for _, keyAction := range g.keyActions {
		if inpututil.IsKeyJustPressed(keyAction.key) {
			keyAction.handler(g.player, g.background)
			if time.Since(g.player.stateTimer) > g.player.animationDuration {
				g.player.stateTimer = time.Now()
				g.player.currentFrame = (g.player.currentFrame + 1) % len(getAnimation(keyAction.ActionState, keyAction.orientation).frames)
				g.player.animationDuration = getAnimation(keyAction.ActionState, keyAction.orientation).duration
				g.player.actionState = keyAction.actionState
				g.player.orientation = keyAction.orientation
			}
		}
	}
	// Update players
	// Update background
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	//
	// Draw players
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func CreateBackgroundObject(path string, initialX, initialY float64) (*BackgroundObject, error) {
	img, err := LoadImage(path)
	if err != nil {
		return nil, err
	}
	return &BackgroundObject{
		Image: img,
		X:     initialX,
		Y:     initialY,
	}, nil
}

func CreateBackgroundObjects(backgroundObjectInits []*BackgroundObjectInit) ([]*BackgroundObject, error) {
	backgroundObjects := []*BackgroundObject{}
	for _, object := range backgroundObjectInits {
		backgroundObject := CreateBackgroundObject(object.path, object.initialX, object.initialY)
		if err != nil {
			return nil, err
		} else {
			backgroundObjects = append(backgroundObjects, backgroundObject)
		}
	}
	return &backgroundObjects
}

func CreateBackground(backgroundObjects []*BackgroundObject, offsetX, offsetY float64) (*Background, error) {
	return &Background{
		Objects: backgroundObjects,
		OffsetX: offsetX,
		OffsetY: offsetY,
	}, nil
}

func CreateKeyAction(key ebiten.Key, handler func(*Player, *Background), actionState ActionState, orientation Orientation) *KeyAction {
	return &KeyAction{
		key:         key,
		handler:     handler,
		actionState: actionState,
		orientation: orientation,
	}
}

func CreateNewAnimation(name ActionState, orientation Orientation, frames []Frame, duration int) *Animation {
	return &Animation{
		name:        name,
		orientation: orientation,
		frames:      frames,
		duration:    duration,
	}
}

func CreateAnimations(animations []Animation) *Animations {
}

// func loadBackgroundObjects(backgroundObjects []*BackgroundObjectInit) error {
//   for _, object := range backgroundObjects {
//     img, err := LoadImage(object.path)
//     if err != nil {
//       return err
//     }
//

// Maybe I should use a map instead of a slice here
// Maybe I should use a keypress event for each frame as well

// playerInit := PlayerInit{
//   PlayerStates: []PlayerState{
//     {
//       ActionState: kick1,
//       Orientation: left,
//       AnimationFrames: []Frame{
//         path:    "OldMan-kick-left-1.png",
//         duration: 1,
//       }
//     },
//     {
//       ActionState: kick2,
//       Orientation: left,
//       AnimationFrames: []Frame{
//         path:    "OldMan-kick-left-2.png",
//         duration: 1,
//       },
//     },
//     {
//       ActionState: stride1,
//       Orientation: left,
//       AnimationFrames: []Frame{
//         path:    "OldMan-stride-left-1.png",
//         duration: 1,
//       },
//     }
//
//
// playerInit := PlayerInit{
// 	frames: []Frame{
// 		{
// 			path:     "OldMan-facing-left.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.X -= p.speed
// 			},
// 		},
// 		{
// 			path:     "OldMan-facing-right.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.X += p.speed
// 			},
// 		},
// 		{
// 			path:     "OldMan-kick-left-1.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(kick1)
// 			},
// 		},
// 		{
// 			path:     "OldMan-kick-right-1.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(kick1)
// 			},
// 		},
// 		{
// 			path:     "OldMan-kick-left-2.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(kick2)
// 			},
// 		},
// 		{
// 			path:     "OldMan-kick-right-2.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(kick2)
// 			},
// 		},
// 		{
// 			path:     "OldMan-stride-left-1.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(stride1)
// 			},
// 		},
// 		{
// 			path:     "OldMan-stride-right-1.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(stride1)
// 			},
// 		},
// 		{
// 			path:     "OldMan-center-jump.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetState(jump)
// 			},
// 		},
// 		{
// 			path:     "OldMan-stride-left-2.png",
// 			duration: 1,
// 			action: func(p *Player) {
// 				p.SetOrientation(left)
// 			},
// 		},
//      {
//        path:     "OldMan-stride-right-2.png",
//        duration: 1,
//        action: func(p *Player) {
//          p.SetOrientation(right)
//          p.setState(stride2)
//       },
//      },
//      {
//        path:     "OldMan-kick-left-1.png",
//        duration: 1,
//        action: func(p *Player) {
//          p.SetState(kick1)
//        },
//      },
//      {
//        path:     "OldMan-kick-right-1.png",
//        duration: 1,
//        action: func(p *Player) {
//          p.SetState(kick1)
//        },
//      },
//      {
//        path:     "OldMan-kick-left-2.png",
//
//      }
//
//      }
// 	},
// }

// Initialize Players
//
// 	return &Game{}, nil
// }

// if ebiten.IsKeyPressed(ebiten.KeyD) {
// 	p.SetOrientation(right)
// 	p.Walk()
// 	if p.X <= 275 {
// 		p.X += p.speed
// 	}
// }
//
// if ebiten.IsKeyPressed(ebiten.KeyA) {
// 	p.SetOrientation(left)
// 	p.Walk()
// 	if p.X >= 20 {
// 		p.X -= p.speed
// 	} else {
// 		CameraLeft()
// 	}
// }
// if ebiten.IsKeyPressed(ebiten.KeyW) {
// 	p.state = jump
// 	p.SetOrientation(center)
// 	p.Y -= p.speed
// }
// if ebiten.IsKeyPressed(ebiten.KeyS) {
// 	p.Y += p.speed
// }
//
// if ebiten.IsKeyPressed(ebiten.KeySpace) {
// 	p.SetState(kick1)
// }

// func NewGame() (*Game, error) {
// 	oldManFrames := []string{"OldMan-facing-left.png", "OldMan-facing-right.png", "OldMan-kick-left-1.png", "OldMan-kick-right-1.png"}
// 	players := []*Player{}
//
// 	backgroundInit := []BackgroundImageInit{
// 		{
// 			path:     "../../assets/tree-1.png",
// 			initialX: 30.0,
// 			initialY: -30.0,
// 		},
// 		{
// 			path:     "../../assets/trailer-1.png",
// 			initialX: 130.0,
// 			initialY: 40.0,
// 		},
// 	}
//
// 	backgroundImages := []*BackgroundImage{}
//
// 	for _, image := range backgroundInit {
// 		image, err := CreateBackgroundImage(image.path, image.initialX, image.initialY)
// 		if err != nil {
// 			return nil, err
// 		}
// 		backgroundImages = append(backgroundImages, image)
//
// 	}
//
// 	background, backgroundErr := NewBackground(backgroundImages, 30, 30)
// 	if backgroundErr != nil {
// 		return nil, backgroundErr
// 	}
// 	oldMan, err := NewPlayer(oldManFrames)
// 	players = append(players, oldMan)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &Game{
// 		players:    players,
// 		background: background,
// 	}, nil
// }
//
// func (g *Game) Update() error {
// 	for _, player := range g.players {
// 		player.Update()
// 	}
// 	return nil
// }
//
// func (g *Game) Draw(screen *ebiten.Image) {
// 	screen.Fill(color.RGBA{255, 255, 255, 255})
// 	g.background.Draw(screen)
// 	for _, player := range g.players {
// 		player.Draw(screen)
// 	}
// }
//
// func (g *Game) CameraLeft() {
// 	g.background.MoveBackground(-1, 0)
// }
//
// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 320, 240
// }
