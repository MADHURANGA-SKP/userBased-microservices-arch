package discovery

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

type Registry struct {
	client *consul.Client
}

func NewRegistry(addr, serviceName string) (*Registry, error){
	config := consul.DefaultConfig()

	config.Address = addr

	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Registry{client}, err
}

func(r *Registry) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	host, portStr, found := strings.Cut(hostPort, ":")

	if !found {
		return errors.New("invalid host:port format. Eg: localhost:8081")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID: instanceID,
		Address: host,
		Port: port,
		Name: serviceName,
		Check: &consul.AgentServiceCheck{
			CheckID: instanceID,
			TLSSkipVerify: true,
			TTL: "5s",
			Timeout: "10s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})
}	

func(r *Registry) DeRegister(ctx context.Context, instanceID, serviceName string) error {
	log.Printf("Deregistering service %s", instanceID)
	return r.client.Agent().CheckDeregister(instanceID)
}	

func(r *Registry) HealthCheck(ctx context.Context, instanceID, serviceName string) error {
	return r.client.Agent().UpdateTTL(instanceID, "online", consul.HealthPassing)
}	

func(r *Registry) Discover(ctx context.Context, serviceName string) ([]string,error) {
	entries, _ , err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil ,err
	}

	var instance []string
	for _, entry := range entries {
		instance = append(instance, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
	}

	return instance, nil
}	