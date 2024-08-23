package main

import (
	"common/broker"
	"common/discovery"
	"common/discovery/consul"
	"common/util"
	"context"
	"log"
	"net"
	db "service_one/db/sqlc"
	grpchandler "service_one/grpcHandler"
	"service_one/services"
	"service_one/usergateway"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main(){
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal(err)
	}

	registry, err := consul.NewRegistry(config.ConsulAddr, config.ServiceUser)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(config.ServiceUser)
	if err := registry.Register(ctx, instanceID, config.ServiceUser, config.GRPCAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(ctx, instanceID, config.ServiceUser); err != nil {
				log.Fatal("failed to check health")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	ch, close := broker.Connect(config.RMQUser, config.RMQPass, config.RMQHost, config.RMQPort)
	defer func ()  {
		close()
		ch.Close()
	} ()

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", config.GRPCAddr)
	if err != nil {
		log.Fatal("failed to listen", err)
	}

	defer l.Close()

	
	connPool, err := pgxpool.New(ctx, config.DBSorurce)
	if err != nil {
		log.Fatal(err)
	}

	util.RunDbMigrations(config.MigrationURL, config.DBSorurce)
	
	store := db.NewStore(connPool)

	gateway := usergateway.NewUserGateway(registry, config)
	
	svc := services.NewService(store, gateway)
	
	grpchandler.NewGRPCHandler(grpcServer, svc , ch)

	log.Printf("server started at port : %s", config.GRPCAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}

