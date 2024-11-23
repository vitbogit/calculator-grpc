package user

import (
	"context"

	"mul/internal/converter"
	desc "mul/pkg/mul_v1"
)

func (i *Implementation) CalculateFractional(ctx context.Context, req *desc.CalculateFractionalRequest) (*desc.CalculateFractionalResponse, error) {
	output, err := i.mulService.CalculateFractional(ctx, *converter.ToFractionalsFromDesc(req), req.GetRounding())
	if err != nil {
		return nil, err
	}

	return &desc.CalculateFractionalResponse{
		C1: output.C1.Text('f', int(req.GetRounding())),
		C2: output.C2.Text('f', int(req.GetRounding())),
	}, nil
}
