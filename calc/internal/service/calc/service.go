package sum

import (
	def "calc/internal/service"

	"google.golang.org/grpc"
)

const (
	serviceSumName           = "sum"
	serviceSumAdressEnvName  = "SERVICE_SUM_ADDRESS"
	serviceRootName          = "root"
	serviceRootAdressEnvName = "SERVICE_ROOT_ADDRESS"
	serviceSubName           = "sub"
	serviceSubAdressEnvName  = "SERVICE_SUB_ADDRESS"
	serviceMulName           = "mul"
	serviceMulAdressEnvName  = "SERVICE_MUL_ADDRESS"
	serviceDivName           = "div"
	serviceDivAdressEnvName  = "SERVICE_DIV_ADDRESS"
	servicePercName          = "perc"
	servicePercAdressEnvName = "SERVICE_PERC_ADDRESS"
	servicePowName           = "pow"
	servicePowAdressEnvName  = "SERVICE_POW_ADDRESS"
)

const (
	serviceResultTypeNonFract  = "nonfract"
	serviceResultTypeFract     = "fract"
	serviceResultTypeUndefined = "undefined"
)

type serviceResultType string

var _ def.CalcService = (*service)(nil)

type service struct {
	outerServices map[string]outerService
}

func NewService() *service {
	return &service{outerServices: make(map[string]outerService)}
}

type outerService struct {
	name       string
	conn       *grpc.ClientConn
	resultType serviceResultType
}
