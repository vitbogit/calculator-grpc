package app

import (
	"log"

	api "sub/internal/api/sub"
	"sub/internal/config"
	"sub/internal/service"
	subService "sub/internal/service/sub"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	subService service.SubService

	subImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) UserService() service.SubService {
	if s.subService == nil {
		s.subService = subService.NewService()
	}

	return s.subService
}

func (s *serviceProvider) SubImpl() *api.Implementation {
	if s.subImpl == nil {
		s.subImpl = api.NewImplementation(s.UserService())
	}

	return s.subImpl
}
