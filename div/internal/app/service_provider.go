package app

import (
	api "div/internal/api/div"
	"div/internal/config"
	"div/internal/service"
	divService "div/internal/service/div"
	"log"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	divService service.DivService

	divImpl *api.Implementation
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

func (s *serviceProvider) UserService() service.DivService {
	if s.divService == nil {
		s.divService = divService.NewService()
	}

	return s.divService
}

func (s *serviceProvider) DivImpl() *api.Implementation {
	if s.divImpl == nil {
		s.divImpl = api.NewImplementation(s.UserService())
	}

	return s.divImpl
}
