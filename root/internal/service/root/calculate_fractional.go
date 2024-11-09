package sum

import (
	"context"
	"fmt"
	"math/big"

	"root/internal/model"
	mymath "root/pkg/math"
)

func (s *service) CalculateFractional(ctx context.Context, input model.Fractionals, rounding uint32) (model.Fractional, error) {
	var empty model.Fractional
	var c big.Float

	a1, a2 := input.A1, input.A2

	if a1.Sign() < 0 {
		if a2.Sign() < 0 {
			a1.Neg(a1)
			a2.Neg(a2)
		} else {
			return model.Fractional{}, fmt.Errorf("calculation error: %w: %w", ErrInvalidInput, ErrNegativeNumber)
		}
	}

	if a2.IsInt() && a2.Sign() == 0 {
		return empty, fmt.Errorf("calculation failed: %w: %w", ErrInvalidInput, ErrDivisionByZero)
	}

	if !a2.IsInt() || !a1.IsInt() {

		c.Sqrt(new(big.Float).Quo(a1, a2))

		// c, err = math.ApplyPrecision(c, rounding)
		// if err != nil {
		// 	return empty, fmt.Errorf("calculation falied: %w: %w", ErrInvalidInput, err)
		// }

		one := new(big.Float)
		one.SetInt64(1)
		return model.Fractional{
			C1: &c,
			C2: one,
		}, nil
	}

	var c1Raw, c2Raw big.Float
	c1Raw.Sqrt(a1)
	c2Raw.Sqrt(a2)
	if !c1Raw.IsInt() || !c2Raw.IsInt() {

		c.Sqrt(new(big.Float).Quo(&c1Raw, &c2Raw))

		one := new(big.Float)
		one.SetInt64(1)
		return model.Fractional{
			C1: &c,
			C2: one,
		}, nil
	}

	c1Int, _ := c1Raw.Int64()
	c2Int, _ := c2Raw.Int64()
	c1c2gcd := mymath.GCD(c1Int, c2Int)
	c1Int = c1Int / c1c2gcd
	c2Int = c2Int / c1c2gcd

	c1, c2 := new(big.Float), new(big.Float)
	c1.SetInt64(c1Int)
	c2.SetInt64(c2Int)

	return model.Fractional{
		C1: c1,
		C2: c2,
	}, nil
}
