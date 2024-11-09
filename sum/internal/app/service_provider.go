package app

import (
	"log"

	api "sum/internal/api/sum"
	"sum/internal/config"
	"sum/internal/service"
	sumService "sum/internal/service/sum"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	sumService service.SumService

	sumImpl *api.Implementation
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

func (s *serviceProvider) SumImpl() *api.Implementation {
	if s.sumImpl == nil {
		s.sumImpl = api.NewImplementation(s.UserService())
	}

	return s.sumImpl
}
