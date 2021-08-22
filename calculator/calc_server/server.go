package main

import (
	"context"
	"fmt"
	"grpc/calculator/calcpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("[+] Calc function was invoked with %v", req)
	first_int := req.GetCalculation().GetFirstInt()
	second_int := req.GetCalculation().GetSecondInt()

	result := first_int + second_int

	res := &calcpb.CalcResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeNumberDecompose(req *calcpb.PrimeNumberDecomposeRequest, stream calcpb.CalcService_PrimeNumberDecomposeServer) error {
	fmt.Printf("[+] GreetManyTimes function was invoked with %v", req)
	prime_number := req.GetPrimeNumber()

	k := int64(2)

	for prime_number > 1 {
		if prime_number%k == 0 {
			stream.Send(&calcpb.PrimeNumberDecomposeResponse{
				Result: k,
			})

			prime_number = prime_number / k
		} else {
			k++
			fmt.Printf("[+] Divisor incremented to %v", k)
		}
	}

	return nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("[+] Server open on port 50051")

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed open server: %v", err)
	}
}
