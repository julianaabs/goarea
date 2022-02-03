package goarea

import "math"

const Pi = 3.1416

func Circ(radio float64) float64 {
	return math.Pow(radio, 2) * Pi
}

func Rect(base, height float64) float64 {
	return base * height
}

func _TriangleEq(base, height float64) float64 {
	return (base * height) / 2
}
