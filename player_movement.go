package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"sidneiwill.dev/BaseRPGMovement/player"
)

func (g *Game) updatePlayerMovement() error {
	// Movement Controls
	const movementSpeed = 1.0

	moving := false

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= movementSpeed
		if err := g.player.StartWalking(player.DirectionUp); err != nil {
			return err
		}

		moving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += movementSpeed
		if err := g.player.StartWalking(
			player.DirectionDown,
		); err != nil {
			return err
		}
		moving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= movementSpeed
		g.flipX = true
		if err := g.player.StartWalking(
			player.DirectionLeft,
		); err != nil {
			return err
		}
		moving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += movementSpeed
		g.flipX = false
		if err := g.player.StartWalking(
			player.DirectionRight,
		); err != nil {
			return err
		}
		moving = true
	}

	if !moving {
		g.player.StopWalking()
	}

	return nil
}
