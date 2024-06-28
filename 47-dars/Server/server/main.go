package main

import (
	"fmt"
	"log"
	t "mymode/protos/transport/protos"
	w "mymode/protos/weather/protos"
	"mymode/service"
	"mymode/storage"
	"mymode/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	serverW := service.NewWeatherRepo(*postgres.NewWeatherRepo(db))
	serverT := service.NewTransportRepo(*postgres.NewTransportRepo(db))

	grpc := grpc.NewServer()

	w.RegisterWeatherServer(grpc, serverW)
	t.RegisterTransportServer(grpc, serverT)

	if err = grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server listening at %v", listener.Addr())
}
