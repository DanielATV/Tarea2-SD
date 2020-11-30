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
	"strings"
	"os"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	Log map [string]string
	Libros []string
	Mode int
	Id string
	mux sync.Mutex
	State int
}

func escribir(libro string, partes int, lis []string ){
	//fmt.Println(libro,partes,lis)
	archi, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
        log.Fatal(err)
    }
    datawriter := bufio.NewWriter(archi)
	aux := strconv.Itoa(partes)
	aux = libro + " " + aux
	var aux2 string
	_,_ = datawriter.WriteString(aux+"\n")
	for ind, val := range lis {
		aux = strconv.Itoa(ind)
		//aux2 = strconv.Itoa(val)
		aux2 = val
		aux = libro+"_"+aux + " " +aux2
        _,_ = datawriter.WriteString(aux+"\n")
    }
    datawriter.Flush()
	archi.Close()
}


func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) WriteLog(ctx context.Context, in *LogInfo) (*Message, error) {
	log.Printf("Se desea escribir: %s", in.Log)
	prop := strings.Split(in.Log,"%%%")
	if s.Mode == 0{
		//Logica centralizada
		s.mux.Lock()
		flag:= 0

		
		for flag ==0 {

			if s.State == 0{
				flag = 1
				s.State = 1
				escribir(in.Nombre,int(in.Partes),prop)
			} else{
				time.Sleep(time.Duration(5) * time.Second)
			}
			

		}
		
		s.State = 0
		s.mux.Unlock()
		

	} else {
		//Logica distribuida
		escribir(in.Nombre,int(in.Partes),prop)

	}

	


	return &Message{Body: "Recurso Liberado"}, nil
}

func (s *Server) DistributeChunk(ctx context.Context, in *Chunk) (*Message, error) {
	log.Printf("Recibido el fragmentro del libro: %s", in.Nombre)
	// write to disk
	fileName := "./DB/"+in.Nombre+"_" + strconv.FormatUint(uint64(in.Indice), 10)
	_, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write/save buffer to disk
	ioutil.WriteFile(fileName, in.Chunk, os.ModeAppend)

	fmt.Println("Split to : ", fileName)

	return &Message{Body: "Chunk recibido"}, nil
}

func (s *Server) CheckStatus(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Chequeando estado del Nodo: %s", in.Body)
	return &Message{Body: "ACK"}, nil
}
func (s *Server) SendPropuesta(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Propuesta recibida: %s", in.Body)
	var prop string

	if s.Mode == 0{
		//logica centralizada

		prop= "1%%%1%%%1%%%1%%%1%%%1%%%1%%%1%%%1"

		//checkear propuesta

	} else{
		//logica distribuida

	}
	
	//guardar en el log (otro metodo)
	
	return &Message{Body: prop}, nil
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

	var distribution string
	
	if s.Mode == 0{
		//logica centralizada

		//conexion al NameNode
		var conn *grpc.ClientConn
		conn, error := grpc.Dial(":9001", grpc.WithInsecure())
		if error != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := NewChatServiceClient(conn)

		//crear propuesta
		response, err1 := c.SendPropuesta(context.Background(), &Message{Body: "1%%%2%%%3%%%1%%%2%%%3%%%1%%%2%%%3",
		Id: s.Id})
		if err1 != nil {
			log.Fatalf("Error when calling SayHello: %s", err1)
		}
		log.Printf("Propuesta Recibida %s", response.Body)

		distribution = response.Body

		//Escribir en el log
		response, err2 := c.WriteLog(context.Background(), &LogInfo{Log: distribution, Nombre: libro,
		Partes: int64(len(chunkList))})
		if err2 != nil {
			log.Fatalf("Error when calling SayHello: %s", err2)
		}
		log.Printf("Propuesta Recibida %s", response.Body)


	} else {
		//logica distribuida
	}
	
	

	//Distribuir los chunks

	var sep []string
	cont:= 0
	sep = strings.Split(distribution,"%%%")
	for i := 0; i < len(chunkList); i++ {

		if sep[cont] == s.Id{

			// write to disk
			fileName := "./DB/"+libro+"_" + strconv.FormatUint(uint64(i), 10)
			_, err := os.Create(fileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// write/save buffer to disk
			ioutil.WriteFile(fileName, chunkList[i], os.ModeAppend)

			fmt.Println("Split to : ", fileName)

		} else {
			if sep[cont] == "1"{

				//conexion al NameNode1
				var conn *grpc.ClientConn
				conn, error := grpc.Dial(":9000", grpc.WithInsecure())
				if error != nil {
					log.Fatalf("did not connect: %s", err)
				}
				conn.Close()

				c := NewChatServiceClient(conn)

				//Envio del Chunk
				filler, err := c.DistributeChunk(context.Background(), &Chunk{Chunk: chunkList[i],
					Nombre: libro,
					Indice: int64(i)})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}

				log.Printf(filler.Body)
				conn.Close()


			}

			if sep[cont] == "2"{
				//conexion al NameNode2
				var conn *grpc.ClientConn
				conn, error := grpc.Dial(":9002", grpc.WithInsecure())
				if error != nil {
					log.Fatalf("did not connect: %s", err)
				}
				conn.Close()

				c := NewChatServiceClient(conn)

				//Envio del Chunk
				filler, err := c.DistributeChunk(context.Background(), &Chunk{Chunk: chunkList[i],
					Nombre: libro,
					Indice: int64(i)})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}

				log.Printf(filler.Body)
				conn.Close()
				
			}

			if sep[cont] == "3"{
				//conexion al NameNode3
				var conn *grpc.ClientConn
				conn, error := grpc.Dial(":9003", grpc.WithInsecure())
				if error != nil {
					log.Fatalf("did not connect: %s", err)
				}
				conn.Close()

				c := NewChatServiceClient(conn)

				//Envio del Chunk
				filler, err := c.DistributeChunk(context.Background(), &Chunk{Chunk: chunkList[i],
					Nombre: libro,
					Indice: int64(i)})
				if err != nil {
					log.Fatalf("Error when calling SayHello: %s", err)
				}

				log.Printf(filler.Body)
				conn.Close()
				
			}
		}
			

			

		cont = cont + 1
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
