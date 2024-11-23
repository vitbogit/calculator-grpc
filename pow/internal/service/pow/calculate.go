package pow

import (
	"context"

	"pow/internal/model"

	"github.com/ALTree/bigfloat"
)

const (
	SetPrec = 256
	MaxPrec = 50
)

func (s *service) Calculate(ctx context.Context, input model.Floats, rounding uint32) (model.Float, error) {
	//var c big.Float

	//c.SetPrec(math.MaxUint)

	// b, _ := input.B.Int64()
	// if b < 0 || b > 100 {
	// 	return model.Float{}, errors.New("percentage out of range")
	// }

	//fmt.Println("debug: sum/calc: ", input.A.Text('f', 50), input.B.Text('f', 50))
	// c.SetPrec(SetPrec).(input.A, input.B)

	// c.SetMode(2)
	// c.SetPrec(uint(rounding))

	// c, err = math.ApplyPrecision(c, rounding)
	// if err != nil {
	// 	return empty, fmt.Errorf("calculation falied: %w: %w", ErrInvalidInput, err)
	// }

	base := input.A     // 2.0
	exponent := input.B // 10.0

	output := model.Float{
		C: bigfloat.Pow(base, exponent),
	}

	return output, nil
}
