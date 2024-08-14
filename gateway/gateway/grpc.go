package gateway

import (
	pb "common/api/user_service/proto"
	"common/discovery"
	"context"
	"log"
)

type gateway struct {
	registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
	return &gateway{registry}
}

func (g *gateway) CreateUser(ctx context.Context, p *pb.CreateUserRequest)(*pb.User, error) {
	conn, err := discovery.ServiceConnection(context.Background(), "users", g.registry)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	return c.CreateUser(ctx, &pb.CreateUserRequest{
		FirstName: p.FirstName,
		LastName: p.LastName,
		Email: p.Email,
		UserName: p.UserName,
		Password: p.Password,
	})
}
