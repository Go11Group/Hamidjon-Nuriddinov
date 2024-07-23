package service

import (
	"fmt"
	pb "homework/genproto/course"
	"homework/pkg/logger"
	redisDb "homework/storage/redis"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type NewCourseService struct{
	pb.UnimplementedCourseServiceServer
	Rdb *redisDb.NewRedis
	Logger *slog.Logger
}

func NewCourseServiceRepo(rdb *redis.Client)*NewCourseService{
	return &NewCourseService{
		Rdb: redisDb.NewRedisRepo(rdb),
		Logger: logger.NewLogger(),
	}
}

func(C *NewCourseService) CreateCourse(ctx context.Context, req *pb.Course)(*pb.CourseResp, error){
	resp, err  := C.Rdb.CreateCourse(ctx, req)
	if err != nil{
		C.Logger.Error(fmt.Sprintf("Databazadan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func(C *NewCourseService) GetCourseById(ctx context.Context, req *pb.CourseId)(*pb.CourseResp, error){
	resp, err  := C.Rdb.GetCourseById(ctx, req)
	if err != nil{
		C.Logger.Error(fmt.Sprintf("Databazadan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func(C *NewCourseService) UpdateCourse(ctx context.Context, req *pb.CourseResp)(*pb.UpdateResp, error){
	resp, err  := C.Rdb.UpdateCourse(ctx, req)
	if err != nil{
		C.Logger.Error(fmt.Sprintf("Databazadan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func(C *NewCourseService) DeleteCourse(ctx context.Context, req *pb.CourseId)(*pb.Status, error){
	resp, err  := C.Rdb.DeleteCourse(ctx, req)
	if err != nil{
		C.Logger.Error(fmt.Sprintf("Databazadan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}