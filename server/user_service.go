package server

import (
	"context"
	"errors"

	pb "github.com/msquare-2/grpc-user-service/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		users: []*pb.User{
			{Id: 1, Fname: "Mustafa", City: "Indore", Phone: 1234568890, Height: 5.8, Married: false},
			{Id: 2, Fname: "John", City: "Mumbai", Phone: 2345678903, Height: 5.9, Married: true},
			{Id: 3, Fname: "Alice", City: "Gurgoan", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 4, Fname: "Rahul", City: "Delhi", Phone: 2345678401, Height: 5.9, Married: false},
			{Id: 5, Fname: "Jatin", City: "Ujjain", Phone: 1234967890, Height: 5.8, Married: true},
			{Id: 6, Fname: "Rajesh", City: "Bangalore", Phone: 2345278901, Height: 5.9, Married: false},
			{Id: 7, Fname: "Namrata", City: "Pune", Phone: 1234564890, Height: 5.8, Married: true},
			{Id: 8, Fname: "Sonu", City: "Chiplun", Phone: 2345618901, Height: 5.9, Married: false},
		},
	}
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.UserResponse{User: user}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.UsersResponse, error) {
	var foundUsers []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				foundUsers = append(foundUsers, user)
			}
		}
	}
	return &pb.UsersResponse{Users: foundUsers}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.UsersResponse, error) {
	var foundUsers []*pb.User
	for _, user := range s.users {
		if (req.City == "" || user.City == req.City) &&
			(req.Phone == 0 || user.Phone == req.Phone) &&
			(!req.Married || user.Married == req.Married) {
			foundUsers = append(foundUsers, user)
		}
	}
	return &pb.UsersResponse{Users: foundUsers}, nil
}
