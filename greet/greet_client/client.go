package main

import (
	// "fmt"
	"fmt"
	"grpc/greet/greetpb"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}

	// defer makes this code run at the end of the program
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client: %f", c)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBidiStream(c)
}

func doBidiStream(c greetpb.GreetServiceClient) {
	fmt.Println("[+] Starting to do a BiDi Streaming RPC...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucky",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucrisha",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Luana",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lula",
			},
		},

		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Luna",
			},
		},
	}

	waitc := make(chan struct{})
	go func() {
		// func to send messages
		for _, req := range requests {
			fmt.Printf("\n[+] Sending message: %v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		// func to receive messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("[*] Error while receiving: %v", err)
				break
			}

			fmt.Printf("\n[+] Received: %v\n", res.GetResult())
		}
		close(waitc)

	}()

	// block until everything is done
	<-waitc
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("[+] Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucky",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucrisha",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Luana",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lula",
			},
		},

		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Luna",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("[*] Error while calling LongGreet: %v\n", err)
	}

	for _, req := range requests {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[*] Error while receiving response from LongGreet: %v\n", err)
	}

	fmt.Printf("[+] LongGreet response: %v\n", res)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("[+] Starting Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lola",
			LastName:  "Loud",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("[*] Error while calling GreetManyTimes RPC: %v\n", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("[*] Error while reading Stream: %v\n", err)
		}

		log.Printf("[+] Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("[+] Starting Unary rpc...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lewis",
			LastName:  "Me",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("[**] Error while calling Greet RPC: %v\n", err)
	}

	log.Println(res.Result)
}
