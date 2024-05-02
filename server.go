package awesomeProject

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "C:\\Users\\admin\\GolandProjects\\awesomeProject\\user.proto"
	"google.golang.org/grpc"
)

type server struct {
	users []*pb.UserResponse
}

func (s *server) AddUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := &pb.UserResponse{
		Id:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}
	s.users = append(s.users, user)
	return user, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *server) ListUsers(ctx context.Context, req *pb.Empty) (*pb.UserListResponse, error) {
	return &pb.UserListResponse{Users: s.users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("Server started...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
