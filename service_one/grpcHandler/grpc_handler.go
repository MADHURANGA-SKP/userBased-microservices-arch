package grpchandler

import (
	pb "common/api/user_service/proto"
	"common/broker"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"service_one/types"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
	service types.UserService
	channel *amqp.Channel
}

func NewGRPCHandler(grpcServer *grpc.Server, service types.UserService, channel *amqp.Channel){
	handler := &grpcHandler{
		service: service,
		channel: channel,
	}

	pb.RegisterUserServiceServer(grpcServer, handler)
}

func(server *grpcHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest)(*pb.CreateUserResponse, error){
	q, err := server.channel.QueueDeclare(broker.UserCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	tr := otel.Tracer("amqp")
	amqpContext, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - publish - %s", q.Name))
	defer messageSpan.End()

	user, err := server.service.ValidateUser(amqpContext, req)
	if err != nil {
		return nil, err
	}

	o, err := server.service.CreateUser(amqpContext, user)
	if err != nil {
		return nil, err
	}

	marshalledOrder, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	// inject the headers
	headers := broker.InjectAMQPHeaders(amqpContext)

	server.channel.PublishWithContext(amqpContext, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         marshalledOrder,
		DeliveryMode: amqp.Persistent,
		Headers:      headers,
	})

	return o, nil
}

func(server *grpcHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest)(*pb.UpdateUserResponse, error){
	return server.service.UpdateUser(ctx, req)
}

func(server *grpcHandler) GetUser(ctx context.Context, req *pb.GetUserRequest)(*pb.GetUserResponse, error){
	return server.service.GetUser(ctx, req)
}


