package sprite

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	// Name of the frames in playback order
	Frames []string

	// Number of ticks before changing frames
	TicksPerFrame int

	// Whether the animation should run in a loop
	Loop bool
}

type Manager struct {
	sheet      *ebiten.Image
	frames     map[string]*ebiten.Image
	animations map[string]Animation
}

func NewManager(sheet *ebiten.Image) *Manager {
	if sheet == nil {
		panic("sprite: spritesheet cannot be a nil")
	}

	return &Manager{
		sheet:      sheet,
		frames:     make(map[string]*ebiten.Image),
		animations: make(map[string]Animation),
	}
}

func (m *Manager) AddFrame(
	name string,
	x int,
	y int,
	width int,
	height int,
) error {
	if name == "" {
		return fmt.Errorf("sprite: frame name cannot be empty")
	}

	if width <= 0 || height <= 0 {
		return fmt.Errorf("sprite: frame %q has an invalid size", name)
	}

	rect := image.Rect(x, y, x+width, y+height)
	bounds := m.sheet.Bounds()

	if rect.Min.X < bounds.Min.X || rect.Min.Y < bounds.Min.Y || rect.Max.X > bounds.Max.X || rect.Max.Y > bounds.Max.Y {
		return fmt.Errorf(
			"sprite: frame %q is outside the spritesheet",
			name,
		)
	}

	frame := m.sheet.SubImage(rect).(*ebiten.Image)
	m.frames[name] = frame

	return nil
}

func (m *Manager) AddGridRow(
	prefix string,
	row int,
	firstColumn int,
	count int,
	frameWidth int,
	frameHeight int,
) error {
	if count <= 0 {
		return fmt.Errorf("sprite: frame count must be greater than zero")
	}

	for index := 0; index < count; index++ {
		name := fmt.Sprintf("%s_%d", prefix, index)

		column := firstColumn + index
		x := column * frameWidth
		y := row * frameHeight

		if err := m.AddFrame(
			name,
			x,
			y,
			frameWidth,
			frameHeight,
		); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) AddAnimation(
	name string,
	frameNames []string,
	ticksPerFrame int,
	loop bool,
) error {
	if name == "" {
		return fmt.Errorf("sprite: animation name cannot be empty")
	}

	if len(frameNames) == 0 {
		return fmt.Errorf("sprite: animation %q must contain at least one frame", name)
	}

	if ticksPerFrame <= 0 {
		return fmt.Errorf("sprite: animation %q must have a positive frame duration", name)
	}

	for _, frameName := range frameNames {
		if _, exists := m.frames[frameName]; !exists {
			return fmt.Errorf("sprite: frame %q does not exist", frameName)
		}
	}

	framesCopy := append([]string(nil), frameNames...)

	m.animations[name] = Animation{
		Frames:        framesCopy,
		TicksPerFrame: ticksPerFrame,
		Loop:          loop,
	}
	return nil
}

type Animator struct {
	manager *Manager

	animationName string
	frameIndex    int
	elapsedTicks  int
	finished      bool
}

func NewAnimator(
	manager *Manager,
	initialAnimation string,
) (*Animator, error) {
	if manager == nil {
		return nil, fmt.Errorf("sprite: manager cannot be nil")
	}

	animator := &Animator{
		manager: manager,
	}

	if err := animator.Restart(initialAnimation); err != nil {
		return nil, err
	}

	return animator, nil
}

// Play changes the animation only when it is different from the current one.
//
// This means Play("walk") can safely be called every Update without restarting
// the walk animation every frame.
func (a *Animator) Play(name string) error {
	if a.animationName == name {
		return nil
	}

	return a.Restart(name)
}

// Stop the current animation
func (a *Animator) Stop(name string) error {
	if _, exists := a.manager.animations[name]; !exists {
		return fmt.Errorf(
			"sprite: animation %q does not exist",
			name,
		)
	}
	a.animationName = name
	a.frameIndex = 0
	a.elapsedTicks = 0
	a.finished = true
	return nil
}

// Restart starts an animation from its first frame.
func (a *Animator) Restart(name string) error {
	if _, exists := a.manager.animations[name]; !exists {
		return fmt.Errorf(
			"sprite: animation %q does not exist",
			name,
		)
	}

	a.animationName = name
	a.frameIndex = 0
	a.elapsedTicks = 0
	a.finished = false

	return nil
}

func (a *Animator) Update() {
	animation, exists := a.manager.animations[a.animationName]
	if !exists || a.finished {
		return
	}

	a.elapsedTicks++

	if a.elapsedTicks < animation.TicksPerFrame {
		return
	}

	a.elapsedTicks = 0
	a.frameIndex++

	if a.frameIndex < len(animation.Frames) {
		return
	}

	if animation.Loop {
		a.frameIndex = 0
		return
	}

	a.frameIndex = len(animation.Frames) - 1
	a.finished = true
}

func (a *Animator) CurrentFrame() *ebiten.Image {
	animation, exists := a.manager.animations[a.animationName]
	if !exists || len(animation.Frames) == 0 {
		return nil
	}

	frameName := animation.Frames[a.frameIndex]

	return a.manager.frames[frameName]
}

func (a *Animator) AnimationName() string {
	return a.animationName
}

func (a *Animator) Finished() bool {
	return a.finished
}

// Draw draws the current frame.
//
// x and y represent the upper-left corner of the character.
func (a *Animator) Draw(
	screen *ebiten.Image,
	x float64,
	y float64,
	scale float64,
	flipX bool,
) {
	frame := a.CurrentFrame()
	if frame == nil {
		return
	}

	if scale <= 0 {
		scale = 1
	}

	options := &ebiten.DrawImageOptions{}

	if flipX {
		options.GeoM.Scale(-scale, scale)

		// Scaling negatively moves the image to the left.
		// This translation moves it back to the expected position.
		options.GeoM.Translate(
			float64(frame.Bounds().Dx())*scale,
			0,
		)
	} else {
		options.GeoM.Scale(scale, scale)
	}

	options.GeoM.Translate(x, y)

	screen.DrawImage(frame, options)
}
