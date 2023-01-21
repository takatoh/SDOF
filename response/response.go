package response

type Response struct {
	Acc float64
	Vel float64
	Dis float64
}

func NewResponse(acc, vel, dis float64) *Response {
	p := new(Response)
	p.Acc = acc
	p.Vel = vel
	p.Dis = dis
	return p
}
