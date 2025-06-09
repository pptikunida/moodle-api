package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/moodle-api/model/web"
	"github.com/rizkycahyono97/moodle-api/services"
	"log"
	"net/http"
)

type MoodleController struct {
	moodleService services.MoodleService
}

func NewMoodleController(moodleService services.MoodleService) *MoodleController {
	return &MoodleController{moodleService: moodleService}
}

func (s *MoodleController) CheckStatus(c *gin.Context) {
	result, err := s.moodleService.CheckStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "OK",
		Message: "OK",
		Data:    result,
	})
}

func (s *MoodleController) CreateUser(c *gin.Context) {
	var req web.MoodleUserCreateRequest

	fmt.Println("[DEBUG] Received Body:", req) // log

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[CreateUser] Error: %v", err) // log
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "INVALID_PARAMS",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	log.Printf("[CreateUser] Received Request: %+v", req) // log

	// Call service
	result, err := s.moodleService.CreateUser(req)
	if err != nil {
		log.Println("[CreateUser] Error:", err)
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "OK",
		Message: "OK",
		Data:    result,
	})
}

func (s *MoodleController) GetUserByField(c *gin.Context) {
	var req web.MoodleUserGetByFieldRequest

	// Bind JSON request body ke struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "INVALID_PARAMS",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// service
	users, err := s.moodleService.GetUserByField(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "OK",
		Message: "OK",
		Data:    users,
	})
}
