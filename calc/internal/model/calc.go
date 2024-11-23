package model

import "math/big"

type CalcNFRequest struct {
	A        *big.Float
	B        *big.Float
	Rounding uint32
}

type CalcFRequest struct {
	A1       *big.Float
	A2       *big.Float
	B1       *big.Float
	B2       *big.Float
	Rounding uint32
}

type CalcRequest struct {
	CalcNFRequest *CalcNFRequest
	CalcFRequest  *CalcFRequest
}

type CalcNFResponse struct {
	C *big.Float
}

type CalcFResponse struct {
	C1 *big.Float
	C2 *big.Float
}

type CalcResponse struct {
	CalcNFResponse *CalcNFResponse
	CalcFResponse  *CalcFResponse
	Precise        uint32
}
