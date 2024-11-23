package service

import (
	"calc/internal/model"
	"context"
)

type CalcService interface {
	Call(ctx context.Context, call model.CallRequest) (*model.CalcResponse, error)
}
