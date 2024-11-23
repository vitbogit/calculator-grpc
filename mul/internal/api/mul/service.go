package user

import (
	"mul/internal/service"

	desc "mul/pkg/mul_v1"
)

type Implementation struct {
	desc.UnimplementedMulServer
	mulService service.MulService
}

func NewImplementation(userService service.MulService) *Implementation {
	return &Implementation{
		mulService: userService,
	}
}
