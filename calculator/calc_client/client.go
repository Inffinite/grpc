package main

import (
	"context"
	"fmt"
	"grpc/calculator/calcpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)

	doUnary(c)
	doServerStreaming(c)
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
