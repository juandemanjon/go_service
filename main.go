package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/juandemanjon/go_service/chatservice"
	"github.com/juandemanjon/go_service/gen"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	s := grpc.NewServer()

	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(s, healthcheck)
	reflection.Register(s)

	chatservice := chatservice.NewServer()
	gen.RegisterChatServiceServer(s, chatservice)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed: %v", err)
	}

}
