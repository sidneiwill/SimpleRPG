package mathutil

func Clamp(value, min, max float64) float64 {
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

func Lerp(current, target, amount float64) float64 {
	return current + (target-current)*amount
}
