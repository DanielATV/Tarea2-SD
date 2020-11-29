package chat

import (
	"log"
	"io"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"bufio"
	
	"io/ioutil"
	//"math"
	"os"
	"strconv"
)

type Server struct {
	Log map [string]string
	Libros []string
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
func (s *Server) SendPropuesta(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	//logica centralizada
	//guardar en el log (otro metodo)
	//logica distribuida
	return &Message{Body: "Propuesta OK"}, nil
}

func (s *Server) SendChunk(stream ChatService_SendChunkServer) (err error) {


	var chunkList [][]byte
	var libro string
	var buffer *Chunk
	flag:= 0
	for {
		buffer, err = stream.Recv()
		if err == io.EOF {
			log.Printf("LLego el libro %s", s.Libros[0])
			break
		}
		if err != nil {
			return err
		}

		if flag ==0 {
			libro = buffer.Nombre
			s.Libros = append(s.Libros,libro)
			flag = 1
			chunkList = append(chunkList,buffer.Chunk)
		}

		chunkList = append(chunkList,buffer.Chunk)
		
	}

	//Envia propuesta

	//logica centralizada
	var conn *grpc.ClientConn
	conn, error := grpc.Dial(":9001", grpc.WithInsecure())
	if error != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := NewChatServiceClient(conn)
	response, err1 := c.SendPropuesta(context.Background(), &Message{Body: "Aqui la propuesta"})
	if err1 != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)


	//Distribuir los chunks
	for i := 0; i < len(chunkList); i++ {

	
		// write to disk
		fileName := "bigfile_" + strconv.FormatUint(uint64(i), 10)
		_, err := os.Create(fileName)

		if err != nil {
				fmt.Println(err)
				os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, chunkList[i], os.ModeAppend)

		fmt.Println("Split to : ", fileName)
	}

	


	return stream.SendAndClose(&Message{Body: "Termino transferencia"})
	
	


}

func (s *Server) LibrosDis(ctx context.Context, in *Message) (*Message, error) {
	
	var actual string
	flag:= 0

	for index, value := range s.Libros {
		fmt.Println(index, value)

		if flag == 0{
			actual = value
			flag= 1
		} else{
			actual = actual + "%%%" + value

		}

	}
	return &Message{Body: actual}, nil
}

func (s *Server) RequestChunk(ctx context.Context, in *Message) (*Chunk, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	//read a chunk
	currentChunkFileName := in.Body
	newFileChunk, err := os.Open(currentChunkFileName)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	defer newFileChunk.Close()

	chunkInfo, err4 := newFileChunk.Stat()

	if err4 != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// calculate the bytes size of each chunk
	// we are not going to rely on previous data and constant

	var chunkSize int64 = chunkInfo.Size()
	chunkBufferBytes := make([]byte, chunkSize)

	// read into chunkBufferBytes
	reader := bufio.NewReader(newFileChunk)
	_, err = reader.Read(chunkBufferBytes)

	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}

	return &Chunk{Chunk: chunkBufferBytes}, nil
}
