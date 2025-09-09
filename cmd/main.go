package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Dyngela/icar_grpc/protos/gen"
	"google.golang.org/grpc"
)

// Implement the ClientServiceServer interface
type clientServiceServer struct {
	pb.UnimplementedClientServiceServer
}

// GetClients implements the GetClients RPC
func (s *clientServiceServer) GetClients(ctx context.Context, req *pb.GetClientsRequest) (*pb.ClientsResponse, error) {
	log.Printf("Received GetClients request: %+v", req.Base)

	// Dummy response for now
	clients := []*pb.Client{
		{
			InternalClientId: "123",
			CompanyName:      "Acme Corp",
			IsCompany:        true,
			IsActive:         true,
		},
	}

	resp := &pb.ClientsResponse{
		Base:    &pb.BaseResponse{Status: "success"},
		Clients: clients,
	}

	return resp, nil
}

// CreateClient implements the CreateClient RPC
func (s *clientServiceServer) CreateClient(ctx context.Context, req *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	log.Printf("Received CreateClient request: %+v", req.Client)
	// For demo, pretend we created a client
	return &pb.CreateClientResponse{
		Base:     &pb.BaseResponse{Status: "OK"},
		ClientId: "new-client-id",
	}, nil
}

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	pb.RegisterClientServiceServer(grpcServer, &clientServiceServer{})

	log.Printf("gRPC server listening on port %d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
