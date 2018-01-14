package engine

import (
	"log"
	"net"

	"fmt"
	"github.com/What-If-I/gowsomebrowser/helpers"
	pb "github.com/What-If-I/gowsomebrowser/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var registeredApps = map[string]*pb.AppInfo{}

func RunGRPC(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAppServiceServer(server, &appService{})
	pb.RegisterViewServiceServer(server, &viewService{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type appService struct{}

func (eng *appService) SendLayout(ctx context.Context, in *pb.LayoutMessage) (*pb.Message, error) {
	fmt.Println("Got layout", in)
	return &pb.Message{Content: "Got layout!"}, nil
}

func (eng *appService) Register(ctx context.Context, in *pb.Message) (*pb.AppInfo, error) {
	uuid := helpers.GenUUID4()
	app := &pb.AppInfo{Id: uuid}
	registeredApps[uuid] = app

	return app, nil
}

type viewService struct{}

func (service *viewService) RunApp(link *pb.Link, stream pb.ViewService_RunAppServer) error {
	return nil
}
