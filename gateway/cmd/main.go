package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	"github.com/rigmas/microservices/gateway/graph"
	"github.com/rigmas/microservices/gateway/graph/generated"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}

	//setup gRPC client connection
	var customerClient customer_grpc.CustomerServiceClient
	var productClient product_grpc.ProductServiceClient

	customerConnection, err := acquireConnection("customer")
	if err != nil {
		log.Fatalf("connection to customer_service failed")
	} else {
		customerClient = customer_grpc.NewCustomerServiceClient(customerConnection)
		log.Printf("connection to customer_service established")
		defer customerConnection.Close()
	}

	productConnection, err := acquireConnection("product")
	if err != nil {
		log.Fatalf("connection to product_service failed")
	} else {
		productClient = product_grpc.NewProductServiceClient(productConnection)
		log.Printf("connection to product_service established")
		defer productConnection.Close()
	}

	resolver := graph.Resolver{
		CustomerService: customerClient,
		ProductService:  productClient,
	}

	//inject gprc services into graphql
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver},
		),
	)

	r := chi.NewRouter()           //Initialize chi router
	r.Use(cors.AllowAll().Handler) //allow cors

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("Server is running on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
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
