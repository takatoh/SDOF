package directintegration

import "github.com/takatoh/msdofs/response"

const (
	Theta = 1.4
)

func WilsonTheta(h, w, dt float64, nn int, ddy []float64) []*response.Response {
	tdt := Theta * dt
	w2 := w * w
	hw := 2.0 * h * w

	cdenom := 1.0 + tdt*hw/2.0 + w2*tdt*tdt/6.0
	cnume2 := hw * w2 * tdt
	cnume3 := tdt*hw*tdt/2.0 + w2*tdt*tdt/3.0

	var resp []*response.Response
	ddx := 0.0
	dx := 0.0
	x := 0.0
	resp = append(resp, response.NewResponse(ddx, dx, x))

	for m := 1; m < nn; m++ {
		f := (Theta-1.0)*ddy[m] - Theta*ddy[m+1]
		ath := (f - w2*x - cnume2*dx - cnume3*ddx) / cdenom
		ddx2 := ((Theta-1.0)*ddx + ath) / Theta
		dx2 := dx + (ddx+ddx2)*dt/2.0
		x2 := x + dx*dt + (2.0*ddx+ddx2)*dt*dt/6.0
		resp = append(resp, response.NewResponse(ddx2+ddy[m], dx2, x2))
		ddx, dx, x = ddx2, dx2, x2
	}

	return resp
}
