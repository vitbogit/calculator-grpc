package converter

import (
	"calc/internal/model"
	desc "calc/pkg/calc_v1"
)

func ToCalcRequestsFromDesc(info *desc.CallRequest) []*model.CalcRequest {
	calcRequests := info.GetCalcRequests()

	modelCalcRequests := make([]*model.CalcRequest, 0, len(calcRequests))

	for _, calcRequest := range calcRequests {
		modelCalcRequests = append(modelCalcRequests, ToCalcRequestFromDesc(calcRequest))
	}

	return modelCalcRequests
}

func ToCalcRequestFromDesc(info *desc.CalcRequest) *model.CalcRequest {
	return &model.CalcRequest{
		CalcNFRequest: ToCalcNFRequestFromDesc(info.GetCalcNFRequest()),
		CalcFRequest:  ToCalcFRequestFromDesc(info.GetCalcFRequest()),
	}
}

func ToCalcNFRequestFromDesc(info *desc.CalcNFRequest) *model.CalcNFRequest {
	if len(info.GetA()) == 0 && len(info.GetB()) == 0 {
		return nil
	}
	return &model.CalcNFRequest{
		A:        ToBigFloatFromString(info.GetA()),
		B:        ToBigFloatFromString(info.GetB()),
		Rounding: info.GetRounding(),
	}
}

func ToCalcFRequestFromDesc(info *desc.CalcFRequest) *model.CalcFRequest {
	if len(info.GetA1()) == 0 && len(info.GetA2()) == 0 && len(info.GetB1()) == 0 && len(info.GetB2()) == 0 {
		return nil
	}
	return &model.CalcFRequest{
		A1:       ToBigFloatFromString(info.GetA1()),
		A2:       ToBigFloatFromString(info.GetA2()),
		B1:       ToBigFloatFromString(info.GetB1()),
		B2:       ToBigFloatFromString(info.GetB2()),
		Rounding: info.GetRounding(),
	}
}
