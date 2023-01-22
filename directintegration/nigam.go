package directintegration

import (
	"math"

	"github.com/takatoh/sdof/response"
)

func Nigam(h, w, dt float64, nn int, ddy []float64) []*response.Response {
	w2 := w * w
	hw := h * w
	wd := w * math.Sqrt(1.0-h*h)
	wdt := wd * dt
	e := math.Exp(-hw * dt)
	cwdt := math.Cos(wdt)
	swdt := math.Sin(wdt)
	a11 := e * (cwdt + hw*swdt/wd)
	a12 := e * swdt / wd
	a21 := -e * w2 * swdt / wd
	a22 := e * (cwdt - hw*swdt/wd)
	ss := -hw*swdt - wd*cwdt
	cc := -hw*cwdt + wd*swdt
	s1 := (e*ss + wd) / w2
	c1 := (e*cc + hw) / w2
	s2 := (e*dt*ss + hw*s1 + wd*c1) / w2
	c2 := (e*dt*cc + hw*c1 - wd*s1) / w2
	s3 := dt*s1 - s2
	c3 := dt*c1 - c2
	b11 := -s2 / wdt
	b12 := -s3 / wdt
	b21 := (hw*s2 - wd*c2) / wdt
	b22 := (hw*s3 - wd*c3) / wdt

	var resp []*response.Response
	ddx := 2.0 * h * w * ddy[0] * dt
	dx := -ddy[0] * dt
	x := 0.0
	resp = append(resp, response.NewResponse(ddx, dx, x))

	for m := 1; m < nn; m++ {
		dxf := dx
		xf := x
		ddym := ddy[m]
		ddyf := ddy[m-1]
		x = a12*dxf + a11*xf + b12*ddym + b11*ddyf
		dx := a22*dxf + a21*xf + b22*ddym + b21*ddyf
		ddx := -2.0*hw*dx - w2*x
		resp = append(resp, response.NewResponse(ddx, dx, x))
	}

	return resp
}
