package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	pb "src/proto"
	"strings"
)

type PersonManagementServer struct {
	pb.UnimplementedPersonServiceServer
}

func (p *PersonManagementServer) SayHello(ctx context.Context, person *pb.Person) (*pb.Person, error) {
	log.Printf("Recieved message from client: %s\n", person)
	return person, nil
}

func (p *PersonManagementServer) QueryLogFiles(ctx context.Context, query *pb.QueryInput) (*pb.QueryResults, error) {
	results, count := scanLogs(query.GetQuery())
	return &pb.QueryResults{LogLines: results, Count: count}, nil
}

func scanLogs(query string) ([]string, int32) {
	file, err := os.Open("sample.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int32
	var results []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // GET the line string
		if strings.Contains(line, query) {
			fmt.Println(line)
			count += 1
			results = append(results, line)
		}
	}
	return results, count
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
