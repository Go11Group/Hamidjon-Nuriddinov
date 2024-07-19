package redisDb

import (
	"context"
	pb "homework/genproto"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type NewRedis struct{
	Rdb *redis.Client
}

func NewRedisRepo(rdb *redis.Client)*NewRedis{
	return &NewRedis{
		Rdb: rdb,
	}
}

func(R *NewRedis) CreateCenter(ctx context.Context, req *pb.Center)(*pb.CenterResp, error){
	id := uuid.New()
	err := R.Rdb.HSet(ctx, id.String(), "name:", req.Name, "description:", req.Description, 
			"address:", req.Address, "phone:", req.Phone).Err()
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return &pb.CenterResp{
		Id: id.String(),
		Name: req.Name,
		Description: req.Description,
		Address: req.Address,
		Phone: req.Phone,
	}, nil
}

func(R *NewRedis) GetCenterById(ctx context.Context, req *pb.CenterId)(*pb.CenterResp){
	center := R.Rdb.HGetAll(ctx, req.Id).Val()
	return &pb.CenterResp{
		Id: req.Id,
		Name: center["name:"],
		Description: center["description:"],
		Address: center["address:"],
		Phone: center["phone:"],
	}
}

func(R *NewRedis) UpdateCenter(ctx context.Context, req *pb.CenterResp)(*pb.UpdateResp, error){
	err := R.Rdb.HSet(ctx, req.Id, "name:", req.Name, "description:", req.Description, 
			"address:", req.Address, "phone:", req.Phone).Err()
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return &pb.UpdateResp{
		Id: req.Id,
		Name: req.Name,
		Description: req.Description,
		Address: req.Address,
		Phone: req.Phone,		
	}, nil
}

func(R *NewRedis) DeleteCenter(ctx context.Context, req *pb.CenterId)(*pb.Status, error){
	err := R.Rdb.HDel(ctx, req.Id, "name:", "description:", "address:", "phone:").Err()
	if err != nil{
		return &pb.Status{
			Status: false,
			Message: "Ma'lumotlar o'chirilmadi.",
		}, err
	}
	return &pb.Status{
		Status: true,
		Message: "Ma'lumotlar muvaffaqiyatli o'chirildi.",
	}, nil
}