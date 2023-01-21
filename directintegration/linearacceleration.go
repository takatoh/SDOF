package directintegration

import "github.com/takatoh/msdofs/response"

func LinearAcc(h, w, dt float64, nn int, ddy []float64) []*response.Response {
	a := 1.0 + h*w*dt + (w*w*dt*dt)/6.0
	b := w * w
	c := 2.0*h*w + w*w*dt
	d := h*w*dt + (w*w*dt*dt)/3.0

	var resp []*response.Response
	ddx := -ddy[0]
	dx := 0.0
	x := 0.0
	resp = append(resp, response.NewResponse(ddx, dx, x))

	var ddx2, dx2, x2 float64
	for m := 1; m < nn; m++ {
		ddx2 = -(ddy[m] + b*x + c*dx + d*ddx) / a
		dx2 = dx + (ddx+ddx2)*dt/2.0
		x2 = x + dx*dt + (2.0*ddx+ddx2)*dt*dt
		resp = append(resp, response.NewResponse(ddx2+ddy[m], dx2, x2))
		ddx, dx, x = ddx2, dx2, x2
	}

	return resp
}
