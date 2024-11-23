package app

import (
	"log"

	api "pow/internal/api/pow"
	"pow/internal/config"
	"pow/internal/service"
	powService "pow/internal/service/pow"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	powService service.PowService

	powImpl *api.Implementation
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

func (s *serviceProvider) UserService() service.PowService {
	if s.powService == nil {
		s.powService = powService.NewService()
	}

	return s.powService
}

func (s *serviceProvider) PercImpl() *api.Implementation {
	if s.powImpl == nil {
		s.powImpl = api.NewImplementation(s.UserService())
	}

	return s.powImpl
}
