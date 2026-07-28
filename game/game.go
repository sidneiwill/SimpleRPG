package game

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	inventoryOpen   bool
	inventoryCursor int
	inventoryItems  []string
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
		inventoryItems: []string{
			"Lantern",
			"Red Potion",
			"Iron Axe",
		},
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

	if g.inventoryOpen {
		g.drawInventory(screen)
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		g.inventoryOpen = !g.inventoryOpen
		if g.inventoryOpen {
			g.player.StopWalking()
		}

	}

	mapWidth := float64(g.mapImage.Bounds().Dx())
	mapHeight := float64(g.mapImage.Bounds().Dy())

	g.playerX = mathutil.Clamp(g.playerX, 0, mapWidth-16)
	g.playerY = mathutil.Clamp(g.playerY, 0, mapHeight-32)

	// Update the player
	g.player.Animator.Update()

	// Updates the camera
	g.camera.Follow(g.playerX, g.playerY, ScreenWidth, ScreenHeight, int(mapWidth), int(mapHeight), cameraSmoothness)

	if g.inventoryOpen {
		g.updateInventory()
		return nil
	}

	if err := g.updatePlayerMovement(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Layout(
	outsideWidth int,
	outsideHeight int,
) (int, int) {
	return ScreenWidth, ScreenHeight
}
