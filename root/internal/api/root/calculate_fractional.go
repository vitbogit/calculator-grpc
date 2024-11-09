package user

import (
	"context"

	"root/internal/converter"
	desc "root/pkg/root_v1"
)

func (i *Implementation) CalculateFractional(ctx context.Context, req *desc.CalculateFractionalRequest) (*desc.CalculateFractionalResponse, error) {
	output, err := i.sumService.CalculateFractional(ctx, *converter.ToFractionalsFromDesc(req), req.GetRounding())
	if err != nil {
		return nil, err
	}

	return &desc.CalculateFractionalResponse{
		C1: output.C1.String(),
		C2: output.C2.String(),
	}, nil
}
