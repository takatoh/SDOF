package directintegration

func LinearAcc(h, w, dt float64, nn int, ddy []float64) ([]float64, []float64, []float64) {
	a := 1.0 + h*w*dt + w*w*dt*dt/6.0
	b := w * w
	c := 2.0*h*w + w*w*dt
	d := h*w*dt + w*w*dt*dt/3.0

	ddx := make([]float64, 0)
	dx := make([]float64, 0)
	x := make([]float64, 0)
	ddx = append(ddx, -ddy[0])
	dx = append(dx, 0.0)
	x = append(x, 0.0)

	var a1, v1, d1 float64
	a0 := ddx[0]
	for m := 1; m < nn; m++ {
		v0 := dx[m-1]
		d0 := x[m-1]
		a1 = -(ddy[m] + b*d0 + c*v0 + d*a0) / a
		v1 = v0 + (a0+a1)*dt/2.0
		d1 = d0 + v0*dt + (2.0*a0+a1)*dt*dt/6.0
		ddx = append(ddx, ddy[m]+a1)
		dx = append(dx, v1)
		x = append(x, d1)
		a0 = a1
	}

	return ddx, dx, x
}
