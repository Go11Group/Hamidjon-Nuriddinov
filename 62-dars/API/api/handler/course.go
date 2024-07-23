package handler

import (
	pb "api/genproto/course"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	req := pb.Course{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateCourse error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.CourseService.CreateCourse(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("CreateCourse request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCourseById(c *gin.Context) {
	resp, err := h.CourseService.GetCourseById(c, &pb.CourseId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetCourseById request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	req := pb.CourseResp{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateCourse error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.CourseService.UpdateCourse(c, &req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("UpdateCourse request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	resp, err := h.CourseService.DeleteCourse(c, &pb.CourseId{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteCourse request error: %v", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
