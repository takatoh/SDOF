package directintegration

func LinearAcc(h, w, dt float64, nn int, ddy []float64) ([]float64, []float64, []float64) {
	a := 1.0 + h*w*dt + (w*w*dt*dt)/6.0
	b := w * w
	c := 2.0*h*w + w*w*dt
	d := h*w*dt + (w*w*dt*dt)/3.0

	var ddx, dx, x []float64
	ddx = append(ddx, -ddy[0])
	dx = append(dx, 0.0)
	x = append(x, 0.0)

	var ddx2, dx2, x2 float64
	for m := 1; m < nn; m++ {
		ddx2 = -(ddy[m] + b*x[m-1] + c*dx[m-1] + d*ddx[m-1]) / a
		dx2 = dx[m-1] + (ddx[m-1]+ddx2)*dt/2.0
		x2 = x[m-1] + dx[m-1]*dt + (2.0*ddx[m-1]+ddx2)*dt*dt/6.0
		ddx = append(ddx, ddx2)
		dx = append(dx, dx2)
		x = append(x, x2)
	}

	return ddx, dx, x
}
