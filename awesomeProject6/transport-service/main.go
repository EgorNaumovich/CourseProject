package main

import (
	protobuf "awesomeProject6/transport-service/proto/transport"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type repository interface {
	Available(*protobuf.Description) (*protobuf.Transport, error)
}

type TransportRepository struct {
	transports []*protobuf.Transport
}

// Available  - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *TransportRepository) Available(dscr *protobuf.Description) (*protobuf.Transport, error) {
	for _, transport := range repo.transports {
		if dscr.ContainerCapacity <= transport.ContainerCapacity && dscr.Weight <= transport.Weight {
			return transport, nil
		}
	}
	return nil, errors.New("no transport available")
}

type transportService struct {
	repo repository
}

func (s *transportService) Available(ctx context.Context, req *protobuf.Description)  (res *protobuf.Response, err error) {

	// Find the next available vessel
	transport, err := s.repo.Available(req)
	if err != nil {
		return nil, err
	}

	// Set the vessel as part of the response message type
	res.Transport = transport
	return res, nil
}

func main() {
	transports := []*protobuf.Transport{
		{Id: "transport1", Name: "Name1", Weight: 20000, ContainerCapacity: 250},
	}
	repo := &TransportRepository{transports}

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
	srv := grpc.NewServer()

	protobuf.RegisterTransportServiceServer(srv, &transportService{repo})

	// Register reflection service on gRPC server.
	reflection.Register(srv)

	log.Println("Server", ":9090")
	if err := srv.Serve(l); err != nil {
		log.Fatalf("Serve error: %v", err)
	}

}
