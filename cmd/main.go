package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/Negative3103/grpc/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}
}
