package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shatwik7/polycrate/libs/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAssetServiceServer
}

func (s *server) GetAsset(ctx context.Context, req *pb.GetAssetRequest) (*pb.GetAssetResponse, error) {
	log.Printf("Received GetAsset request for ID: %s", req.GetId())

	// Dummy response
	res := &pb.GetAssetResponse{
		Id:   req.GetId(),
		Name: "Example Asset",
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAssetServiceServer(grpcServer, &server{})

	log.Println("AssetService gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
