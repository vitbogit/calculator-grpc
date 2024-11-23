package sum

import (
	"calc/internal/model"
	desc_div "calc/pkg/div_v1"
	desc_mul "calc/pkg/mul_v1"
	desc_perc "calc/pkg/perc_v1"
	desc_pow "calc/pkg/pow_v1"
	desc_root "calc/pkg/root_v1"
	desc_sub "calc/pkg/sub_v1"
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

func SubService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	fmt.Println("calling sub service")

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
		service.conn, err = grpc.NewClient(os.Getenv(serviceSubAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_sub.NewSubClient(service.conn)

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
		resp, err := client.Calculate(context.Background(), &desc_sub.CalculateRequest{A: A, B: B, Rounding: rounding})
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

		resp, err := client.CalculateFractional(context.Background(), &desc_sub.CalculateFractionalRequest{A1: A1, A2: A2, B1: B1, B2: B2, Rounding: rounding})
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

func MulService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	fmt.Println("calling mul service")

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
		service.conn, err = grpc.NewClient(os.Getenv(serviceMulAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_mul.NewMulClient(service.conn)

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
		resp, err := client.Calculate(context.Background(), &desc_mul.CalculateRequest{A: A, B: B, Rounding: rounding})
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
		fmt.Println("calling with ", A1, A2, B1, B2)
		if err != nil {
			return nil, -1, fmt.Errorf("%s: on call number №%s", err.Error(), strconv.Itoa(*callNumber))
		}

		resp, err := client.CalculateFractional(context.Background(), &desc_mul.CalculateFractionalRequest{A1: A1, A2: A2, B1: B1, B2: B2, Rounding: rounding})
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

func DivService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	fmt.Println("calling div service")

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
		service.conn, err = grpc.NewClient(os.Getenv(serviceDivAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_div.NewDivClient(service.conn)

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
		resp, err := client.Calculate(context.Background(), &desc_div.CalculateRequest{A: A, B: B, Rounding: rounding})
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

		resp, err := client.CalculateFractional(context.Background(), &desc_div.CalculateFractionalRequest{A1: A1, A2: A2, B1: B1, B2: B2, Rounding: rounding})
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

func PercService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	fmt.Println("calling perc service")

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
		service.conn, err = grpc.NewClient(os.Getenv(servicePercAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_perc.NewPercClient(service.conn)

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
		resp, err := client.Calculate(context.Background(), &desc_perc.CalculateRequest{A: A, B: B, Rounding: rounding})
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

		resp, err := client.CalculateFractional(context.Background(), &desc_perc.CalculateFractionalRequest{A1: A1, A2: A2, B1: B1, B2: B2, Rounding: rounding})
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

func PowService(service *outerService, callNumber *int, call *model.CallRequest, lastCalcResponse *model.CalcResponse) (lc *model.CalcResponse, cl int, err error) {
	fmt.Println("calling pow service")

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
		service.conn, err = grpc.NewClient(os.Getenv(servicePowAdressEnvName), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("can`t connect to service: %s : %s\n", service.name, err.Error())
			return nil, -1, fmt.Errorf("can`t connect to service: %s", service.name)
		}
	}
	client := desc_pow.NewPowClient(service.conn)

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
		resp, err := client.Calculate(context.Background(), &desc_pow.CalculateRequest{A: A, B: B, Rounding: rounding})
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

		resp, err := client.CalculateFractional(context.Background(), &desc_pow.CalculateFractionalRequest{A1: A1, A2: A2, B1: B1, B2: B2, Rounding: rounding})
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
