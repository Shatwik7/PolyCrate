package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/shatwik7/polycrate/libs/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up gRPC client connection
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewAssetServiceClient(conn)

	// Initialize Fiber app
	app := fiber.New()

	// Define REST endpoint
	app.Get("/assets/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		grpcRes, err := client.GetAsset(ctx, &pb.GetAssetRequest{Id: id})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"id":   grpcRes.GetId(),
			"name": grpcRes.GetName(),
		})
	})

	log.Println("Gateway REST API listening on :8080")
	log.Fatal(app.Listen(":8080"))
}
