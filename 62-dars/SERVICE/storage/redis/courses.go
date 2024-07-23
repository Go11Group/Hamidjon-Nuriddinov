package redisDb

import (
	"context"
	pb "homework/genproto/course"
	"strconv"

	"github.com/google/uuid"
)

func (C *NewRedis) CreateCourse(ctx context.Context, req *pb.Course) (*pb.CourseResp, error) {
	id := uuid.New()
	err := C.Rdb.HSet(ctx, id.String(), "center_id:", req.CenterId, "name:", req.Name, "description:", req.Description, "teacher:", req.Teacher, "continuity:", req.Continuity).Err()
	if err != nil {
		return nil, err
	}
	return &pb.CourseResp{
		Id:          id.String(),
		CenterId:    req.CenterId,
		Name:        req.Name,
		Description: req.Description,
		Teacher:     req.Teacher,
		Continuity:  req.Continuity,
	}, nil
}

func (C *NewRedis) GetCourseById(ctx context.Context, req *pb.CourseId) (*pb.CourseResp, error) {
	course := C.Rdb.HGetAll(ctx, req.Id).Val()
	cont, err := strconv.Atoi(course["continuity:"])
	if err != nil{
		return nil, err
	}
	return &pb.CourseResp{
		Id:          req.Id,
		CenterId:    course["center_id:"],
		Name:        course["name:"],
		Description: course["description:"],
		Teacher:     course["teacher:"],
		Continuity:  int32(cont),
	}, nil
}

func(C *NewRedis) UpdateCourse(ctx context.Context, req *pb.CourseResp)(*pb.UpdateResp, error){
	err := C.Rdb.HSet(ctx, req.Id, "center_id:", req.CenterId, "name:", req.Name, "description:", req.Description, "teacher:", req.Teacher, "continuity:", req.Continuity).Err()
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResp{
		Id:          req.Id,
		CenterId:    req.CenterId,
		Name:        req.Name,
		Description: req.Description,
		Teacher:     req.Teacher,
		Continuity:  req.Continuity,
	}, nil
}

func(C *NewRedis) DeleteCourse(ctx context.Context, req *pb.CourseId)(*pb.Status, error){
	err := C.Rdb.HDel(ctx, req.Id, "center_id:", "name:", "description:", "teacher:", "continuity:").Err()
	if err != nil{
		return &pb.Status{
			Status: false,
			Message: "Ma;lumotlar o'chirilmadi",
		}, err
	}
	return &pb.Status{
		Status: true,
		Message: "Ma'lumotlar muvaffaqiyatli o'chirildi",
	}, nil
}
