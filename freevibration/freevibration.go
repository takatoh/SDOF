package freevibration

import "math"

func NaturalCircularFreqency(m, k float64) float64 {
	return math.Sqrt(k / m)
}

func NaturalFreqency(m, k float64) float64 {
	return math.Sqrt(m/k) / (2.0 * math.Pi)
}

func NaturalPeriod(m, k float64) float64 {
	return 2.0 * math.Pi * math.Sqrt(m/k)
}

func DampedNaturalCircularFreqency(omega, h float64) float64 {
	return omega * math.Sqrt(1.0-h*h)
}

func DampedNaturalPeriod(omega, h float64) float64 {
	return 2.0 * math.Pi / (omega * math.Sqrt(1.0-h*h))
}
