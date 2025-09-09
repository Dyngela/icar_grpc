package go_client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/icar_grpc/protos/gen"
)

// ClientServiceClient is a client for the ClientService
type ClientServiceClient struct {
	conn   *grpc.ClientConn
	client pb.ClientServiceClient
}

// NewClientServiceClient creates a new client for the ClientService
func NewClientServiceClient(serverAddr string) (*ClientServiceClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewClientServiceClient(conn)
	return &ClientServiceClient{
		conn:   conn,
		client: client,
	}, nil
}

// Close closes the connection to the server
func (c *ClientServiceClient) Close() error {
	return c.conn.Close()
}

// GetClients retrieves all clients
func (c *ClientServiceClient) GetClients(apiKey, empressa string) ([]*pb.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := &pb.GetClientsRequest{
		Base: &pb.BaseRequest{
			ApiKey:   apiKey,
			Empressa: empressa,
		},
	}

	resp, err := c.client.GetClients(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Base.Status != "success" {
		log.Printf("Error: %s - %s", resp.Base.ErrorId, resp.Base.ErrorMessage)
		return nil, nil
	}

	return resp.Clients, nil
}

// CreateClient creates a new client
func (c *ClientServiceClient) CreateClient(apiKey, empressa string, client *pb.Client) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Set last update date to now
	client.LastUpdateDate = timestamppb.Now()

	req := &pb.CreateClientRequest{
		Base: &pb.BaseRequest{
			ApiKey:   apiKey,
			Empressa: empressa,
		},
		Client: client,
	}

	resp, err := c.client.CreateClient(ctx, req)
	if err != nil {
		return "", err
	}

	if resp.Base.Status != "success" {
		log.Printf("Error: %s - %s", resp.Base.ErrorId, resp.Base.ErrorMessage)
		return "", nil
	}

	return resp.ClientId, nil
}

// UpdateClient updates an existing client
func (c *ClientServiceClient) UpdateClient(apiKey, empressa string, client *pb.Client) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Set last update date to now
	client.LastUpdateDate = timestamppb.Now()

	req := &pb.UpdateClientRequest{
		Base: &pb.BaseRequest{
			ApiKey:   apiKey,
			Empressa: empressa,
		},
		Client: client,
	}

	resp, err := c.client.UpdateClient(ctx, req)
	if err != nil {
		return false, err
	}

	if resp.Base.Status != "success" {
		log.Printf("Error: %s - %s", resp.Base.ErrorId, resp.Base.ErrorMessage)
		return false, nil
	}

	return resp.Success, nil
}
