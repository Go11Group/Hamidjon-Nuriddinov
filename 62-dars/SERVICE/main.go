package main

import (
	"fmt"
	center "homework/genproto/center"
	course "homework/genproto/course"
	"homework/service"
	redisDb "homework/storage/redis"
	"log"
	"net"

	"google.golang.org/grpc"
)


func main(){
	listener, err := net.Listen("tcp", ":7777")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()

	rdb := redisDb.ConnectRedis()

	centerService := service.NewCenterService(rdb)
	courseService := service.NewCourseServiceRepo(rdb)

	service := grpc.NewServer()

	center.RegisterCenterServiceServer(service, centerService)
	course.RegisterCourseServiceServer(service, courseService)

	fmt.Printf("Server listen port: %v", "7777")
	if err = service.Serve(listener); err != nil{
		log.Fatal(err)
	}
}