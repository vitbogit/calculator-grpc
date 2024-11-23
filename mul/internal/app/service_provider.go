package app

import (
	"log"

	api "mul/internal/api/mul"
	"mul/internal/config"
	"mul/internal/service"
	mulService "mul/internal/service/mul"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	mulService service.MulService

	mulImpl *api.Implementation
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

func (s *serviceProvider) UserService() service.MulService {
	if s.mulService == nil {
		s.mulService = mulService.NewService()
	}

	return s.mulService
}

func (s *serviceProvider) MulImpl() *api.Implementation {
	if s.mulImpl == nil {
		s.mulImpl = api.NewImplementation(s.UserService())
	}

	return s.mulImpl
}
