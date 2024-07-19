package handler

import (
	pb "api/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(h *Handler) CreateCenter(c *gin.Context){
	req := pb.Center{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateCenter error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	resp, err := h.CenterService.CreateCenter(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateCenter request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) GetCenterById(c *gin.Context){
	resp, err := h.CenterService.GetCenterById(c, &pb.CenterId{Id: c.Param("id")})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetCenterById request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) UpdateCenter(c *gin.Context){
	req := pb.CenterResp{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("UpdateCenter error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return 
	}

	resp, err := h.CenterService.UpdateCenter(c, &req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("UpdateCenter request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) DeleteCenter(c *gin.Context){
	resp, err := h.CenterService.DeleteCenter(c, &pb.CenterId{Id: c.Param("id")})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("DeleteCenter request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}