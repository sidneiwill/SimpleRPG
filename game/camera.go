package game

import "sidneiwill.dev/BaseRPGMovement/internal/mathutil"

type Camera struct {
	X float64
	Y float64
}

func (c *Camera) Follow(
	targetX float64,
	targetY float64,
	screenWidth int,
	screenHeight int,
	mapWidth int,
	mapHeight int,
	smoothness float64,
) {
	// Center the camera on the target
	targetCameraX := targetX - float64(screenWidth)/2
	targetCameraY := targetY - float64(screenHeight)/2

	// Keep the camera inside the map bounds
	targetCameraX = mathutil.Clamp(targetCameraX, 0, float64(mapWidth-screenWidth))
	targetCameraY = mathutil.Clamp(targetCameraY, 0, float64(mapHeight-screenHeight))

	// Smoothly move the camera toward the target position
	c.X = mathutil.Lerp(c.X, targetCameraX, smoothness)
	c.Y = mathutil.Lerp(c.Y, targetCameraY, smoothness)
}
