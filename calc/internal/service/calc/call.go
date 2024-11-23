package sum

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"calc/internal/model"
)

const (
	SetPrec = 256
	MaxPrec = 50
)

func (s *service) Call(ctx context.Context, call model.CallRequest) (*model.CalcResponse, error) {
	fmt.Println("log: -----------------------------------------------")
	var err error

	if call.Services == nil || call.Services.Services == nil {
		return nil, fmt.Errorf("can`t construct response: %w", ErrEmptySequence)
	}

	services := call.Services.Services
	//outputType := serviceResultTypeUndefined

	lastCalcResponse := &model.CalcResponse{}
	callNumber := 0
	for _, serviceName := range services {

		// Prepare service
		if _, ok := s.outerServices[serviceName]; !ok {
			s.outerServices[serviceName] = outerService{name: serviceName}
		}
		service := s.outerServices[serviceName]

		// Last call result for first call
		fmt.Println("CALL ", callNumber)

		switch service.name {
		case serviceSumName:
			lastCalcResponse, callNumber, err = SumService(&service, &callNumber, &call, lastCalcResponse)

		case serviceRootName:
			lastCalcResponse, callNumber, err = RootService(&service, &callNumber, &call, lastCalcResponse)
		}
		if err != nil {
			return nil, fmt.Errorf("service err: %s", err.Error())
		}
		// callNumber += 1
		fmt.Printf("debug: %v\n", lastCalcResponse)

	}

	// if outputType == serviceResultTypeUndefined {
	// 	return nil, fmt.Errorf("can`t construct response: %w", ErrEmptySequence)
	// }

	return lastCalcResponse, nil
}

func GetNumbers(callNumber int, calcRequests []*model.CalcRequest) (A, A1, A2 string, nf bool, rounding uint32, err error) {
	if calcRequests == nil || len(calcRequests) < 1 || len(calcRequests) < callNumber {
		return "", "", "", false, 0, fmt.Errorf("not provided enough numbers")
	}

	if calcRequests[callNumber].CalcNFRequest != nil {
		if calcRequests[callNumber].CalcNFRequest.A == nil {
			return "", "", "", false, 0, fmt.Errorf("bad number №%s", strconv.Itoa(callNumber))
		}
		A = calcRequests[callNumber].CalcNFRequest.A.Text('f', MaxPrec)
		nf = true
		rounding = calcRequests[callNumber].CalcNFRequest.Rounding
	} else if calcRequests[callNumber].CalcFRequest != nil {
		if calcRequests[callNumber].CalcFRequest.A1 == nil {
			return "", "", "", false, 0, fmt.Errorf("bad number №%s", strconv.Itoa(callNumber))
		}
		if calcRequests[callNumber].CalcFRequest.A2 == nil {
			return "", "", "", false, 0, fmt.Errorf("bad number №%s", strconv.Itoa(callNumber))
		}
		A1 = calcRequests[callNumber].CalcFRequest.A1.Text('f', MaxPrec)
		A2 = calcRequests[callNumber].CalcFRequest.A2.Text('f', MaxPrec)
		rounding = calcRequests[callNumber].CalcFRequest.Rounding
		nf = false
	} else {
		return "", "", "", false, 0, fmt.Errorf("missing number №%s", strconv.Itoa(callNumber))
	}

	return A, A1, A2, nf, rounding, nil
}

func GetRounding(callNumber int, calcRequests []*model.CalcRequest) (rounding uint32, err error) {
	if calcRequests == nil || len(calcRequests) < 1 || len(calcRequests) < callNumber {
		return 0, fmt.Errorf("not provided enough numbers")
	}

	if calcRequests[callNumber].CalcNFRequest != nil {
		rounding = calcRequests[callNumber].CalcNFRequest.Rounding
	} else if calcRequests[callNumber].CalcFRequest != nil {
		rounding = calcRequests[callNumber].CalcFRequest.Rounding
	} else {
		return 0, fmt.Errorf("missing number №%s", strconv.Itoa(callNumber))
	}

	return rounding, nil
}

func GetCallResponse(A, A1, A2 string, Anf bool) *model.CalcResponse {
	if Anf {
		num := new(big.Float)
		num.SetPrec(SetPrec).SetString(A)
		return &model.CalcResponse{CalcNFResponse: &model.CalcNFResponse{C: num}}
	} else {
		num1 := new(big.Float)
		num1.SetPrec(SetPrec).SetString(A1)
		num2 := new(big.Float)
		num2.SetPrec(SetPrec).SetString(A2)
		return &model.CalcResponse{CalcFResponse: &model.CalcFResponse{C1: num1, C2: num2}}
	}
}

