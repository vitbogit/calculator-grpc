package app

import (
	"log"

	api "calc/internal/api/calc"
	"calc/internal/config"
	"calc/internal/service"
	calcService "calc/internal/service/calc"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	calcService service.CalcService

	calcImpl *api.Implementation
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

func (s *serviceProvider) CalcService() service.CalcService {
	if s.calcService == nil {
		s.calcService = calcService.NewService()
	}

	return s.calcService
}

func (s *serviceProvider) CalcImpl() *api.Implementation {
	if s.calcImpl == nil {
		s.calcImpl = api.NewImplementation(s.CalcService())
	}

	return s.calcImpl
}
