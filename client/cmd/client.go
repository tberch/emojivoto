package main

import (
	"context"
	"os"
	"time"

	pb "github.com/buoyantio/emojivoto/emojivoto-web/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	client := pb.NewVotingServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = client.VoteDoughnut(ctx, &pb.VoteRequest{})
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	print("Success")
}
