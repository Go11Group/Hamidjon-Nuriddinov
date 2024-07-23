package service

import (
	"context"
	"fmt"
	pb "homework/genproto/center"
	"homework/pkg/logger"
	redisDb "homework/storage/redis"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type CenterService struct {
	pb.UnimplementedCenterServiceServer
	Rdb    redisDb.NewRedis
	Logger *slog.Logger
}

func NewCenterService(rdb *redis.Client) *CenterService {
	return &CenterService{
		Rdb:    *redisDb.NewRedisRepo(rdb),
		Logger: logger.NewLogger(),
	}
}

func (C *CenterService) CreateCenter(ctx context.Context, req *pb.Center) (*pb.CenterResp, error) {
	resp, err := C.Rdb.CreateCenter(ctx, req)
	if err != nil {
		C.Logger.Error(fmt.Sprintf("Redisdan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (C *CenterService) GetCenterById(ctx context.Context, req *pb.CenterId) (*pb.CenterResp, error) {
	resp := C.Rdb.GetCenterById(ctx, req)
	return resp, nil
}

func (C *CenterService) UpdateCenter(ctx context.Context, req *pb.CenterResp) (*pb.UpdateResp, error) {
	resp, err := C.Rdb.UpdateCenter(ctx, req)
	if err != nil {
		C.Logger.Error(fmt.Sprintf("Redisdan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (C *CenterService) DeleteCenter(ctx context.Context, req *pb.CenterId) (*pb.Status, error) {
	resp, err := C.Rdb.DeleteCenter(ctx, req)
	if err != nil {
		C.Logger.Error(fmt.Sprintf("Redisdan ma'lumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}
