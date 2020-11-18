package config

import (
	"context"
	"log"

	pb "project-employee/model"

	"google.golang.org/grpc"
)

var Host = "camskoleksi.com:8091"
var Ctx = context.Background()
var Client pb.UserServiceClient

func Connected() {
	conn, err := grpc.Dial(Host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	Client = pb.NewUserServiceClient(conn)
}
