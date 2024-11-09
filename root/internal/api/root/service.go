package user

import (
	"root/internal/service"

	desc "root/pkg/root_v1"
)

type Implementation struct {
	desc.UnimplementedRootServer
	sumService service.SumService
}

func NewImplementation(userService service.SumService) *Implementation {
	return &Implementation{
		sumService: userService,
	}
}
