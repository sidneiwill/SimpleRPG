package game

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"sidneiwill.dev/BaseRPGMovement/internal/mathutil"
	"sidneiwill.dev/BaseRPGMovement/player"
)

const (
	ScreenWidth      = 320
	ScreenHeight     = 240
	WindowWidth      = 640
	WindowHeight     = 480
	cameraSmoothness = 0.08
)

type Game struct {
	player   *player.Player
	mapImage *ebiten.Image

	playerX float64
	playerY float64
	flipX   bool

	camera Camera
}

func NewGame() (*Game, error) {
	mapImage, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(MapImageBytes))
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
		playerX:  ScreenWidth / 2,
		playerY:  ScreenHeight / 2,
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

	g.playerX = mathutil.Clamp(g.playerX, 0, mapWidth-16)
	g.playerY = mathutil.Clamp(g.playerY, 0, mapHeight-32)

	// Update the player
	g.player.Animator.Update()

	// Updates the camera
	g.camera.Follow(g.playerX, g.playerY, ScreenWidth, ScreenHeight, int(mapWidth), int(mapHeight), cameraSmoothness)

	return nil
}

func (g *Game) Layout(
	outsideWidth int,
	outsideHeight int,
) (int, int) {
	return ScreenWidth, ScreenHeight
}
