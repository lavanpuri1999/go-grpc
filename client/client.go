package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "src/proto"
	"time"
)

const address = "localhost:9000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Wrond number: %v", err)
	}
	defer conn.Close()

	c := pb.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Alice"] = 43
	new_users["Bob"] = 30

	for name, age := range new_users {
		r, err := c.SayHello(ctx, &pb.Person{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Error bro : %v", err)
		}
		log.Printf("User Details: %s, %d\n", r.GetName(), r.GetAge())
	}

}
