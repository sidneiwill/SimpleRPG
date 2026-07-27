package main

import (
	"bytes"
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"sidneiwill.dev/BaseRPGMovement/player"
)

//go:embed assets/map.png
var mapImageBytes []byte

//go:embed assets/character.png
var playerSpriteSheet []byte

const (
	screenWidth  = 320
	screenHeight = 240
	windowWidth  = 640
	windowHeight = 480
)

type Game struct {
	player   *player.Player
	mapImage *ebiten.Image

	playerX float64
	playerY float64
	flipX   bool

	cameraX float64
	cameraY float64
}

func clamp(value, min, max float64) float64 {
	if max < min {
		return min
	}

	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func NewGame() (*Game, error) {
	mapImage, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(mapImageBytes))
	if err != nil {
		return nil, err
	}

	player, err := player.NewAnimator(playerSpriteSheet)
	if err != nil {
		return nil, err
	}
	player.Animator.Stop(player.Animator.AnimationName())

	return &Game{
		player:   player,
		mapImage: mapImage,
		playerX:  0,
		playerY:  0,
	}, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mapOptions := &ebiten.DrawImageOptions{}
	mapOptions.GeoM.Translate(-g.cameraX, -g.cameraY)
	screen.DrawImage(g.mapImage, mapOptions)

	drawX := g.playerX - float64(g.player.Animator.CurrentFrame().Bounds().Dx()/2)
	drawY := g.playerY - float64(g.player.Animator.CurrentFrame().Bounds().Dy()/2)

	g.player.Animator.Draw(
		screen,
		drawX-g.cameraX,
		drawY-g.cameraY,
		1,       // Scale
		g.flipX, // Horizontal direction
	)
}

func (g *Game) Update() error {
	// Movement Controls
	const movementSpeed = 1.0

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

	mapWidth := float64(g.mapImage.Bounds().Dx())
	mapHeight := float64(g.mapImage.Bounds().Dy())

	g.playerX = clamp(g.playerX, 0, mapWidth-16)
	g.playerY = clamp(g.playerY, 0, mapHeight-32)

	// Update the player
	g.player.Animator.Update()

	// Updates the camera
	g.cameraX = g.playerX - screenWidth/2
	g.cameraY = g.playerY - screenHeight/2

	g.cameraX = clamp(g.cameraX, 0, mapWidth-screenWidth)
	g.cameraY = clamp(g.cameraY, 0, mapHeight-screenHeight)

	return nil
}

func (g *Game) Layout(
	outsideWidth int,
	outsideHeight int,
) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Pokemon Basic Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
