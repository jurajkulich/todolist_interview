package service

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"time"

	pb "github.com/yuro8/todolist/gen/pb"
)


type TodoListServer struct {
	firestoreClient *firestore.Client
}


func (server *TodoListServer) PrepareServer() error {
	ctx := context.Background()
	const saPath = "/home/juraj/Documents/Programming/backend/todolist/todolist-interview-firebase-adminsdk-en7bv-53a9847f1f.json"
	sa := option.WithCredentialsFile(saPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
		return err
	}
	server.firestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore app: %v", err)
		return err
	}
	return nil
}

func (server *TodoListServer) CreateTodoItem(ctx context.Context, todoItem *pb.TodoItem) (*pb.TodoItem, error) {
	currentTime := time.Now().Unix()
	ref, _, err := server.firestoreClient.Collection("items").Add(ctx, map[string]interface{}{
		"title": todoItem.Title,
		"description": todoItem.Description,
		"last_updated": currentTime,
		"done": pb.TodoItem_UNDONE,
	})
	if err != nil {
		return nil, err
	}
	todoitem := pb.TodoItem{Id: ref.ID, Title: todoItem.Title, Description: todoItem.Description,
		Done: pb.TodoItem_UNDONE, LastUpdated: &timestamp.Timestamp{ Seconds: currentTime}}
	return &todoitem, err
}

func (server *TodoListServer) UpdateTodoItem(ctx context.Context, todoItem *pb.TodoItem) (*pb.TodoItem, error) {
	mapped := map[string]interface{}{}
	if todoItem.Title != "" {
		mapped["title"] = todoItem.Title
	}
	if todoItem.Description != "" {
		mapped["description"] = todoItem.Description
	}
	if todoItem.Done != pb.TodoItem_UNDEFINED && todoItem.Done <= pb.TodoItem_UNDONE {
		mapped["done"]  = todoItem.Done
	}

	currentTime := time.Now().Unix()
	_, err := server.firestoreClient.Collection("items").Doc(todoItem.Id).Set(ctx, mapped, firestore.MergeAll)
	if err != nil {
		return nil, err
	}
	todoitem := pb.TodoItem{Id: todoItem.Id, Title: todoItem.Title, Description: todoItem.Description, LastUpdated: &timestamp.Timestamp{ Seconds: currentTime}}
	return &todoitem, err
}

func (server *TodoListServer) DeleteTodoItem(ctx context.Context, todoItem *pb.TodoItem) (*pb.TodoItem, error) {
	_, err := server.firestoreClient.Collection("items").Doc(todoItem.Id).Delete(ctx)
	if err != nil {
		return nil, err
	} else {
		return todoItem, err
	}
}

func (server *TodoListServer) GetTodoItem(ctx context.Context, todoItem *pb.TodoItem) (*pb.TodoItem, error) {
	dataSnap, err := server.firestoreClient.Collection("items").Doc(todoItem.Id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var item pb.TodoItem
	err = dataSnap.DataTo(&item)
	if err != nil {
		return nil, err
	}
	item.Id = dataSnap.Ref.ID
	return &item, nil
}

func (server *TodoListServer) ListTodoItems(ctx context.Context, empty *empty.Empty) (*pb.TodoItemList, error) {
	var items []*pb.TodoItem
	iter := server.firestoreClient.Collection("items").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var item pb.TodoItem
		err = doc.DataTo(&item)
		if err != nil {
			return nil, err
		}
		item.Id = doc.Ref.ID
		items = append(items, &item)
	}
	return &pb.TodoItemList{Todolist: items}, nil
}

func (server *TodoListServer) StopServer() error {
	return server.firestoreClient.Close()
}