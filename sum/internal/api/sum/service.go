package user

import (
	"sum/internal/service"

	desc "sum/pkg/sum_v1"
)

type Implementation struct {
	desc.UnimplementedSumServer
	sumService service.SumService
}

func NewImplementation(userService service.SumService) *Implementation {
	return &Implementation{
		sumService: userService,
	}
}
