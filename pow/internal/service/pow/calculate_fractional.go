package pow

import (
	"context"
	"errors"

	"pow/internal/model"
)

func (s *service) CalculateFractional(ctx context.Context, input model.Fractionals, rounding uint32) (model.Fractional, error) {
	return model.Fractional{}, errors.New("not supported")
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
