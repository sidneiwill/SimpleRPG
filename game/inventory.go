package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	panelX = 32
	panelY = 24
	panelW = 256
	panelH = 184
)

func (g *Game) updateInventory() {
	if len(g.inventoryItems) == 0 {
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.inventoryCursor--
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.inventoryCursor++
	}

	if g.inventoryCursor < 0 {
		g.inventoryCursor = len(g.inventoryItems) - 1
	}

	if g.inventoryCursor >= len(g.inventoryItems) {
		g.inventoryCursor = 0
	}
}

func (g *Game) drawInventory(screen *ebiten.Image) {
	vector.FillRect(screen, float32(panelX), float32(panelY), float32(panelW), float32(panelH), color.RGBA{220, 220, 255, 255}, true)
	vector.FillRect(screen, float32(panelX), float32(panelY), float32(panelW), float32(panelH), color.RGBA{16, 24, 40,
		235}, true)
	vector.FillRect(screen, float32(panelX), float32(panelY), float32(panelW), 2, color.RGBA{220, 220, 255, 255}, true)
	vector.FillRect(screen, float32(panelX), float32(panelY+panelH-2), float32(panelW), 2, color.RGBA{220, 220, 255,
		255}, true)
	vector.FillRect(screen, float32(panelX), float32(panelY), 2, float32(panelH), color.RGBA{220, 220, 255, 255}, true)
	vector.FillRect(screen, float32(panelX+panelW-2), float32(panelY), 2, float32(panelH), color.RGBA{220, 220, 255,
		255}, true)
	ebitenutil.DebugPrintAt(screen, "Backpack", panelX+8, panelY+8)
	vector.FillRect(screen, float32(panelX), float32(panelY)+24, float32(panelW), 1, color.RGBA{220, 220, 255, 255}, true)

	for i, item := range g.inventoryItems {
		y := panelY + 32 + i*32
		prefix := "  "
		if i == g.inventoryCursor {
			prefix = ">>  "
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s%s", prefix, item), 40, y)
	}
}
