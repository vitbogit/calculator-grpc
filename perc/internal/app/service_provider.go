package app

import (
	"log"

	api "perc/internal/api/perc"
	"perc/internal/config"
	"perc/internal/service"
	percService "perc/internal/service/perc"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	percService service.PercService

	percImpl *api.Implementation
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

func (s *serviceProvider) UserService() service.PercService {
	if s.percService == nil {
		s.percService = percService.NewService()
	}

	return s.percService
}

func (s *serviceProvider) PercImpl() *api.Implementation {
	if s.percImpl == nil {
		s.percImpl = api.NewImplementation(s.UserService())
	}

	return s.percImpl
}
