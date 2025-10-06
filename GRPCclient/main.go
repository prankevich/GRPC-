package main

import (
	"GRPC-client/pbf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pbf.NewEmployeeServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	respons, err := client.CreateEmployee(ctx, &pbf.CreateEmployeeRequest{
		Age:   12,
		Name:  "Alex",
		Email: "alex@gmail.com",
	})
	if err != nil {
		log.Fatalf("could not create employee: %v", err)
	}
	fmt.Println("Create", respons)

	respons, err = client.UpdateEmployee(ctx, &pbf.UpdateEmployeeRequest{
		Id:    2,
		Age:   15,
		Name:  "Alex",
		Email: "alex@gmail.com",
	})
	if err != nil {
		log.Fatalf("could not create employee: %v", err)
	}
	fmt.Println("Update", respons)

	respons, err = client.GetEmployeeByID(ctx, &pbf.GetEmployeeByIDRequest{
		Id: 2,
	})
	if err != nil {
		log.Fatalf("could not create employee: %v", err)
	}
	fmt.Println("Get by ID", respons)

}
