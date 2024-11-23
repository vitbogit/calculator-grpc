package user

import (
	"div/internal/service"

	desc "div/pkg/div_v1"
)

type Implementation struct {
	desc.UnimplementedDivServer
	divService service.DivService
}

func NewImplementation(userService service.DivService) *Implementation {
	return &Implementation{
		divService: userService,
	}
}
