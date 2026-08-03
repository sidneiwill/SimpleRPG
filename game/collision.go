package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	obstacleDebugColor = color.RGBA{R: 255, A: 255}
	hitboxDebugColor   = color.RGBA{R: 255, G: 255, A: 255}
)

type collisionRect struct {
	X float64
	Y float64
	W float64
	H float64
}

func rectanglesOverlap(a, b collisionRect) bool {
	return a.X < b.X+b.W &&
		a.X+a.W > b.X &&
		a.Y < b.Y+b.H &&
		a.Y+a.H > b.Y
}

func playerHitbox(x, y float64) collisionRect {
	return collisionRect{
		X: x - 5,
		Y: y + 4,
		W: 10,
		H: 8,
	}
}

func (g *Game) positionBlocked(x, y float64) bool {
	hitbox := playerHitbox(x, y)
	mapWidth := float64(g.mapImage.Bounds().Dx())
	mapHeight := float64(g.mapImage.Bounds().Dy())

	if hitbox.X < 0 ||
		hitbox.Y < 0 ||
		hitbox.X+hitbox.W > mapWidth ||
		hitbox.Y+hitbox.H > mapHeight {
		return true
	}

	for _, obstacle := range g.colliders {
		if rectanglesOverlap(hitbox, obstacle) {
			return true
		}
	}

	return false
}

func (g *Game) movePlayer(dx, dy float64) (float64, float64) {
	var actualX, actualY float64

	nextX := g.playerX + dx
	if !g.positionBlocked(nextX, g.playerY) {
		g.playerX = nextX
		actualX = dx
	}

	nextY := g.playerY + dy
	if !g.positionBlocked(g.playerX, nextY) {
		g.playerY = nextY
		actualY = dy
	}

	return actualX, actualY
}

func (g *Game) drawCollisionDebug(screen *ebiten.Image) {
	for _, obstacle := range g.colliders {
		drawCollisionRect(screen, obstacle, g.camera, obstacleDebugColor)
	}

	drawCollisionRect(
		screen,
		playerHitbox(g.playerX, g.playerY),
		g.camera,
		hitboxDebugColor,
	)
}

func drawCollisionRect(
	screen *ebiten.Image,
	rect collisionRect,
	camera Camera,
	outline color.Color,
) {
	vector.StrokeRect(
		screen,
		float32(rect.X-camera.X),
		float32(rect.Y-camera.Y),
		float32(rect.W),
		float32(rect.H),
		1,
		outline,
		false,
	)
}
