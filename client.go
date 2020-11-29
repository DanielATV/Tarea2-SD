package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
         
    "math"
	"os"
	//"strings"

	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func main() {

	var clientType int
	clientType = 0

	//Solicitar tipo de cliente
	fmt.Println("Indique el tipo de cliente")
	fmt.Println("0: Uploader")
	fmt.Println("1: Downloader")
	
	fmt.Scanln(&clientType)


	// ClientUploader
	if clientType == 0{

		//Conexion a DataNode
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := chat.NewChatServiceClient(conn)

		var chunkList [][]byte

		//Particion del Libro
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


		// Envio de chunks
		for i := uint64(0); i < totalPartsNum; i++ {

				partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
				partBuffer := make([]byte, partSize)

				file.Read(partBuffer)

				chunkList = append(chunkList,partBuffer)

				stream.Send(&chat.Chunk{Chunk: partBuffer, Nombre: "Dracula-Stoker_Bram.pdf"})
		}

		response, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)


	// ClientDownloader
	} else{

		


		//Consulta Libros disponibles

		//Conexion al NameNode
		/*
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9001", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()


		
		c := chat.NewChatServiceClient(conn)

		response, err := c.LibrosDis(context.Background(), &chat.Message{Body: "OK"})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", response.Body)

		var sep []string
		sep = strings.Split(response.Body,"%%%")
		for index, value := range sep {
			fmt.Println(index, value)
	
		}
		*/
		//Consulta el log

		books := [9]string{"bigfile_0", "bigfile_1", "bigfile_2","bigfile_3", "bigfile_4", "bigfile_5",
		"bigfile_6","bigfile_7","bigfile_8"}

		//Descarga del datanode
		//var conn *grpc.ClientConn
		conn, err6 := grpc.Dial(":9000", grpc.WithInsecure())
		if err6 != nil {
			log.Fatalf("did not connect: %s", err6)
		}
		defer conn.Close()

		cc := chat.NewChatServiceClient(conn)


		newFileName := "libro.pdf"
        _, err := os.Create(newFileName)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        //set the newFileName file to APPEND MODE!!
        // open files r and w

        file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

		for j := 0; j < len(books); j++ {

            
			responseChunk, err := cc.RequestChunk(context.Background(), &chat.Message{Body: books[j]})
			if err != nil {
				log.Fatalf("Error when calling SayHello: %s", err)
			}
			

			// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
			// write/save buffer to disk
			//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

			n, err := file.Write(responseChunk.Chunk)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			file.Sync() //flush to disk

			// free up the buffer for next cycle
			// should not be a problem if the chunk size is small, but
			// can be resource hogging if the chunk size is huge.
			// also a good practice to clean up your own plate after eating

			fmt.Println("Written ", n, " bytes")

			fmt.Println("Recombining part [", j, "] into : ", newFileName)
		}

		// now, we close the newFileName
		file.Close()



	}

	

}