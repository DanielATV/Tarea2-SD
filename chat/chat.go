package chat

import (
	"log"
	"io"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
	return &Message{Body: "Hello From the Server!"}, nil
}
func (s *Server) SendChunk(stream ChatService_SendChunkServer) (err error) {

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
		}
		
	}

	//Envia propuesta
	var conn *grpc.ClientConn
	conn, error := grpc.Dial(":9000", grpc.WithInsecure())
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

