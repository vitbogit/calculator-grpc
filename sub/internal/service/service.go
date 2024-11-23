package service

import (
	"context"
	"sub/internal/model"
)

type SubService interface {
	Calculate(ctx context.Context, terms model.Floats, rounding uint32) (model.Float, error)
	CalculateFractional(ctx context.Context, input model.Fractionals, rounding uint32) (model.Fractional, error)
}
