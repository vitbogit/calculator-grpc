package user

import (
	"context"
	"errors"

	"calc/internal/converter"
	sumService "calc/internal/service/calc"
	desc "calc/pkg/calc_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SetPrec = 256
	MaxPrec = 50
)

func (i *Implementation) Call(ctx context.Context, req *desc.CallRequest) (*desc.CallResponse, error) {

	calcResp, err := i.sumService.Call(ctx, *converter.ToCallRequestFromDesc(req))
	if err != nil {
		if errors.Is(err, sumService.ErrInvalidInput) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}

	switch {
	case calcResp == nil:
		return nil, status.Error(codes.Internal, errors.New("failed getting response").Error())
	case calcResp.CalcNFResponse != nil && calcResp.CalcNFResponse.C != nil:
		return &desc.CallResponse{
			CalcResponse: &desc.CalcResponse{
				CalcNFResponse: &desc.CalcNFResponse{
					C: calcResp.CalcNFResponse.C.Text('f', int(calcResp.Precise)),
				},
			},
		}, nil
	case calcResp.CalcFResponse != nil && calcResp.CalcFResponse.C1 != nil && calcResp.CalcFResponse.C2 != nil:
		return &desc.CallResponse{
			CalcResponse: &desc.CalcResponse{
				CalcFResponse: &desc.CalcFResponse{
					C1: calcResp.CalcFResponse.C1.Text('f', int(calcResp.Precise)),
					C2: calcResp.CalcFResponse.C2.Text('f', int(calcResp.Precise)),
				},
			},
		}, nil
	default:
		return nil, status.Error(codes.Internal, errors.New("haven`t got any response").Error())
	}
}
