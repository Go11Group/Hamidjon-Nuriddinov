package handler

import (
	center "api/genproto"
	"api/pkg"
	"api/pkg/logger"
	"log/slog"
)

type Handler struct{
	CenterService center.CenterServiceClient
	Logger *slog.Logger
}

func NewHandler()*Handler{
	return &Handler{
		CenterService: pkg.NewCenterClient(),
		Logger: logger.NewLogger(),
	}
}