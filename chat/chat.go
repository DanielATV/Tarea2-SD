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
	"math/rand"
)

type Server struct {
	Log map [string]string
	Libros []string
	Mode int
	Id string
	mux sync.Mutex
	State int
}
type mensaje struct{
	nombre string
	largo_chunks int
}

func createRegPropDist (ej1 mensaje, m1 bool, m2 bool, m3 bool) string{ //Funcion que crea una propuesta dependiendo de sus restricciones
	var prop_c string
	var posibles []string

	if (m1 && m2 && m3){
		posibles = []string {"1%%%","2%%%","3%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m1 && m2){
		posibles = []string {"1%%%","2%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m1 && m3){
		posibles = []string{"1%%%","3%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m2 && m3){
		posibles = []string{"2%%%","3%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m1){
		posibles = []string{"1%%%"}
		prop_c = createProp(ej1,posibles)
	}else if(m2){
		posibles = []string{"2%%%"}
		prop_c = createProp(ej1,posibles)
	}else if(m3){
		posibles = []string{"3%%%"}
		prop_c = createProp(ej1,posibles)
	}else if(m1 && m2 && m3){
		posibles = []string {"1%%%","2%%%","3%%%"}
	}
	return prop_c

}

func createProp (ej1 mensaje,posibles []string) string{ //Funcion que crea una propuesta
	var prop_c string
	var prob int
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    var cont int = 0

	if ej1.largo_chunks>=3{


		prop_c = prop_c+posibles[0]
		cont = cont + 1 

		if len(posibles)>1{
			prop_c = prop_c+posibles[1]
			cont = cont +1
		}
		if len(posibles)>2{
			prop_c = prop_c+posibles[2]
			cont = cont + 1
		}
		for i:=cont + 1; i<= ej1.largo_chunks; i++{
			prob = r1.Intn(len(posibles))
			prop_c = prop_c+posibles[prob]
		}
	} else if ej1.largo_chunks == 2{

		prop_c = prop_c+posibles[0]

		if ( len(posibles)>1){
			prop_c = prop_c+posibles[1]
		}else{
			prop_c = prop_c+posibles[0]
		}
		
	}else{
		prop_c = prop_c+posibles[0]
	}

	prop_c = prop_c[0:len(prop_c)-3]
	return prop_c
}

func createRegProp (ej1 mensaje, m1 bool, m2 bool, m3 bool) string{ //Funcion que crea una propuesta dependiendo de sus restricciones
	var prop_c string
	var posibles []string

	if (m1 && m2){
		posibles = []string {"1%%%","2%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m1 && m3){
		posibles = []string{"1%%%","3%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m2 && m3){
		posibles = []string{"2%%%","3%%%"}
		prop_c = createProp(ej1,posibles)
	} else if(m1){
		posibles = []string{"1%%%"}
		prop_c = createProp(ej1,posibles)
	}else if(m2){
		posibles = []string{"2%%%"}
		prop_c = createProp(ej1,posibles)
	}else if(m3){
		posibles = []string{"3%%%"}
		prop_c = createProp(ej1,posibles)
	}
	return prop_c

}

func escribir(libro string, partes int, lis []string ) string{
	//fmt.Println(libro,partes,lis)

	var logRecord string
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
		logRecord = logRecord + libro + "_" + aux+ "&&&"+aux2 + "%%%"
		aux = libro+"_"+aux + " " +aux2

        _,_ = datawriter.WriteString(aux+"\n")
    }
    datawriter.Flush()
	archi.Close()

	logRecord = logRecord[0:len(logRecord)-3]

	return logRecord
}


func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) RequestLog(ctx context.Context, in *Message) (*Message, error) {
	return &Message{Body:s.Log[in.Body]}, nil
}

func (s *Server) WriteLog(ctx context.Context, in *LogInfo) (*Message, error) {
	log.Printf("Se desea escribir: %s", in.Log)

	var logRecord string

	prop := strings.Split(in.Log,"%%%")

	if s.Mode == 0{
		//Logica centralizada
		s.mux.Lock()
		flag:= 0

		
		for flag ==0 {

			if s.State == 0{
				flag = 1
				s.State = 1
				logRecord = escribir(in.Nombre,int(in.Partes),prop)
			} else{
				time.Sleep(time.Duration(5) * time.Second)
			}
			

		}
		
		s.State = 0
		s.mux.Unlock()
		

	} else {
		//Logica distribuida
		logRecord = escribir(in.Nombre,int(in.Partes),prop)

	}

	s.Libros = append(s.Libros,in.Nombre)
	s.Log[in.Nombre] = logRecord
	


	return &Message{Body: "Log actualizado"}, nil
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
	var prop_c string

	if s.Mode == 0{
		//logica centralizada
		

		m1 := true
		m2 := true
		m3 := true

		var sep []string
		sep = strings.Split(in.Body,"%%%")

		//checkear conexiones

	
		//var conn *grpc.ClientConn
		//var conn2 *grpc.ClientConn
		//var conn3 *grpc.ClientConn
		//DataNode 1
		
		_, err := grpc.Dial(":9000", grpc.WithInsecure(),grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),)
		if err != nil {
			fmt.Println("Fallo conexion al DataNode 1")
			m1  = false
		}
		//defer conn.Close()
		
		//DataNode 2
		_, err2 := grpc.Dial(":9002", grpc.WithInsecure(),grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),)
		if err2 != nil {
			fmt.Println("Fallo conexion al DataNode 2")
			m2  = false
		}
	
		//defer conn2.Close()

		//DataNode 3
		_, err3 := grpc.Dial(":9003", grpc.WithInsecure(),grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),)
		if err3 != nil {
			fmt.Println("Fallo conexion al DataNode 3")
			m3  = false
		}
		
		//defer conn3.Close()


		ej1 := mensaje{nombre : "N/A", largo_chunks :len(sep) }

		fmt.Println(m1,m2,m3)

		//checkear propuesta
		if (m1 && m2 && m3){
			// seguir propuesta enviada por el data node
			prop_c = in.Body
		}else{
			prop_c = createRegProp(ej1,m1,m2,m3)
		}

	} else{
		//logica distribuida
		prop_c = in.Body

	}
	
	//guardar en el log (otro metodo)
	
	return &Message{Body: prop_c}, nil
}

func (s *Server) SendChunk(stream ChatService_SendChunkServer) (err error) {


	var chunkList [][]byte
	var libro string
	var buffer *Chunk
	var cantidadMensajes int
	flag:= 0
	for {
		//fmt.Println(len(chunkList))
		buffer, err = stream.Recv()

		//fmt.Println("Llego chunk")
		if err == io.EOF {
			log.Printf("LLego el libro %s", libro)
			break
		}
		if err != nil {
			return err
		}

		if flag ==0 {
			libro = buffer.Nombre
			//s.Libros = append(s.Libros,libro)
			flag = 1
			chunkList = append(chunkList,buffer.Chunk)
			cantidadMensajes = int(buffer.Total)
			continue
		}

	
		chunkList = append(chunkList,buffer.Chunk)
		//fmt.Println(buffer.Indice)
		buffer.Chunk = nil
		
	}

	//fmt.Println(len(chunkList))

	fmt.Println("Envio propuesta")

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
		var prop_c string
		var prob int
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		
		var aux string

		if len(chunkList)>=3{
			prop_c = prop_c + "1%%%"
			prop_c = prop_c + "2%%%"
			prop_c = prop_c + "3%%%"
			for i:=4; i<= len(chunkList); i++{
				prob = r1.Intn(3) +1
				aux = strconv.Itoa(prob)
				prop_c = prop_c + aux + "%%%"
			}
		} else if len(chunkList) == 2{
			prop_c = prop_c + "1%%%"
			prop_c = prop_c + "2%%%"
		}else{
			prop_c = prop_c + "1%%%"
		}
	
		prop_c = prop_c[0:len(prop_c)-3]


		response, err1 := c.SendPropuesta(context.Background(), &Message{Body: prop_c,
		Id: s.Id})
		if err1 != nil {
			fmt.Println("Murio antes del log")
			log.Fatalf("Error when calling SayHello: %s", err1)
		}
		log.Printf("Propuesta Recibida %s", response.Body)

		distribution = response.Body

		//Escribir en el log
		response, err2 := c.WriteLog(context.Background(), &LogInfo{Log: distribution, Nombre: libro,
		Partes: int64(cantidadMensajes)})
		if err2 != nil {
			fmt.Println("Murio en el log")
			log.Fatalf("Error when calling SayHello: %s", err2)
		}
		log.Printf("Propuesta Recibida %s", response.Body)


	} else {
		//logica distribuida


		m1 := true
		m2 := true
		m3 := true
		ej1 := mensaje{nombre : "N/A", largo_chunks :len(chunkList) }


		//crear propuesta
		var prop_c string
		var prob int
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		
		var aux string

		if len(chunkList)>=3{
			prop_c = prop_c + "1%%%"
			prop_c = prop_c + "2%%%"
			prop_c = prop_c + "3%%%"
			for i:=4; i<= len(chunkList); i++{
				prob = r1.Intn(3) +1
				aux = strconv.Itoa(prob)
				prop_c = prop_c + aux + "%%%"
			}
		} else if len(chunkList) == 2{
			prop_c = prop_c + "1%%%"
			prop_c = prop_c + "2%%%"
		}else{
			prop_c = prop_c + "1%%%"
		}
	
		prop_c = prop_c[0:len(prop_c)-3]

		if s.Id == "1"{


		
			//DataNode 2
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(":9002", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}

			

			c := NewChatServiceClient(conn)

			_, err1 := c.SendPropuesta(context.Background(), &Message{Body: prop_c})
			if err1 != nil {
				m2 = false
			}

			conn.Close()

			//DataNode 3
			conn3, err3 := grpc.Dial(":9003", grpc.WithInsecure())
			if err3 != nil {
				log.Fatalf("did not connect: %s", err)
			}

			c = NewChatServiceClient(conn)
			_, err33 := c.SendPropuesta(context.Background(), &Message{Body: prop_c})
			if err33 != nil {
				m3 = false
			}
			conn3.Close()


		}

		if s.Id == "2"{
		//DataNode 1
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}

		c := NewChatServiceClient(conn)

		_, err1 := c.SendPropuesta(context.Background(), &Message{Body: prop_c})
		if err1 != nil {
			m1 = false
		}
		
		conn.Close()
		
		//DataNode 3
		conn3, err3 := grpc.Dial(":9003", grpc.WithInsecure())
		if err3 != nil {
			log.Fatalf("did not connect: %s", err)
		}

		c = NewChatServiceClient(conn)
		_, err33 := c.SendPropuesta(context.Background(), &Message{Body: prop_c})
		if err33 != nil {
			m3 = false
		}
		conn3.Close()

		} else {
		//DataNode 1
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}

		c := NewChatServiceClient(conn)

		_, err1 := c.SendPropuesta(context.Background(), &Message{Body: prop_c})
		if err1 != nil {
			m1 = false
		}
		
		conn.Close()
		
		//DataNode 2
		conn2, err2 := grpc.Dial(":9002", grpc.WithInsecure())
		if err2 != nil {
			log.Fatalf("did not connect: %s", err)
		}

		c = NewChatServiceClient(conn)
		_, err22 := c.SendPropuesta(context.Background(), &Message{Body: prop_c})
		if err22 != nil {
			m2 = false
		}
	
		conn2.Close()

		}

		if (m1 && m2 && m3 ) == false{
			prop_c = createRegPropDist(ej1,m1,m2,m3)
		}
		prop_c = createRegPropDist(ej1,m1,m2,m3)

		distribution = prop_c

		//conexion al NameNode
		var conn *grpc.ClientConn
		conn, error := grpc.Dial(":9001", grpc.WithInsecure())
		if error != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := NewChatServiceClient(conn)

		//Escribir en el log


		//Algoritmo Ricart y agrawala


		response, err2 := c.WriteLog(context.Background(), &LogInfo{Log: distribution, Nombre: libro,
			Partes: int64(cantidadMensajes)})
			if err2 != nil {
				log.Fatalf("Error when calling SayHello: %s", err2)
			}
			log.Printf("Respuesta NameNode %s", response.Body)

	}
	
	

	//Distribuir los chunks

	

	var sep []string
	cont:= 0
	//fmt.Println(len(chunkList))
	sep = strings.Split(distribution,"%%%")
	for i := 0; i < cantidadMensajes; i++ {

		/*
		fmt.Println(cont)
		if cont >= cantidadMensajes{
			fmt.Println("Entre al break")
			break
		
		}
		*/

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
	currentChunkFileName := "./DB/" + in.Body
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
