package sum

import (
	"calc/internal/model"
	desc_root "calc/pkg/root_v1"
	desc_sum "calc/pkg/sum_v1"
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SumService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	if *callNumber == 0 {
		A, A1, A2, Anf, _, err := GetNumbers(*callNumber, call.CalcRequests)
		if err != nil {
			return nil, -1, fmt.Errorf("can`t read number: %s", err.Error())
		}
		lastCalcResponse = GetCallResponse(A, A1, A2, Anf)
		fmt.Println("log: start:", A, A1, A2, Anf)
		*callNumber += 1
	}

	if len(call.CalcRequests) <= *callNumber {
		return nil, -1, fmt.Errorf("not enough args")
	}

	// Prepare client (sum)
	if service.conn == nil {
		service.conn, err = grpc.NewClient(os.Getenv(serviceSumAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_sum.NewSumClient(service.conn)

	// Prepare new args for new request (Besides result of last request)
	B, B1, B2, Bnf, rounding, err := GetNumbers(*callNumber, call.CalcRequests)
	if err != nil {
		return nil, -1, fmt.Errorf("can`t read number: %s", err.Error())
	}
	newArgs := GetCallResponse(B, B1, B2, Bnf)

	fmt.Println("log: new args:", B, B1, B2, strconv.Itoa(int(rounding)), Bnf)

	if Bnf {

		A, B, err := CallNFResponseToStr(lastCalcResponse, newArgs)
		if err != nil {
			return nil, -1, fmt.Errorf("%s: on call number №%s", err.Error(), strconv.Itoa(*callNumber))
		}

		//fmt.Println("debug:", A, B)
		resp, err := client.Calculate(context.Background(), &desc_sum.CalculateRequest{A: A, B: B, Rounding: rounding})
		if err != nil {
			fmt.Printf("got error from service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("got error from service: %s", service.name)
		}

		num := new(big.Float)
		num.SetPrec(SetPrec).SetString(resp.C)
		fmt.Println("log: result:", resp.C)
		lastCalcResponse.CalcNFResponse = &model.CalcNFResponse{C: num}
		lastCalcResponse.CalcFResponse = nil
		lastCalcResponse.Precise = rounding
	} else {

		A1, A2, B1, B2, err := CallFResponseToStr(lastCalcResponse, newArgs)
		if err != nil {
			return nil, -1, fmt.Errorf("%s: on call number №%s", err.Error(), strconv.Itoa(*callNumber))
		}

		resp, err := client.CalculateFractional(context.Background(), &desc_sum.CalculateFractionalRequest{A1: A1, A2: A2, B1: B1, B2: B2, Rounding: rounding})
		if err != nil {
			fmt.Printf("got error from service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("got error from service: %s", service.name)
		}

		num1 := new(big.Float)
		num1.SetPrec(SetPrec).SetString(resp.C1)
		num2 := new(big.Float)
		num2.SetPrec(SetPrec).SetString(resp.C2)
		fmt.Println("log: result:", resp.C1, resp.C2)
		lastCalcResponse.CalcFResponse = &model.CalcFResponse{C1: num1, C2: num2}
		lastCalcResponse.CalcNFResponse = nil
		lastCalcResponse.Precise = rounding
	}
	*callNumber += 1
	return lastCalcResponse, *callNumber, nil
}

func RootService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	var rounding uint32

	// if len(call.CalcRequests) <= *callNumber {
	// 	return nil, -1, fmt.Errorf("not enough args")
	// }

	if *callNumber == 0 {
		A, A1, A2, Anf, r, err := GetNumbers(*callNumber, call.CalcRequests)
		if err != nil {
			return nil, -1, fmt.Errorf("can`t read number: %s", err.Error())
		}
		lastCalcResponse = GetCallResponse(A, A1, A2, Anf)
		fmt.Println("log: start:", A, A1, A2, Anf)
		rounding = r // call.CalcRequests[*callNumber].CalcNFRequest.Rounding
		*callNumber += 1
	} else {
		rounding = lastCalcResponse.Precise
	}
	// else {
	// 	// rounding, err = GetRounding(*callNumber, call.CalcRequests)
	// 	if err != nil {
	// 		return nil, -1, fmt.Errorf("can`t read number: %s", err.Error())
	// 	}
	// }

	// fmt.Println("debug ", len(call.CalcRequests), *callNumber)

	// Prepare client (sum)
	if service.conn == nil {
		service.conn, err = grpc.NewClient(os.Getenv(serviceRootAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_root.NewRootClient(service.conn)

	if lastCalcResponse.CalcNFResponse != nil {

		A, err := CallNFResponseToStrSingle(lastCalcResponse)
		if err != nil {
			return nil, -1, fmt.Errorf("%s: on call number №%s", err.Error(), strconv.Itoa(*callNumber))
		}

		//fmt.Println("debug:", A, B)
		resp, err := client.Calculate(context.Background(), &desc_root.CalculateRequest{A: A, Rounding: rounding})
		if err != nil {
			fmt.Printf("got error from service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("got error from service: %s", service.name)
		}

		num := new(big.Float)
		num.SetPrec(SetPrec).SetString(resp.C)
		fmt.Println("log: result:", resp.C)
		lastCalcResponse.CalcNFResponse = &model.CalcNFResponse{C: num}
		lastCalcResponse.CalcFResponse = nil
		lastCalcResponse.Precise = rounding
	} else {

		A1, A2, err := CallFResponseToStrSingle(lastCalcResponse)
		if err != nil {
			return nil, -1, fmt.Errorf("%s: on call number №%s", err.Error(), strconv.Itoa(*callNumber))
		}

		resp, err := client.CalculateFractional(context.Background(), &desc_root.CalculateFractionalRequest{A1: A1, A2: A2, Rounding: rounding})
		if err != nil {
			fmt.Printf("got error from service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("got error from service: %s", service.name)
		}

		num1 := new(big.Float)
		num1.SetPrec(SetPrec).SetString(resp.C1)
		num2 := new(big.Float)
		num2.SetPrec(SetPrec).SetString(resp.C2)
		fmt.Println("log: result:", resp.C1, resp.C2)
		lastCalcResponse.CalcFResponse = &model.CalcFResponse{C1: num1, C2: num2}
		lastCalcResponse.CalcNFResponse = nil
		lastCalcResponse.Precise = rounding
	}

	fmt.Printf("debug: %v\n", lastCalcResponse)

	return lastCalcResponse, *callNumber, nil
}
