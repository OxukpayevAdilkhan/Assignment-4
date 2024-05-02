package awesomeProject

import (
	"context"
	"log"

	pb "C:\\Users\\admin\\GolandProjects\\awesomeProject\\user.proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// AddUser
	addUserResp, err := client.AddUser(context.Background(), &pb.UserRequest{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		log.Fatalf("AddUser failed: %v", err)
	}
	log.Printf("AddUser response: %v", addUserResp)

	// GetUser
	getUserResp, err := client.GetUser(context.Background(), &pb.UserRequest{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	log.Printf("GetUser response: %v", getUserResp)

	// ListUsers
	listUsersResp, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("ListUsers failed: %v", err)
	}
	log.Printf("ListUsers response: %v", listUsersResp)
}
