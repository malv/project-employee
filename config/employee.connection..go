package config

import (
	"context"
	"log"

	pb "project-employee/proto/model"

	"google.golang.org/grpc"
)

var HostEmployee = "localhost:8888"
var CtxEmployee = context.Background()
var ClientEmployee pb.EmployeeServiceClient

func ConnectedEmployee() {
	conn, err := grpc.Dial(HostEmployee, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not connected err =>", err)
	}

	ClientEmployee = pb.NewEmployeeServiceClient(conn)
}
