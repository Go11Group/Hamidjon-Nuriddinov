package pkg

import (
	center "api/genproto/center"
	"api/genproto/course"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCenterClient() center.CenterServiceClient{
	listener, err := grpc.NewClient(":7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println(err)
		return nil
	}

	conn := center.NewCenterServiceClient(listener)
	return conn
}

func NewCourseClient() course.CourseServiceClient{
	listener, err := grpc.NewClient(":7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println(err)
		return nil
	}

	conn := course.NewCourseServiceClient(listener)
	return conn
}
