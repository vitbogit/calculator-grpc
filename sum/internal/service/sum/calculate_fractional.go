package sum

import (
	"context"
	"fmt"
	"math/big"

	"sum/internal/model"
	"sum/pkg/math"
)

func (s *service) CalculateFractional(ctx context.Context, input model.Fractionals, rounding uint32) (model.Fractional, error) {
	var empty model.Fractional
	var c big.Float

	a1, a2, b1, b2 := input.A1, input.A2, input.B1, input.B2

	if a2.IsInt() && a2.Sign() == 0 || b2.IsInt() && b2.Sign() == 0 {
		return empty, fmt.Errorf("calculation failed: %w: %w", ErrInvalidInput, ErrDivisionByZero)
	}

	if !a2.IsInt() || !a1.IsInt() || !b1.IsInt() || !b2.IsInt() {

		c.Add(new(big.Float).Quo(a1, a2), new(big.Float).Quo(b1, b2))

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

	a1Int, _ := a1.Int64()
	a2Int, _ := a2.Int64()
	b1Int, _ := b1.Int64()
	b2Int, _ := b2.Int64()

	LCM, k1, k2 := math.LCMWithCoeffs(a2Int, b2Int)

	fmt.Println("debug a1 a2 b1 b2 ", a1Int, a2Int, b1Int, b2Int)
	fmt.Println("debug LCM ", LCM, k1, k2)

	c1Int := a1Int*k1 + b1Int*k2
	c2Int := LCM
	fmt.Println("debug c1 c2 ", c1Int, c2Int)
	c1c2gcd := math.GCD(c1Int, c2Int)
	c1Int = c1Int / c1c2gcd
	c2Int = c2Int / c1c2gcd
	fmt.Println("debug c1 c2 ", c1Int, c2Int)

	c1, c2 := new(big.Float), new(big.Float)
	c1.SetPrec(SetPrec).SetInt64(c1Int)
	c2.SetPrec(SetPrec).SetInt64(c2Int)

	return model.Fractional{
		C1: c1,
		C2: c2,
	}, nil
}
