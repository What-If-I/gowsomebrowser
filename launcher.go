//go:generate protoc -I ./proto/ --go_out=plugins=grpc:./proto/ ./proto/engine.proto
//go:generate protoc -I ./proto/ --go_out=plugins=grpc:./proto/ ./proto/layout/layout.proto

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
	//appName := "main.go"
	engine.RunGRPC(defaultPort)
	//go runView()

	//url := "http://127.0.0.1:8000/" + appName
	//downloadFromUrl(url)
	//runGoApp(appName)
}

func downloadByUrl(url string) {
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


func runGoApp(path string) {
	// todo: pass port as an arg to the app
	cmd := exec.Command("go", "run " + path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run %s: %s", path, err)
	}
	fmt.Printf(out.String())
}

// todo: that should run view app
func runView() {}
