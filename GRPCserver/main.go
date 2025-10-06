package main

import (
	"awesomeProject/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedEmployeeServiceServer
	employees []*pb.Employee
}

func (s *server) GetEmployeeByID(ctx context.Context, req *pb.GetEmployeeByIDRequest) (*pb.EmployeeResponse, error) {
	for _, employee := range s.employees {
		if employee.Id == req.Id {
			return &pb.EmployeeResponse{Employee: employee}, nil
		}
	}
	return nil, fmt.Errorf("employee with id %v not found", req.Id)
}
func (s *server) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.EmployeeResponse, error) {
	for _, employee := range s.employees {
		if employee.Id == req.Id {
			employee.Name = req.Name
			employee.Email = req.Email
			employee.Age = req.Age
			return &pb.EmployeeResponse{Employee: employee}, nil
		}
	}
	return nil, fmt.Errorf("employee with id %v not found", req.Id)
}
func (s *server) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.EmployeeResponse, error) {
	newID := int32(len(s.employees) + 1)
	newEmployee := &pb.Employee{
		Id:    newID,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}
	s.employees = append(s.employees, newEmployee)

	return &pb.EmployeeResponse{Employee: newEmployee}, nil
}

func (s *server) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.EmployeeResponse, error) {
	for i, employee := range s.employees {
		if employee.Id == req.Id {
			// Удаляем сотрудника из среза
			s.employees = append(s.employees[:i], s.employees[i+1:]...)
			return &pb.EmployeeResponse{Employee: employee}, nil
		}
	}
	return nil, fmt.Errorf("employee with id %v not found", req.Id)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(grpcServer, &server{})

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
