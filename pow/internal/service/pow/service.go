package pow

import (
	def "pow/internal/service"
)

var _ def.PowService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
