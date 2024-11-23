package user

import (
	"calc/internal/service"

	desc "calc/pkg/calc_v1"
)

type Implementation struct {
	desc.UnimplementedCalcServer
	sumService service.CalcService
}

func NewImplementation(userService service.CalcService) *Implementation {
	return &Implementation{
		sumService: userService,
	}
}
