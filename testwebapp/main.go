package main

import (
	pb "github.com/What-If-I/gowsomebrowser/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	layoutStub "github.com/What-If-I/gowsomebrowser/testwebapp/layout"
)

const address = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("GRPC connection error:", err)
	}
	defer conn.Close()

	grid := layoutStub.GetTestLayout()

	client := pb.NewAppServiceClient(conn)

	appName := pb.Message{Content: "TestApp"}
	appID, _ := client.Register(context.Background(), &appName)
	log.Println("got AppId", appID)

	layoutMessage := pb.LayoutMessage{AppInfo: appID, Grid: &grid}

	resp, err := client.SendLayout(context.Background(), &layoutMessage)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
