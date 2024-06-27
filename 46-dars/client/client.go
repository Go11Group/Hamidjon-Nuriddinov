package main

import (
	"context"
	"log"
	t "mymode/protos/transport/protos"
	w "mymode/protos/weather/protos"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	wServ := w.NewWeatherClient(conn)
	tServ := t.NewTransportClient(conn)
	GetCurrentWeather(wServ)
	GetBusSchedule(tServ)
}

func GetCurrentWeather(wServ w.WeatherClient){
	t := w.Time{
		Time: "14:00",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := wServ.GetCurrentWeather(ctx, &t)
	if err != nil{
		log.Println(err)
	}
	log.Println(resp)
}

func GetBusSchedule(tServ t.TransportClient){
	number := t.Number{
		Number: 15,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := tServ.GetBusSchedule(ctx, &number)
	if err != nil{
		log.Println(err)
	}
	log.Println(resp)
}

