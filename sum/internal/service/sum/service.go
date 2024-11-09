package sum

import (
	def "sum/internal/service"
)

var _ def.SumService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
