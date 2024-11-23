package converter

import (
	"calc/internal/model"
	desc "calc/pkg/calc_v1"
)

func ToCallRequestFromDesc(info *desc.CallRequest) *model.CallRequest {
	return &model.CallRequest{
		Services:     ToServicesFromDesc(info),
		CalcRequests: ToCalcRequestsFromDesc(info),
	}
}
