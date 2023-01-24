package directintegration

const (
	Theta = 1.4
)

func WilsonTheta(h, w, dt float64, nn int, ddy []float64) ([]float64, []float64, []float64) {
	tdt := Theta * dt
	w2 := w * w
	hw := 2.0 * h * w

	cdenom := 1.0 + tdt*hw/2.0 + w2*tdt*tdt/6.0
	cnume2 := hw * w2 * tdt
	cnume3 := tdt*hw*tdt/2.0 + w2*tdt*tdt/3.0

	ddx := make([]float64, 0)
	dx := make([]float64, 0)
	x := make([]float64, 0)
	ddx = append(ddx, 0.0)
	dx = append(dx, 0.0)
	x = append(x, 0.0)

	for m := 1; m < nn; m++ {
		f := (Theta-1.0)*ddy[m] - Theta*ddy[m+1]
		ath := (f - w2*x[m-1] - cnume2*dx[m-1] - cnume3*ddx[m-1]) / cdenom
		ddx2 := ((Theta-1.0)*ddx[m-1] + ath) / Theta
		dx2 := dx[m-1] + (ddx[m-1]+ddx2)*dt/2.0
		x2 := x[m-1] + dx[m-1]*dt + (2.0*ddx[m-1]+ddx2)*dt*dt/6.0
		ddx = append(ddx, ddx2)
		dx = append(dx, dx2)
		x = append(x, x2)
	}

	return ddx, dx, x
}
