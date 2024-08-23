package services

import (
	"common"
	pb "common/api/user_service/proto"
	"context"
	db "service_one/db/sqlc"
	"service_one/usergateway"
)

type Service struct {
	store db.Store
	gateway usergateway.UserGateway
}

func NewService(store db.Store, gateway usergateway.UserGateway) *Service {
	return &Service{store, gateway}
}

func (service *Service) CreateUser(ctx context.Context, p *pb.CreateUserRequest)(*pb.CreateUserResponse, error){
	o, err := service.store.CreateUser(ctx, db.CreateUserParams{
		FirstName: p.FirstName,
		LastName: p.LastName,
		Email: p.Email,
		UserName: p.UserName,
		Password: p.Password,
	})
	if err != nil {
		return nil, err
	}
	
	response := &pb.CreateUserResponse{
		User: &pb.User{
			FirstName: o.FirstName,
			LastName: o.LastName,
			Email: o.Email,
			UserName: o.UserName,
			Password: o.Password,
		},
	}

	return response, nil
}

func (service *Service) GetUser(ctx context.Context, p *pb.GetUserRequest)(*pb.GetUserResponse, error){
	o ,err := service.store.GetUser(ctx, p.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{User: &pb.User{UserID: o.UserID}}, nil
}

func (service *Service) UpdateUser(ctx context.Context, p *pb.UpdateUserRequest)(*pb.UpdateUserResponse, error){
	o ,err := service.store.UpdateUser(ctx, 
		db.UpdateUserParams{
			FirstName: p.FirstName,
			LastName: p.LastName,
			Email: p.Email,
			UserName: p.UserName,
		})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{User: &pb.User{
			FirstName: o.FirstName,
			LastName: o.LastName,
			Email: o.Email,
			UserName: o.UserName,
	}}, nil
}

func (service *Service) ValidateUser(ctx context.Context, p *pb.CreateUserRequest)(*pb.CreateUserRequest, error){
	if p == nil {
		return nil , common.ErrNoItems
	}

	//validation

	return p, nil
}