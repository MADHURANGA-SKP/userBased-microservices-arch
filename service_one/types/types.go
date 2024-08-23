package types

import (
	pb "common/api/user_service/proto"
	"context"
)

type UserService interface {
	CreateUser(context.Context,*pb.CreateUserRequest)(*pb.CreateUserResponse, error)
	ValidateUser(context.Context,*pb.CreateUserRequest)(*pb.CreateUserRequest, error)
	UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	GetUser(context.Context,*pb.GetUserRequest)(*pb.GetUserResponse, error)
}

type UserStore interface {
	Create(context.Context, CreateUserRequest)( CreateUserResponse, error)
	Get(context.Context, GetUserRequest)( GetUserResponse, error)
	Update(context.Context, UpdateUserRequest)( UpdateUserResponse, error)
}

type CreateUserRequest struct{
	pb.CreateUserRequest
}

type CreateUserResponse struct{
	pb.CreateUserResponse
}

type GetUserRequest struct {
	pb.GetUserRequest
}

type GetUserResponse struct {
	pb.GetUserResponse
}

type UpdateUserRequest struct {
	pb.UpdateUserRequest
}

type UpdateUserResponse struct {
	pb.UpdateUserResponse
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email 	  string `json:"email"`
	Username  string `json:"user_name"`
	Password  string `json:"password"`
}