package player

import (
	"bytes"
	_ "embed"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"sidneiwill.dev/BaseRPGMovement/sprite"
)

type Direction int

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

type Player struct {
	Animator  *sprite.Animator
	Direction Direction
	moving    bool
}

func NewAnimator(spriteSheet []byte) (*Player, error) {
	sheet, _, err := ebitenutil.NewImageFromReader(
		bytes.NewReader(spriteSheet),
	)
	if err != nil {
		return nil, err
	}

	manager := sprite.NewManager(sheet)

	// Row 1: walk_0 through walk_5.
	if err := manager.AddGridRow(
		"up",
		1,
		0,
		4,
		48,
		48,
	); err != nil {
		return nil, err
	}

	if err := manager.AddGridRow(
		"down",
		0,
		0,
		4,
		48,
		48,
	); err != nil {
		return nil, err
	}

	if err := manager.AddGridRow(
		"right",
		3,
		0,
		4,
		48,
		48,
	); err != nil {
		return nil, err
	}

	if err := manager.AddAnimation(
		"walk_up",
		[]string{
			"up_2",
			"up_3",
			"up_0",
			"up_1",
		},
		6,
		true,
	); err != nil {
		return nil, err
	}

	if err := manager.AddAnimation(
		"walk_down",
		[]string{
			"down_2",
			"down_3",
			"down_0",
			"down_1",
		},
		6,
		true,
	); err != nil {
		return nil, err
	}

	if err := manager.AddAnimation(
		"walk_right",
		[]string{
			"right_2",
			"right_3",
			"right_1",
			"right_0",
		},
		6,
		true,
	); err != nil {
		return nil, err
	}

	animator, err := sprite.NewAnimator(manager, "walk_down")
	if err != nil {
		return nil, err
	}

	player := &Player{
		Animator:  animator,
		Direction: DirectionDown,
	}

	return player, nil
}

func (p *Player) StartWalking(direction Direction) error {
	p.Direction = direction
	p.moving = true

	return p.Animator.Play(animationForDirection(direction))
}

func (p *Player) StopWalking() {
	if !p.moving {
		return
	}

	p.moving = false
	p.Animator.Stop(p.Animator.AnimationName())
}

func animationForDirection(direction Direction) string {
	switch direction {
	case DirectionUp:
		return "walk_up"

	case DirectionRight:
		return "walk_right"

	case DirectionDown:
		return "walk_down"

	case DirectionLeft:
		return "walk_right"

	default:
		return "walk_down"
	}
}
