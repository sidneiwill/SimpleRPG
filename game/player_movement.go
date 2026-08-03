package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"sidneiwill.dev/BaseRPGMovement/player"
)

func movementDelta(up, down, left, right bool, speed float64) (float64, float64) {
	var dx, dy float64

	if up {
		dy--
	}
	if down {
		dy++
	}
	if left {
		dx--
	}
	if right {
		dx++
	}

	length := math.Hypot(dx, dy)
	if length == 0 {
		return 0, 0
	}
	return dx / length * speed, dy / length * speed
}

func (g *Game) updatePlayerMovement() error {
	// Movement Controls
	const movementSpeed = 1.0

	up := ebiten.IsKeyPressed(ebiten.KeyW)
	down := ebiten.IsKeyPressed(ebiten.KeyS)
	left := ebiten.IsKeyPressed(ebiten.KeyA)
	right := ebiten.IsKeyPressed(ebiten.KeyD)

	dx, dy := movementDelta(up, down, left, right, movementSpeed)

	actualDX, actualDY := g.movePlayer(dx, dy)

	if actualDX == 0 && actualDY == 0 {
		g.player.StopWalking()
		return nil
	}

	var direction player.Direction

	// This deliberately gives horizontal animation priority on diagonals.
	switch {
	case actualDX < 0:
		direction = player.DirectionLeft
		g.flipX = true
	case actualDX > 0:
		direction = player.DirectionRight
		g.flipX = false
	case actualDY < 0:
		direction = player.DirectionUp
	case actualDY > 0:
		direction = player.DirectionDown
	}

	return g.player.StartWalking(direction)
}
