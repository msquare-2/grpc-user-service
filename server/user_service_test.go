package server

import (
    "context"
    "testing"
    pb "github.com/msquare-2/grpc-user-service/proto"
)

func TestGetUser(t *testing.T) {
    s := NewUserServiceServer()

    req := &pb.GetUserRequest{Id: 1}
    res, err := s.GetUser(context.Background(), req)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if res.User.Id != 1 {
        t.Errorf("expected user ID to be 1, got %d", res.User.Id)
    }
}

func TestGetUsers(t *testing.T) {
    s := NewUserServiceServer()

    req := &pb.GetUsersRequest{Ids: []int32{1, 2}}
    res, err := s.GetUsers(context.Background(), req)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(res.Users) != 2 {
        t.Errorf("expected 2 users, got %d", len(res.Users))
    }
}

func TestSearchUsers(t *testing.T) {
    s := NewUserServiceServer()

    req := &pb.SearchUsersRequest{City: "LA"}
    res, err := s.SearchUsers(context.Background(), req)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(res.Users) != 1 {
        t.Errorf("expected 1 user, got %d", len(res.Users))
    }
}
