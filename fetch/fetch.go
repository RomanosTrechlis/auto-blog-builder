package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/RomanosTrechlis/auto-blog-builder/service"
	"github.com/RomanosTrechlis/blog-generator/cli"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement fetchService
type server struct{}

func (s *server) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchResponse, error) {
	cli.Download(in.DsType, in.DsRepository, in.DsDestFolder)
	return &pb.FetchResponse{Res: "Fetching completed"}, nil
}

func main() {
	fmt.Println("Fetch service started.")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := server{}
	pb.RegisterFetcherServer(s, &srv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
