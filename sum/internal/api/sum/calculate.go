package user

import (
	"context"
	"errors"

	"sum/internal/converter"
	sumService "sum/internal/service/sum"
	desc "sum/pkg/sum_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Calculate(ctx context.Context, req *desc.CalculateRequest) (*desc.CalculateResponse, error) {
	output, err := i.sumService.Calculate(ctx, *converter.ToFloatsFromDesc(req), req.GetRounding())
	if err != nil {
		if errors.Is(err, sumService.ErrInvalidInput) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}

	return &desc.CalculateResponse{
		C: output.C.Text('f', int(req.GetRounding())),
	}, nil
}
