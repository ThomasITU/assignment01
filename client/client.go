package main

import (
	"assignment01/timePack"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	tC := timePack.NewTimeServiceClient(conn)

	// Contact the server and print out its response.
	ctx := context.Background()

	for {
		time.Sleep(time.Second * 1)
		fmt.Println(tC.GetTime(ctx, &timePack.GetTimeRequest{}))
	}
}
