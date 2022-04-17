package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	dataapi "HackDavisFood/services/data_api/internal/app/data_api"
)

const (
	recoveryMsg = "recovering from panic: %v"
)

var config struct {
	ServerURI string `env:"DATA_API_SERVICE_URI" envDefault:"0.0.0.0:50052"`
}

var (
	recoveryFunc grpc_recovery.RecoveryHandlerFunc
)

func init() {
	// parse environment config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

}

func main() {
	log.Println("in main")

	// start grpc server
	lis, err := net.Listen("tcp", config.ServerURI)
	if err != nil {
		panic(err)
	}

	recoveryFunc = func(p interface{}) error {
		m := fmt.Sprintf(recoveryMsg, p)
		err := status.Errorf(codes.Unknown, m)
		log.Println("Error: ", err)
		return err
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	// Register all services
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	dataapi.StartDataAPIServer(grpcServer)
	healthServer.SetServingStatus("data_api", grpc_health_v1.HealthCheckResponse_SERVING)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint
	grpcServer.GracefulStop()
}
