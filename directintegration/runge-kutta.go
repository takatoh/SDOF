package directintegration

func RK4(h, w, dt float64, nn int, ddy []float64) ([]float64, []float64, []float64) {
	am := 1.0              // Mass
	ak := w * w * am       // Stiffness
	ac := 2.0 * h * w * am // Damping constant

	ddx := make([]float64, nn)
	dx := make([]float64, nn)
	x := make([]float64, nn)

	var f1, f2, fm float64
	var x1, x2, dx1, dx2 float64
	var racc, rvel, rdis float64
	var kx1, kx2, kx3, kx4, kdx1, kdx2, kdx3, kdx4 float64
	racc = 0.0
	rvel = 0.0
	rdis = 0.0
	for m := 0; m < nn; m++ {
		if m >= 1 {
			f1 = -am * ddy[m-1]
		} else {
			f1 = 0.0
		}
		f2 = -am * ddy[m]
		fm = 0.5 * (f1 + f2)
		x1 = rdis
		dx1 = rvel

		kx1, kdx1 = func4rk(am, ac, ak, f1, x1, dx1)
		kx2, kdx2 = func4rk(am, ac, ak, fm, x1+dt*kx1/2.0, dx1+dt*kdx1/2.0)
		kx3, kdx3 = func4rk(am, ac, ak, fm, x1+dt*kx2/2.0, dx1+dt*kdx2/2.0)
		kx4, kdx4 = func4rk(am, ac, ak, f2, x1+dt*kx3, dx1+dt*kdx3)
		x2 = x1 + dt*(kx1+2.0*kx2+2.0*kx3+kx4)/6.0
		dx2 = dx1 + dt*(kdx1+2.0*kdx2+2.0*kdx3+kdx4)/6.0

		racc = f2/am - ac/am*dx2 - ak/am*x2
		rvel = dx2
		rdis = x2

		ddx[m] = racc + ddy[m]
		dx[m] = rvel
		x[m] = rdis
	}

	return ddx, dx, x
}

// Function for RK4.
func func4rk(am, ac, ak, ff, x, y float64) (float64, float64) {
	dxdt := y                         // Velocity
	dydt := ff/am - ac/am*y - ak/am*x // Acceleration
	return dxdt, dydt
}
