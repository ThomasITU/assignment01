package main

import (
	"assignment01/timePack"
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type Server struct {
	timePack.UnimplementedTimeServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	timePack.RegisterTimeServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s  %v", port, err)
	}

}

func (s *Server) GetTime(ctx context.Context, message *timePack.GetTimeRequest) (*timePack.Response, error) {
	return &timePack.Response{Name: time.Now().String()}, nil
}
