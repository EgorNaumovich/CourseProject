package main

import (
	protobuf "awesomeProject6/delivery-service/proto/delivery"
	tr "awesomeProject6/transport-service/proto/transport"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

type repository interface {
	Create(delivery *protobuf.Delivery) (*protobuf.Delivery, error)
	GetAll() []*protobuf.Delivery
}

type Repository struct {
	mtx        sync.RWMutex
	deliveries []*protobuf.Delivery
}

func (repo *Repository) Create(delivery *protobuf.Delivery) (*protobuf.Delivery, error) {
	repo.mtx.Lock()
	updated := append(repo.deliveries, delivery)
	repo.deliveries = updated
	repo.mtx.Unlock()
	return delivery, nil
}

func (repo *Repository) GetAll() []*protobuf.Delivery {
	return repo.deliveries
}

type service struct {
	repo repository
}

type description struct {
	capacity int32
	weight   int32
}

func (s *service) CreateDelivery(ctx context.Context, req *protobuf.Delivery) (*protobuf.Response, error) {

	// Save our consignment
	delivery, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &protobuf.Response{Created: true, Delivery: delivery}, nil
}

func (s *service) GetDeliveries(ctx context.Context, req *protobuf.GetRequest) (*protobuf.Response, error) {
	deliveries := s.repo.GetAll()
	return &protobuf.Response{Deliveries: deliveries}, nil
}

func main() {

	repo := &Repository{}

	// Set-up our gRPC server.
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	srv := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	protobuf.RegisterDeliveryServiceServer(srv, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(srv)

	log.Println("Server", ":8080")
	if err := srv.Serve(l); err != nil {
		log.Fatalf("Serve error: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()
	client := tr.NewTransportServiceClient(conn)
	filter := &description
	{
		len(repo.deliveries[0].Containers),

	}
	gettransport, err := client.Available(context.Background(), repo.deliveries[0].Description)

}
