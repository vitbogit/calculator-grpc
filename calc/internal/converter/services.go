package converter

import (
	"calc/internal/model"
	desc "calc/pkg/calc_v1"
)

func ToServicesFromDesc(info *desc.CallRequest) *model.Services {
	services := info.GetServices()

	if services == nil || services.Services == nil {
		return &model.Services{
			Services: make([]string, 0),
		}
	}

	return &model.Services{
		Services: services.Services,
	}
}
