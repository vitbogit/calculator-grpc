package app

import (
	"log"

	api "root/internal/api/root"
	"root/internal/config"
	"root/internal/service"
	sumService "root/internal/service/root"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	sumService service.SumService

	userImpl *api.Implementation
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

func (s *serviceProvider) UserService() service.SumService {
	if s.sumService == nil {
		s.sumService = sumService.NewService()
	}

	return s.sumService
}

func (s *serviceProvider) UserImpl() *api.Implementation {
	if s.userImpl == nil {
		s.userImpl = api.NewImplementation(s.UserService())
	}

	return s.userImpl
}
