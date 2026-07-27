package game

import (
	_ "embed"
)

//go:embed assets/map.png
var MapImageBytes []byte

//go:embed assets/character.png
var playerSpriteSheet []byte
