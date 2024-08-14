package common

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)


type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPort string) error
	DeRegister(ctx context.Context, instanceID, serviceName string) error
	HealthCheck(ctx context.Context, instanceID, serviceName string) error
	Discover(ctx context.Context, serviceName string) ([]string,error)
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}