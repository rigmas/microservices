package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rigmas/microservices/order/handlers"
	"github.com/rigmas/microservices/order/handlers/order_grpc"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
	"google.golang.org/grpc"
)

func main() {
	var productClient product_grpc.ProductServiceClient

	productConnection, err := acquireConnection("product")
	if err != nil {
		log.Fatalf("connection to product_service failed")
	} else {
		productClient = product_grpc.NewProductServiceClient(productConnection)
		log.Printf("connection to product_service established")
		defer productConnection.Close()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	s := handlers.Server{
		Orders:        "",
		ProductClient: productClient,
	}

	//register service to gRPC
	order_grpc.RegisterOrderServiceServer(grpcServer, &s)

	log.Println("order service grpc is running")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to rune due to %s\n", err)
	}
}

func acquireConnection(serviceName string) (*grpc.ClientConn, error) {
	dialCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		dialCtx,
		fmt.Sprintf("%s:9000", serviceName),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
