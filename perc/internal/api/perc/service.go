package user

import (
	"perc/internal/service"

	desc "perc/pkg/perc_v1"
)

type Implementation struct {
	desc.UnimplementedPercServer
	percService service.PercService
}

func NewImplementation(userService service.PercService) *Implementation {
	return &Implementation{
		percService: userService,
	}
}
