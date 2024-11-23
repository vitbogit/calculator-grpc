package sub

import (
	def "sub/internal/service"
)

var _ def.SubService = (*service)(nil)

type service struct {
}

func NewService() *service {
	return &service{}
}
