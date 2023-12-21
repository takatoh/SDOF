package freevibration

import "math"

func NaturalCircularFrequency(m, k float64) float64 {
	return math.Sqrt(k / m)
}

func NaturalFrequency(m, k float64) float64 {
	return math.Sqrt(k/m) / (2.0 * math.Pi)
}

func NaturalPeriod(m, k float64) float64 {
	return 2.0 * math.Pi * math.Sqrt(m/k)
}

func DampedNaturalCircularFrequency(omega, h float64) float64 {
	return omega * math.Sqrt(1.0-h*h)
}

func DampedNaturalPeriod(omega, h float64) float64 {
	return 2.0 * math.Pi / (omega * math.Sqrt(1.0-h*h))
}
