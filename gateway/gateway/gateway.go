package gateway

import (
	pb "common/api/user_service/proto"
	"context"
)

type OMSGateway interface {
	CreateUser(ctx context.Context, p *pb.CreateUserRequest)(*pb.CreateUserResponse, error)
	// GetUser()
	// UpdateUser()

	// CreateOrder()
	// GetOrder()
	// UpdateOrder()

	// CreateOrder()
	// GetOrder()
	// UpdateOrder()
}