package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rigmas/microservices/customer/handlers"
	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	s := handlers.Server{
		Token: "",
	}

	//register service to gRPC
	customer_grpc.RegisterCustomerServiceServer(grpcServer, &s)

	log.Println("customer service grpc is running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to rune due to %s\n", err)
	}
}
