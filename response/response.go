package response

type Response struct {
	Period float64
	Sa     float64
	Sv     float64
	Sd     float64
}

func NewResponse(period, sa, sv, sd float64) *Response {
	p := new(Response)
	p.Period = period
	p.Sa = sa
	p.Sv = sv
	p.Sd = sd
	return p
}
