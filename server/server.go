package server

import (
    "log"
    "net"

    pb "github.com/msquare-2/grpc-user-service/proto"
    "google.golang.org/grpc"
)

func StartGRPCServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, NewUserServiceServer())

    log.Printf("server listening at %v", lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
