package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/bookish-goggles/protogen"
	"google.golang.org/grpc"
)

const (
	maxSendSize = 100 * 1024 * 1024
	maxRecvSize = 100 * 1024 * 1024
)

// server is used to implement KVStore Service
type server struct {
	pb.UnimplementedKVStoreServer
}

//----------------------------------------------------------------------------
//  KVStore Service API
//----------------------------------------------------------------------------

func (s *server) Get(ctx context.Context,
	in *pb.GetReq) (*pb.GetRes, error) {
	log.Printf("Received Get(key=%s)", in.Key)
	val, err := state.KVStoreInstance.Get(in.Key)
	return &pb.GetRes{Val: val, Err: &err}, nil
}

func (s *server) Put(ctx context.Context,
	in *pb.PutReq) (*pb.PutRes, error) {
	log.Printf("Received Put(key=%s, val=%s)", in.Key, in.Val)
	err := state.KVStoreInstance.Put(in.Key, in.Val)
	return &pb.PutRes{Err: &err}, nil
}

func StartServer() {
	state = new(ServerState)
	state.Init()

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxRecvSize), grpc.MaxSendMsgSize(maxSendSize))
	pb.RegisterKVStoreServer(s, &server{})

	fmt.Println("KVStore server ready to service requests...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	s.Stop()
}
