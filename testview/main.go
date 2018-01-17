package main

import (
	pb "github.com/What-If-I/gowsomebrowser/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const address = "127.0.0.1:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("GRPC connection error:", err)
	}
	defer conn.Close()

	client := pb.NewViewServiceClient(conn)

	stream, err := client.RunApp(context.Background(), &pb.Link{"somelink"})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := stream.Recv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
