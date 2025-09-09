package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/Dyngela/icar_grpc/protos/gen"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewClientServiceClient(conn)

	resp, err := client.GetClients(context.Background(), &pb.GetClientsRequest{
		Base: &pb.BaseRequest{
			ApiKey:   "dummy-key",
			Empressa: "dealer-uuid",
		},
	})
	if err != nil {
		log.Fatalf("GetClients failed: %v", err)
	}

	for _, c := range resp.Clients {
		fmt.Printf("Client: %s %s (%s)\n", c.FirstName, c.LastName, c.CompanyName)
	}
}
