package discovery

import (
	"context"
	"log"
	"math/rand"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ServiceConnection(
	ctx context.Context,
	serviceName string,
	registry Registry,

)(
	*grpc.ClientConn, 
	error,
){
	addr, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	
	log.Printf("Discovered %d instance of %s", len(addr), serviceName)

	//Randomly select an instance
	return grpc.Dial(
		addr[rand.Intn(len(addr))],
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
}
