package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/lafskelton/opensanic/pkg/server/proto/sanicdb"
	"google.golang.org/grpc"
)

type server struct {
	pb.SanicDBServer
	sanic *SanicDB
}

type lexiReplyArray struct {
	mu     sync.Mutex
	docs   []string
	keys   []string
	values []interface{}
}
type numReplyArray struct {
	mu     sync.Mutex
	docs   []string
	keys   []uint32
	values []interface{}
}

//#### gRPC service methods
//GET
func (s *server) GET(ctx context.Context, req *pb.GETRequest) (*pb.GETReply, error) {
	//Not implemented in open sanic
	fmt.Println("gRPc implemented but not functional in open sanic")
	//See github.com/lafskelton/sanic-additions/client for gRPC code
	return nil, nil
}

//SET
func (s *server) SET(ctx context.Context, req *pb.SETRequest) (*pb.SETReply, error) {
	fmt.Println("gRPc implemented but not functional in open sanic")
	//See github.com/lafskelton/sanic-additions/client for gRPC code
	return nil, nil
}

//Network starts the servers network service based on gRPC
func (s *SanicDB) Network() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSanicDBServer(grpcServer, &server{sanic: s})
	grpcServer.Serve(lis)
}
