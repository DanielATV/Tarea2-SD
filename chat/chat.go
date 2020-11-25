package chat

import (
	"log"
	"io"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) SendChunk(stream ChatService_SendChunkServer) (err error) {
	for {
		_, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				goto END
			}

			
			return nil
		}
	}

END:
        // once the transmission finished, send the
        // confirmation if nothign went wrong
	err = stream.SendAndClose(&Message{Body: "Termino transferencia"})
	

	return
}
