package engine

import (
	"log"
	"net"

	"fmt"
	pb "github.com/What-If-I/gowsomebrowser/proto"
	"github.com/What-If-I/gowsomebrowser/proto/layout"
	"github.com/golang/protobuf/jsonpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strings"
)

func RunGRPC(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterEngineServiceServer(server, &engine{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type engine struct{}

func (eng *engine) SendLayout(ctx context.Context, in *browser_layout.Grid) (*pb.Message, error) {
	r := strings.NewReader("")
	jsonpb.Unmarshal(r, in)
	fmt.Println("Got layout", r.Len())
	return &pb.Message{"Got layout!"}, nil
}

func (eng *engine) Connect(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	return &pb.Message{"Connected!"}, nil
}
