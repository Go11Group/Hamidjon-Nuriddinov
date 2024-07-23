package handler

import (
	center "api/genproto/center"
	"api/genproto/course"
	"api/pkg"
	"api/pkg/logger"
	"log/slog"
)

type Handler struct{
	CenterService center.CenterServiceClient
	CourseService course.CourseServiceClient
	Logger *slog.Logger
}

func NewHandler()*Handler{
	return &Handler{
		CenterService: pkg.NewCenterClient(),
		CourseService: pkg.NewCourseClient(),
		Logger: logger.NewLogger(),
	}
}