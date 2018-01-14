package main

import (
	pb "github.com/What-If-I/gowsomebrowser/proto"
	"github.com/What-If-I/gowsomebrowser/proto/layout"
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

	// Constructing layout
	black := layout.Color{Value: "FFFFFF"}
	text := layout.Text{Content: "Test Content", Size: "10", Color: &black}
	log.Println(text)

	px40 := layout.Units{Value: 40, Type: layout.Units_PIXEL}
	percent100 := layout.Units{Value: 100, Type:layout.Units_PERCENT}
	area := layout.Area{Width: &px40, Height: &px40}

	textArea := layout.TextBox{Size: &area, Text: &text, Color: &black}
	elements := []*layout.Element{{&layout.Element_Textbox{Textbox: &textArea}}}
	elementsLayout := []*layout.ElementLayout{
		{ElemPosition: 1, RowStart: 2, RowEnd: 2, ColStart: 2, ColEnd: 2},
	}
	grid := layout.Grid{
		Size: &layout.Area{Width:&percent100, Height:&percent100}, Color: &black,
		Rows: 3, Columns: 3, Elements: elements, Elementslayout: elementsLayout}

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
