package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/shatwik7/polycrate/lib/db"
	userpb "github.com/shatwik7/polycrate/lib/protos/user"
	service "github.com/shatwik7/polycrate/services/user_service"
	grpcserver "google.golang.org/grpc"
)

func main() {
	dataSourceName := "postgres://polycrate:polycreate@localhost:5432/polycrate_db?sslmode=disable"

	if dataSourceName == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Connect to database
	database, err := db.NewDB(dataSourceName)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Test DB connection
	if err := database.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Start TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server and register services
	grpcServer := grpcserver.NewServer()
	userService := service.NewUserServer(database)
	userpb.RegisterUserServiceServer(grpcServer, userService)

	// Run gRPC server in a goroutine
	go func() {
		log.Println("gRPC server running on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("Server shut down cleanly")
}
