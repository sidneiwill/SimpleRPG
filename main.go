package main

import (
	"bytes"
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	camera "sidneiwill.dev/BaseRPGMovement/game"
	"sidneiwill.dev/BaseRPGMovement/player"
)

//go:embed assets/map.png
var mapImageBytes []byte

//go:embed assets/character.png
var playerSpriteSheet []byte

const (
	screenWidth      = 320
	screenHeight     = 240
	windowWidth      = 640
	windowHeight     = 480
	cameraSmoothness = 0.08
)

type Game struct {
	player   *player.Player
	mapImage *ebiten.Image

	playerX float64
	playerY float64
	flipX   bool

	camera camera.Camera
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

func lerp(current, target, amount float64) float64 {
	return current + (target-current)*amount
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
		playerX:  screenWidth / 2,
		playerY:  screenHeight / 2,
	}, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mapOptions := &ebiten.DrawImageOptions{}
	mapOptions.GeoM.Translate(-g.camera.X, -g.camera.Y)
	screen.DrawImage(g.mapImage, mapOptions)

	// Draw the sprite in the center of the "Object"
	drawX := g.playerX - float64(g.player.Animator.CurrentFrame().Bounds().Dx()/2)
	drawY := g.playerY - float64(g.player.Animator.CurrentFrame().Bounds().Dy()/2)

	g.player.Animator.Draw(
		screen,
		drawX-g.camera.X,
		drawY-g.camera.Y,
		1,       // Scale
		g.flipX, // Horizontal direction
	)
}

func (g *Game) Update() error {
	if err := g.updatePlayerMovement(); err != nil {
		return err
	}

	mapWidth := float64(g.mapImage.Bounds().Dx())
	mapHeight := float64(g.mapImage.Bounds().Dy())

	g.playerX = clamp(g.playerX, 0, mapWidth-16)
	g.playerY = clamp(g.playerY, 0, mapHeight-32)

	// Update the player
	g.player.Animator.Update()

	// Updates the camera
	g.camera.Follow(g.playerX, g.playerY, screenWidth, screenHeight, int(mapWidth), int(mapHeight), cameraSmoothness)

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
