package main

import (
	"common"
	"common/discovery"
	"common/discovery/consul"
	"common/util"
	"context"
	"gateway/gateway"
	"log"
	"net/http"
	"time"
)

func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	err = common.SetGlobalTracer(context.TODO(), config.ServiceGateway, config.JeagerAddr)
	if err != nil {
		log.Fatal(err)
	}

	registry, err := consul.NewRegistry(config.ConsulAddr, config.ServiceGateway)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(config.ServiceGateway)
	if err := registry.Register(ctx, instanceID, config.ServiceGateway, config.HttpAddr); err != nil {
		panic(err)
	}
	
	go func() {
		for {
			if err := registry.HealthCheck(ctx, instanceID, config.ServiceGateway); err != nil {
				log.Fatal("failed to check health")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.DeRegister(ctx, instanceID, config.ServiceGateway)

	mux := http.NewServeMux()

	OMSGateway := gateway.NewGRPCGateway(registry)

	handler := NewHandler(OMSGateway)

	handler.registerRoutes(mux)

	log.Printf("stating http server at %s", config.HttpAddr)

	if err := http.ListenAndServe(config.HttpAddr, mux); err != nil {
		log.Fatalf("failed to start http server at %s:", config.ConsulAddr)
	}
}