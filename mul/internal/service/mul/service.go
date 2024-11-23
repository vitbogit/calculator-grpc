package mul

import (
	def "mul/internal/service"
)

var _ def.MulService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
