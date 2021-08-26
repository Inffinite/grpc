package main

import (
	"context"
	"fmt"
	"grpc/calculator/calcpb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)

	// doUnary(c)
	// doServerStreaming(c)
	doErrorUnary(c)
}

func doServerStreaming(c calcpb.CalcServiceClient) {
	fmt.Println("[+] Starting Server Streaming RPC...")

	req := &calcpb.PrimeNumberDecomposeRequest{
		PrimeNumber: 120,
	}

	resStream, err := c.PrimeNumberDecompose(context.Background(), req)
	if err != nil {
		log.Fatalf("[*] Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("[*] Error while reading Stream: %v", err)
		}

		log.Printf("[+] Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("[+] Starting Unary rpc...")

	req := &calcpb.CalcRequest{
		Calculation: &calcpb.Calculation{
			FirstInt:  12,
			SecondInt: 9,
		},
	}
	res, err := c.Calc(context.Background(), req)

	if err != nil {
		log.Fatalf("[*] Error while calling Calculator RPC: %v", err)
	}

	log.Println(res.Result)
}

func doErrorUnary(c calcpb.CalcServiceClient) {
	fmt.Println("[+] Starting Squareroot Unary Error rpc...")

	// correct call
	doErrorCall(c, 34)

	// error call
	doErrorCall(c, -4)

}

func doErrorCall(c calcpb.CalcServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{
		Number: n,
	})

	if err != nil {
		respErr, ok := status.FromError(err)

		if ok {
			// actual error from gRPC (user error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("[-] We probably sent a negative number")
			}
		} else {
			log.Fatalf("[*] Big Error calling SquareRoot: %v", err)
		}
	} else {
		fmt.Printf("[+] SquareRoot of %v: %v", n, res.GetNumberRoot())
	}

}
