package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rigmas/microservices/product/handlers"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	s := handlers.Server{
		Products: "",
	}

	//register service to gRPC
	product_grpc.RegisterProductServiceServer(grpcServer, &s)

	log.Println("product service grpc is running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to rune due to %s\n", err)
	}
}
