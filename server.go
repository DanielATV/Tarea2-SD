package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
	"google.golang.org/grpc"
	
)

func main() {

	var serverType int
	var nodeID string
	//var adress int
	//Solicitar tipo de algortimo
	fmt.Println("Indique el tipo de algoritmo")
	fmt.Println("0: Centralizado")
	fmt.Println("1: Distriuido")
	
	fmt.Scanln(&serverType)

	fmt.Println("Id del nodo")
	
	fmt.Scanln(&nodeID)

	/*

	//Testeo Local
	if nodeID == "1"{
		adress =  9000
	} else if nodeID == "2"{
		adress = 9002
	} else if nodeID == "3"{
		adress = 9003
	} else {
		adress = 9001
	}

	*/

	fmt.Println("Servidor Arriba")

	

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{Log:make(map[string]string), Mode: serverType, Id: nodeID, State: 0}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
