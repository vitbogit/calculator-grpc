package user

import (
	"sub/internal/service"

	desc "sub/pkg/sub_v1"
)

type Implementation struct {
	desc.UnimplementedSubServer
	subService service.SubService
}

func NewImplementation(userService service.SubService) *Implementation {
	return &Implementation{
		subService: userService,
	}
}
