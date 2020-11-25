package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
         
    "math"
    "os"

	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	var chunkList [][]byte


	fileToBeChunked := "./libros/Dracula-Stoker_Bram.pdf" // change here!

	file, err := os.Open(fileToBeChunked)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000 // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	stream, _ := c.SendChunk(context.Background())

	for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			file.Read(partBuffer)

			chunkList = append(chunkList,partBuffer)

			stream.Send(&chat.Chunk{Chunk: partBuffer})
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}