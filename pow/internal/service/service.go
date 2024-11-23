package service

import (
	"context"
	"pow/internal/model"
)

type PowService interface {
	Calculate(ctx context.Context, terms model.Floats, rounding uint32) (model.Float, error)
	CalculateFractional(ctx context.Context, input model.Fractionals, rounding uint32) (model.Fractional, error)
}
