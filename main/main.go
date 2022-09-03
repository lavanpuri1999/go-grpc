package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "src/proto"
)

type PersonManagementServer struct {
	pb.UnimplementedPersonServiceServer
}

func (p *PersonManagementServer) SayHello(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	log.Printf("Recieved message from client: %s\n", person)
	return person, nil
}

func main() {
	/* This is my first sample program. */
	fmt.Println("Hello, World!")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen to port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPersonServiceServer(grpcServer, &PersonManagementServer{})
	fmt.Printf("Server is listening at: %v\n\n", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000 %v", err)
	}
}
