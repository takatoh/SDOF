package directintegration

const (
	Theta = 1.4
)

func WilsonTheta(h, w, dt float64, nn int, ddy []float64) ([]float64, []float64, []float64) {
	tdt := Theta * dt
	w2 := w * w
	hw := 2.0 * h * w

	cdenom := 1.0 + hw*tdt/2.0 + w2*tdt*tdt/6.0
	cnume2 := hw + w2*tdt
	cnume3 := hw*tdt/2.0 + w2*tdt*tdt/3.0

	ddx := make([]float64, 0)
	dx := make([]float64, 0)
	x := make([]float64, 0)
	ddx = append(ddx, 0.0)
	dx = append(dx, 0.0)
	x = append(x, 0.0)

	a0 := ddx[0]
	for m := 1; m < nn-1; m++ {
		f := (Theta-1.0)*ddy[m] - Theta*ddy[m+1]
		ath := (f - w2*x[m-1] - cnume2*dx[m-1] - cnume3*a0) / cdenom
		a1 := ((Theta-1.0)*a0 + ath) / Theta
		x = append(x, x[m-1]+dt*dx[m-1]+(2.0*a0+a1)*dt*dt/6.0)
		dx = append(dx, dx[m-1]+(a0+a1)*dt/2.0)
		ddx = append(ddx, a1+ddy[m])
		a0 = a1
	}

	return ddx, dx, x
}
