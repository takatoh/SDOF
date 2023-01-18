package response

type Response struct {
	Sa float64
	Sv float64
	Sd float64
}

func NewResponse(sa, sv, sd float64) *Response {
	p := new(Response)
	p.Sa = sa
	p.Sv = sv
	p.Sd = sd
	return p
}
