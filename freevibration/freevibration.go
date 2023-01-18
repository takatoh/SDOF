package freevibration

import "math"

func naturalCircularFreqency(m, k float64) float64 {
	return math.Sqrt(k / m)
}

func naturalFreqency(m, k float64) float64 {
	return math.Sqrt(m / k)
}

func naturalPeriod(m, k float64) float64 {
	return 2.0 * math.Pi * math.Sqrt(m/k)
}

func dampedNaturalCircularFreqency(omega, h float64) float64 {
	return omega * math.Sqrt(1.0-h*h)
}

func dampedNaturalPeriod(omega, h float64) float64 {
	return 2.0 * math.Pi / (omega * math.Sqrt(1.0-h*h))
}
