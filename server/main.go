package main

import (
	"log"
	"net"

	"github.com/Joao-Pedro-MB/grpc-server/vehicle"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()

	vehicle.RegisterVehicleServiceServer(grpcServer, &vehicle.VehicleServer{})

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8080: %v", err)
	}
}
