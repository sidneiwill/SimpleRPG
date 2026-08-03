package game

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	audioContext *audio.Context
	musicPlayer  *audio.Player

	playerX            float64
	playerY            float64
	flipX              bool
	colliders          []collisionRect
	showCollisionDebug bool

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

	g := &Game{
		player:   player,
		mapImage: mapImage,
		playerX:  300,
		playerY:  150,
		colliders: []collisionRect{
			// Map obstacles use world coordinates.
			{X: 147, Y: 82, W: 91, H: 67},  // House on the western island.
			{X: 260, Y: 0, W: 23, H: 123},  // Vertical pond.
			{X: 498, Y: 211, W: 78, H: 75}, // House on the southern island.
		},
		inventoryItems: []string{
			"Lantern",
			"Red Potion",
			"Iron Axe",
		},
	}

	if err := g.PlaySong(); err != nil {
		return nil, err
	}

	return g, nil
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

	if g.showCollisionDebug {
		g.drawCollisionDebug(screen)
	}

	if g.inventoryOpen {
		g.drawInventory(screen)
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		g.showCollisionDebug = !g.showCollisionDebug
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		g.inventoryOpen = !g.inventoryOpen
		if g.inventoryOpen {
			g.player.StopWalking()
		}

	}

	mapWidth := float64(g.mapImage.Bounds().Dx())
	mapHeight := float64(g.mapImage.Bounds().Dy())

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
