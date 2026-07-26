package main

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"sidneiwill.dev/BaseRPGMovement/player"
)

//go:embed assets/character.png
var playerSpriteSheet []byte

type Game struct {
	player *player.Player

	playerX float64
	playerY float64
	flipX   bool
}

func NewGame() (*Game, error) {
	player, err := player.NewAnimator(playerSpriteSheet)
	if err != nil {
		return nil, err
	}

	return &Game{
		player:  player,
		playerX: 100,
		playerY: 80,
	}, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Animator.Draw(
		screen,
		g.playerX,
		g.playerY,
		2,       // Scale
		g.flipX, // Horizontal direction
	)
}

func (g *Game) Update() error {
	const movementSpeed = 2.0

	moving := false

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= movementSpeed
		if err := g.player.StartWalking(
			player.DirectionUp,
		); err != nil {
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
		g.flipX = false
		if err := g.player.StartWalking(
			player.DirectionLeft,
		); err != nil {
			return err
		}
		moving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += movementSpeed
		g.flipX = true
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

	g.player.Animator.Update()
	return nil
}

func (g *Game) Layout(
	outsideWidth int,
	outsideHeight int,
) (int, int) {
	return 320, 240
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pokemon Basic Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