func CallNFResponseToStr(lastCalcResponse, newArgs *model.CalcResponse) (A, B string, err error) {
	if lastCalcResponse.CalcNFResponse != nil && lastCalcResponse.CalcNFResponse.C != nil {
		A = lastCalcResponse.CalcNFResponse.C.Text('f', MaxPrec)
	} else if lastCalcResponse.CalcFResponse != nil && lastCalcResponse.CalcFResponse.C1 != nil && lastCalcResponse.CalcFResponse.C2 != nil {
		z := new(big.Float).Quo(lastCalcResponse.CalcFResponse.C1, lastCalcResponse.CalcFResponse.C2)
		A = z.Text('f', MaxPrec)
	} else {
		return "", "", fmt.Errorf("invalid last response")
	}

	if newArgs.CalcNFResponse != nil && newArgs.CalcNFResponse.C != nil {
		B = newArgs.CalcNFResponse.C.Text('f', MaxPrec)
	} else if newArgs.CalcFResponse != nil && newArgs.CalcFResponse.C1 != nil && newArgs.CalcFResponse.C2 != nil {
		z := new(big.Float).Quo(newArgs.CalcFResponse.C1, newArgs.CalcFResponse.C2)
		B = z.Text('f', MaxPrec)
	} else {
		return "", "", fmt.Errorf("invalid args")
	}

	return A, B, nil
}

func CallNFResponseToStrSingle(lastCalcResponse *model.CalcResponse) (A string, err error) {
	if lastCalcResponse.CalcNFResponse != nil && lastCalcResponse.CalcNFResponse.C != nil {
		A = lastCalcResponse.CalcNFResponse.C.Text('f', MaxPrec)
	} else if lastCalcResponse.CalcFResponse != nil && lastCalcResponse.CalcFResponse.C1 != nil && lastCalcResponse.CalcFResponse.C2 != nil {
		z := new(big.Float).Quo(lastCalcResponse.CalcFResponse.C1, lastCalcResponse.CalcFResponse.C2)
		A = z.Text('f', MaxPrec)
	} else {
		return "", fmt.Errorf("invalid last response")
	}

	return A, nil
}

func CallFResponseToStr(lastCalcResponse, newArgs *model.CalcResponse) (A1, A2, B1, B2 string, err error) {
	if lastCalcResponse.CalcFResponse != nil && lastCalcResponse.CalcFResponse.C1 != nil && lastCalcResponse.CalcFResponse.C2 != nil {
		A1 = lastCalcResponse.CalcFResponse.C1.Text('f', MaxPrec)
		A2 = lastCalcResponse.CalcFResponse.C2.Text('f', MaxPrec)
	} else if lastCalcResponse.CalcNFResponse != nil && lastCalcResponse.CalcNFResponse.C != nil {
		A1 = lastCalcResponse.CalcNFResponse.C.Text('f', MaxPrec)
		A2 = "1"
	} else {
		return "", "", "", "", fmt.Errorf("invalid last response")
	}

	if newArgs.CalcFResponse != nil && newArgs.CalcFResponse.C1 != nil && newArgs.CalcFResponse.C2 != nil {
		B1 = newArgs.CalcFResponse.C1.Text('f', MaxPrec)
		B2 = newArgs.CalcFResponse.C2.Text('f', MaxPrec)
	} else if newArgs.CalcNFResponse != nil && newArgs.CalcNFResponse.C != nil {
		B1 = newArgs.CalcNFResponse.C.Text('f', MaxPrec)
		B2 = "1"
	} else {
		return "", "", "", "", fmt.Errorf("invalid args")
	}

	return A1, A2, B1, B2, nil
}

func CallFResponseToStrSingle(lastCalcResponse *model.CalcResponse) (A1, A2 string, err error) {
	if lastCalcResponse.CalcFResponse != nil && lastCalcResponse.CalcFResponse.C1 != nil && lastCalcResponse.CalcFResponse.C2 != nil {
		A1 = lastCalcResponse.CalcFResponse.C1.Text('f', MaxPrec)
		A2 = lastCalcResponse.CalcFResponse.C2.Text('f', MaxPrec)
	} else if lastCalcResponse.CalcNFResponse != nil && lastCalcResponse.CalcNFResponse.C != nil {
		A1 = lastCalcResponse.CalcNFResponse.C.Text('f', MaxPrec)
		A2 = "1"
	} else {
		return "", "", fmt.Errorf("invalid last response")
	}

	return A1, A2, nil
}
