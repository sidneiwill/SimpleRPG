package main

import (
	"bytes"
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"sidneiwill.dev/BaseRPGMovement/sprite"
)

//go:embed assets/character.png
var playerSpriteSheet []byte

type Game struct {
	player *sprite.Animator

	playerX float64
	playerY float64
	flipX   bool
}

func NewGame() (*Game, error) {
	sheet, _, err := ebitenutil.NewImageFromReader(
		bytes.NewReader(playerSpriteSheet),
	)
	if err != nil {
		return nil, err
	}

	manager := sprite.NewManager(sheet)

	// Row 1: walk_0 through walk_5.
	if err := manager.AddGridRow(
		"walk",
		1,
		0,
		3,
		16,
		32,
	); err != nil {
		return nil, err
	}

	if err := manager.AddAnimation(
		"walk",
		[]string{
			"walk_0",
			"walk_1",
			"walk_2",
		},
		6,
		true,
	); err != nil {
		return nil, err
	}

	player, err := sprite.NewAnimator(manager, "walk")
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
	g.player.Draw(
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

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.Play("walk")
		g.playerX -= movementSpeed
		g.flipX = false
		moving = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.Play("walk")
		g.playerX += movementSpeed
		g.flipX = true
		moving = true
	}

	if !ebiten.IsKeyPressed(ebiten.KeyA) || !ebiten.IsKeyPressed(ebiten.KeyD) {
		moving = false
	}

	if moving {
		if err := g.player.Play("walk"); err != nil {
			return err
		}
	} else {
		g.player.Stop("walk")
	}

	g.player.Update()
	return nil
}

func (g *Game) Layout(
	outsideWidth int,
	outsideHeight int,
) (int, int) {
	return 320, 180
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 360)
	ebiten.SetWindowTitle("Pokemon Basic Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
