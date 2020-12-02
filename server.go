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
	//Solicitar tipo de algortimo
	fmt.Println("Indique el tipo de algoritmo")
	fmt.Println("0: Centralizado")
	fmt.Println("1: Distriuido")
	
	fmt.Scanln(&serverType)

	fmt.Println("Servidor Arriba")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9003))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{Log:make(map[string]string), Mode: serverType, Id: "1", State: 0}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
