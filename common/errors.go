package common

import "errors"

var (
	ErrToCreateUser =  errors.New("failed  to create user details")
	ErrToGetUser = errors.New("failed  to get user details")
	ErrToUpdateUser = errors.New("failed  to update user details")
	ErrToDeleteUser = errors.New("failed  to delete user details")

	ErrToCreateOrder =  errors.New("failed  to create order details")
	ErrToGetUOrder = errors.New("failed  to get order details")
	ErrToUpdateOrder = errors.New("failed  to update order details")
	ErrToDeleteOrder = errors.New("failed  to delete order details")

	ErrNoItems =  errors.New("items must have at least one item")
	ErrNoStock = errors.New("seleted items are not in the stock")
)