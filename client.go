package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
         
    "math"
	"os"
	"strings"
	"math/rand"

	"time"

	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func NumeroPuerto()string{
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    posibles := []string{"dist01:9000","dist02:9000","dist03:9000"}
    puerto := posibles[r1.Intn(3)]
    return puerto
}


func main() {

	var clientType int

	//Solicitar tipo de cliente
	fmt.Println("Indique el tipo de cliente")
	fmt.Println("0: Uploader")
	fmt.Println("1: Downloader")
	
	fmt.Scanln(&clientType)



	// ClientUploader
	if clientType == 0{

		//Conexion a DataNode
		randomChoice := NumeroPuerto()
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(randomChoice, grpc.WithInsecure(),grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),)
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()
		
		c := chat.NewChatServiceClient(conn)

		var chunkList [][]byte

		//Seleccion de libro

		fmt.Println("Selecione el libro que desea cargar")

		catalogo := [3]string{"Dracula","Frankestein","Peter Pan"}

		for ind, val := range catalogo {
			fmt.Println(ind, val)
	
		}
		var clientChoice int
		var choiceHolder string
		fmt.Scanln(&clientChoice)

		if clientChoice == 0{
			choiceHolder = "Dracula-Stoker_Bram.pdf"

		} else if clientChoice ==1{
			choiceHolder = "Frankenstein-Mary_Shelley.pdf"
		} else {
			choiceHolder = "Peter_Pan-J._M._Barrie.pdf"
		}

		//Particion del Libro
		fileToBeChunked := "./libros/" + choiceHolder // change here!

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

		//fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

		stream, _ := c.SendChunk(context.Background())

		log.Printf("Tiempo de inicio")


		// Envio de chunks
		for i := uint64(0); i < totalPartsNum; i++ {

				partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
				partBuffer := make([]byte, partSize)

				file.Read(partBuffer)

				chunkList = append(chunkList,partBuffer)

				stream.Send(&chat.Chunk{Chunk: partBuffer, Nombre: choiceHolder, Total: int64(totalPartsNum),
				Indice: int64(i)})

				log.Println("Mensaje Enviado")
		}

		response, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("Error when calling SendChunk: %s", err)
		}
		log.Printf("Respuesta del DataNode: %s", response.Body)


	// ClientDownloader
	} else{

	
		//Consulta Libros disponibles

		//Conexion al NameNode
		
		var conn *grpc.ClientConn
		conn, err := grpc.Dial("dist04:9000", grpc.WithInsecure(),grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),)
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		
		c := chat.NewChatServiceClient(conn)

		response, err := c.LibrosDis(context.Background(), &chat.Message{Body: "OK"})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Lista de libros %s", response.Body)

		var sep []string
		sep = strings.Split(response.Body,"%%%")

		var booksUploaded []string
		for index, value := range sep {
			fmt.Println(index, value)
			booksUploaded = append(booksUploaded,value)
	
		}

		//Elegir libro

		log.Printf("Indique el libro que desea")
		var downloadChoice int
		fmt.Scanln(&downloadChoice)



		//Consulta el log
		source, err := c.RequestLog(context.Background(), &chat.Message{Body: booksUploaded[downloadChoice]})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Ubicacion del archivo: %s", source.Body)

		holder := strings.Split(source.Body,"%%%")

		var bufferAux []string

		//Reconstruir Libro

		newFileName := booksUploaded[downloadChoice]
        _, err9 := os.Create(newFileName)

        if err9 != nil {
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


		var hostname string

		for _ , valor := range holder{

			fmt.Println(valor)

			bufferAux = strings.Split(valor,"&&&")

			//fmt.Printf("%T\n", bufferAux[1])

			if bufferAux[1] == "1"{
				hostname = "dist01:9000"

			} else if bufferAux[1] == "2"{
				hostname = "dist02:9000"

			} else {
				hostname = "dist03:9000"

			}

			//fmt.Println(hostname)



			//Descarga del datanode
			var conn *grpc.ClientConn
			conn, err6 := grpc.Dial(hostname, grpc.WithInsecure())
			if err6 != nil {
				log.Fatalf("did not connect: %s", err6)
			}
			defer conn.Close()

			cc := chat.NewChatServiceClient(conn)

			responseChunk, err := cc.RequestChunk(context.Background(), &chat.Message{Body: bufferAux[0]})
			if err != nil {
				log.Fatalf("Error when calling RequestChunk: %s", err)
			}

			log.Printf("Se recibio el chunk")
			
			chunkBufferBytes := responseChunk.Chunk

			// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
			// write/save buffer to disk
			//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

			n, err := file.Write(chunkBufferBytes)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			file.Sync() //flush to disk

			chunkBufferBytes = nil // reset or empty our buffer

			// free up the buffer for next cycle
			// should not be a problem if the chunk size is small, but
			// can be resource hogging if the chunk size is huge.
			// also a good practice to clean up your own plate after eating

			fmt.Println("Written ", n, " bytes")

			

		}

		// now, we close the newFileName
		file.Close()


		


		

		



	}

	

}