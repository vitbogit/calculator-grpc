package sum

import (
	"context"
	"math/big"

	"sum/internal/model"
)

func (s *service) Calculate(ctx context.Context, input model.Floats, rounding uint32) (model.Float, error) {
	var c big.Float

	//c.SetPrec(math.MaxUint)

	c.Add(input.A, input.B)

	// c.SetMode(2)
	// c.SetPrec(uint(rounding))

	// c, err = math.ApplyPrecision(c, rounding)
	// if err != nil {
	// 	return empty, fmt.Errorf("calculation falied: %w: %w", ErrInvalidInput, err)
	// }

	output := model.Float{
		C: &c,
	}

	return output, nil
}
