package main

import (
	"fmt"
	pb "github.com/JMURv/unona/ratings/api/pb"
	kafka "github.com/JMURv/unona/ratings/internal/broker/kafka"
	redis "github.com/JMURv/unona/ratings/internal/cache/redis"
	ctrl "github.com/JMURv/unona/ratings/internal/controller/rating"
	handler "github.com/JMURv/unona/ratings/internal/handler/grpc"
	tracing "github.com/JMURv/unona/ratings/internal/metrics/jaeger"
	metrics "github.com/JMURv/unona/ratings/internal/metrics/prometheus"

	//mem "github.com/JMURv/unona/ratings/internal/repository/memory"
	db "github.com/JMURv/unona/ratings/internal/repository/db"
	cfg "github.com/JMURv/unona/ratings/pkg/config"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const configName = "dev.config"

// TODO: покрыть сервис тестами, выяснить степень покрытия

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic occurred: %v", err)
			os.Exit(1)
		}
	}()

	// Load configuration
	conf, err := cfg.LoadConfig(configName)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	port := conf.Port
	serviceName := conf.ServiceName

	// Start metrics and tracing
	metric := metrics.New()
	trace := tracing.New(serviceName, &conf.Jaeger)

	// Setting up main app
	broker := kafka.New(&conf.Kafka)
	cache := redis.New(&conf.Redis)
	repo := db.New(&conf.DB)

	svc := ctrl.New(repo, cache, broker)
	h := handler.New(svc)

	srv := metric.ConfigureServerGRPC() // grpc.NewServer()
	pb.RegisterRatingServiceServer(srv, h)
	reflection.Register(srv)

	// Start http server for prometheus
	go metric.Start(conf.Port + 1)

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c
		log.Println("Shutting down gracefully...")

		broker.Close()
		cache.Close()
		metric.Close()
		if err = trace.Close(); err != nil {
			log.Printf("Error closing tracer: %v", err)
		}
		srv.GracefulStop()
		os.Exit(0)
	}()

	// Start main server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("%v service is listening", serviceName)
	srv.Serve(lis)
}
