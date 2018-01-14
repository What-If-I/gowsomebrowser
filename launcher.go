//go:generate protoc -I ./proto/ --go_out=plugins=grpc:./proto/ ./proto/engine.proto
//go:generate protoc -I ./proto/ --go_out=plugins=grpc:./proto/layout/ ./proto/layout.proto

package main

import (
	"bytes"
	"fmt"
	"github.com/What-If-I/gowsomebrowser/engine"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const defaultPort = ":50051"

func main() {
	go engine.RunGRPC(defaultPort)

	url := "http://127.0.0.1:8000/main.go"
	downloadFromUrl(url)
	cmd := exec.Command("go", "run", "main.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String())
}

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	log.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Error while creating", fileName, "-", err)
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error while downloading", url, "-", err)
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Fatalln("Error while downloading", url, "-", err)
	}

	log.Println(n, "bytes downloaded.")
}
