package user

import (
	"pow/internal/service"
	desc "pow/pkg/pow_v1"
)

type Implementation struct {
	desc.UnimplementedPowServer
	powService service.PowService
}

func NewImplementation(userService service.PowService) *Implementation {
	return &Implementation{
		powService: userService,
	}
}
