package main

import (
	"common"
	"common/discovery"
	"common/discovery/consul"
	"common/util"
	"context"
	"gateway/gateway"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	config, err := util.LoadConfig("..")
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

	router := gin.Default()

	OMSGateway := gateway.NewGRPCGateway(registry, config)

	handler := NewHandler(OMSGateway)

	handler.registerRoutes(router)

	log.Printf("Starting HTTP server at %s", config.HttpAddr)

	// Start the Gin server
	if err := router.Run(config.HttpAddr); err != nil {
		log.Fatal("Failed to start HTTP server")
	}
}