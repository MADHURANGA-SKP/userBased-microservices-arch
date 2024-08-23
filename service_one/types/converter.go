package types

import (
	pb "common/api/user_service/proto"
	db "service_one/db/sqlc"
)

func Converter(user db.User) *pb.User {
	return &pb.User{
		UserID: user.UserID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		UserName: user.UserName,
		Password: user.Password,
	}
}

