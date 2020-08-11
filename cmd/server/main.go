package main

import (
	"github.com/yuro8/todolist/service"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/yuro8/todolist/gen/pb"

)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var server service.TodoListServer
	err = server.PrepareServer()
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &server)

	defer func(s *service.TodoListServer) {
		err = s.StopServer()
		if err != nil {
			log.Printf("failed to stop server: %v", err)
		}
	}(&server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
