package sum

import (
	def "root/internal/service"
)

var _ def.SumService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
