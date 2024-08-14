package main

import pb "common/api/user_service/proto"

type CreateUserRequest struct {
	User *pb.User `json:"user"`
	// UserID int64 `json:"user_id"`
	// FirstName string `json:"first_name"`
	// LastName string `json:"last_name"`
	// Email string `json:"email"`
	// UserName string `json:"user_name"`
	// Password string `json:"password"`
}