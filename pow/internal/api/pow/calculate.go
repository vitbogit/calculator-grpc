package user

import (
	"context"
	"errors"

	"pow/internal/converter"
	powService "pow/internal/service/pow"
	desc "pow/pkg/pow_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SetPrec = 256
	MaxPrec = 50
)

func (i *Implementation) Calculate(ctx context.Context, req *desc.CalculateRequest) (*desc.CalculateResponse, error) {
	output, err := i.powService.Calculate(ctx, *converter.ToFloatsFromDesc(req), req.GetRounding())
	if err != nil {
		if errors.Is(err, powService.ErrInvalidInput) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}

	return &desc.CalculateResponse{
		C: output.C.Text('f', int(req.GetRounding())),
	}, nil
}
