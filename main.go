package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"sidneiwill.dev/BaseRPGMovement/game"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(game.WindowWidth, game.WindowHeight)
	ebiten.SetWindowTitle("Pokemon Basic Game")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
