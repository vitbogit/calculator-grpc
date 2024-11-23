package div

import (
	def "div/internal/service"
)

var _ def.DivService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
