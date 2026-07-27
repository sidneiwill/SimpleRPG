package camera

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
	targetCameraX := targetX - float64((screenWidth))/2
	targetCameraY := targetY - float64((screenHeight))/2

	targetCameraX = clamp(targetCameraX, 0, float64(mapWidth-screenWidth))
	targetCameraY = clamp(targetCameraY, 0, float64(mapHeight-screenHeight))

	c.X = lerp(c.X, targetCameraX, smoothness)
	c.Y = lerp(c.Y, targetCameraY, smoothness)
}

// Take a value and clamp it to a given limit
func clamp(value, min, max float64) float64 {
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
