package server

import (
    "context"
    pb "github.com/msquare-2/grpc-user-service/proto"
    "errors"
)

type UserServiceServer struct {
    pb.UnimplementedUserServiceServer
    users []*pb.User
}

func NewUserServiceServer() *UserServiceServer {
    return &UserServiceServer{
        users: []*pb.User{
            {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
            {Id: 2, Fname: "John", City: "NY", Phone: 2345678901, Height: 5.9, Married: false},
            // Add more users here
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
        if (req.City != "" || user.City == req.City) &&
           (req.Phone == 0 || user.Phone == req.Phone) &&
           (user.Married == req.Married) {
            foundUsers = append(foundUsers, user)
        }
    }
    return &pb.UsersResponse{Users: foundUsers}, nil
}
