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

// server is used to implement generateService
type server struct{}

func (s *server) Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	siteInfo := pb.MapRequestToSiteInformation(in)
	cli.Generate(&siteInfo)
	return &pb.GenerateResponse{Res: "Generated"}, nil
}

func main() {
	fmt.Println("Generate service started.")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := server{}
	pb.RegisterGeneratorServer(s, &srv)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
