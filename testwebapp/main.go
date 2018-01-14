package main

import (
	pb "github.com/What-If-I/gowsomebrowser/proto"
	layout "github.com/What-If-I/gowsomebrowser/proto/layout"
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
	log.Println(pb.Message{"Hey"})

	black := layout.Color{Value: "FFFFFF"}
	text := layout.Text{Content: "Test Content", Size: "10", Color: &black}
	log.Println(text)

	px10 := layout.Units{Value: 10}
	area := layout.Area{
		Width: &px10, Height: &px10, MarginLeft: &layout.Units{}, MarginTop: &layout.Units{}, MarginRight: &layout.Units{},
		MarginBottom: &layout.Units{}, PaddingLeft: &layout.Units{}, PaddingTop: &layout.Units{},
		PaddingRight: &layout.Units{}, PaddingBottom: &layout.Units{}}
	textArea := layout.TextBox{Size: &area, Text: &text, Color: &black}
	elements := []*layout.Element{{&layout.Element_Textbox{Textbox: &textArea}}}
	grid := layout.Grid{
		Size: &area, Color: &black,
		Rows: []*layout.Grid{}, Columns: []*layout.Grid{}, Elements: elements}

	client := pb.NewEngineServiceClient(conn)
	resp, err := client.SendLayout(context.Background(), &grid)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
