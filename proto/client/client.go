package main

import (
	"context"
	"log"
	pb "project-employee/proto/model"

	"google.golang.org/grpc"
)

var host = "localhost:8888"
var client pb.EmployeeServiceClient
var ctx = context.Background()

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Not Connected => ", err)
	}
	defer conn.Close()
	client = pb.NewEmployeeServiceClient(conn)
	// getEmployees()
}

// func getEmployees() {
// 	resp, err := client.GetEmployees(ctx, &pb.Empty{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDU2MTA2OTEsInVzZXJuYW1lIjoiYWRtaW4yIn0.AYjwBx_D9BF-UXfWvFbIR-e9ZTNJneNtwreXhceLo1s"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Println(resp)
// }
