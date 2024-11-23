package div

import (
	"context"
	"errors"
	"math/big"

	"div/internal/model"
)

const (
	SetPrec = 256
	MaxPrec = 50
)

func (s *service) Calculate(ctx context.Context, input model.Floats, rounding uint32) (model.Float, error) {
	var c big.Float

	//c.SetPrec(math.MaxUint)

	//fmt.Println("debug: sum/calc: ", input.A.Text('f', 50), input.B.Text('f', 50))
	if input.B.Sign() == 0 {
		return model.Float{}, errors.New("division by zero")
	}
	c.SetPrec(SetPrec).Quo(input.A, input.B)

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
