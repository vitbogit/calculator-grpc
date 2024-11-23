package perc

import (
	def "perc/internal/service"
)

var _ def.PercService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
