package main

import (
	"context"
	"fmt"
	pb "github.com/yuro8/todolist/gen/pb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	service := pb.NewTodoServiceClient(cc)
	resp, err := service.CreateTodoItem(context.Background(), &pb.TodoItem{
		Title:       "My new item",
		Description: "My super description",
	})

	fmt.Printf("Response: %v, error: %v", resp, err)
}
