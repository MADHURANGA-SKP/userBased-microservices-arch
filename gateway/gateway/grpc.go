package gateway

import (
	pb "common/api/user_service/proto"
	"common/discovery"
	util "common/util"
	"context"
	"log"
)

type gateway struct {
	registry discovery.Registry
	config util.Config
}

func NewGRPCGateway(registry discovery.Registry, config util.Config) *gateway {
	return &gateway{registry, config}
}

func (g *gateway) CreateUser(ctx context.Context, p *pb.CreateUserRequest)(*pb.CreateUserResponse, error) {
	conn, err := discovery.ServiceConnection(context.Background(), g.config.ServiceUser, g.registry)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}

	c := pb.NewUserServiceClient(conn)

	reponse, err := c.CreateUser(ctx, &pb.CreateUserRequest{
		FirstName: p.FirstName,
		LastName: p.LastName,
		Email: p.Email,
		UserName: p.UserName,
		Password: p.Password,
	})
	if err != nil {
		return nil, err
	}

	return reponse, nil
}
