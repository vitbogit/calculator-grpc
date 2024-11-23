package model

type CallRequest struct {
	Services     *Services
	CalcRequests []*CalcRequest
}

type CallResponse struct {
	Services     *Services
	CalcRequests []*CalcRequest
	Precise      uint32
}
