package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rizkycahyono97/moodle-api/model/web"
	"github.com/rizkycahyono97/moodle-api/services"
	"github.com/rizkycahyono97/moodle-api/utils/validation"
)

type MoodleController struct {
	moodleService services.MoodleService
}

func NewMoodleController(moodleService services.MoodleService) *MoodleController {
	return &MoodleController{moodleService: moodleService}
}

func (s *MoodleController) CoreWebserviceGetSiteInfo(c *gin.Context) {
	result, err := s.moodleService.CoreWebserviceGetSiteInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: err.Error(),
			Data:    nil,
		})
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "OK",
		Data:    result,
	})
}

func (s *MoodleController) CoreUserCreateUsers(c *gin.Context) {
	var req web.MoodleUserCreateRequest

	fmt.Println("[DEBUG] Received Body:", req) // log

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[CoreUserCreateUsers] Error: %v", err) // log
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	log.Printf("[CoreUserCreateUsers] Received Request: %+v", req) // log

	// Call service
	result, err := s.moodleService.CoreUserCreateUsers(req)
	if err != nil {
		log.Println("[CoreUserCreateUsers] Error:", err)
		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "OK",
		Data:    result,
	})
}

func (s *MoodleController) CoreUserGetUsersByField(c *gin.Context) {
	var req web.MoodleUserGetByFieldRequest

	// Bind JSON request body ke struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Format body request tidak valid.",
			Data:    err.Error(),
		})
		return
	}

	// service
	users, err := s.moodleService.CoreUserGetUsersByField(req)
	if err != nil {
		log.Printf("[DIAGNOSA] Controller menerima error. Tipe: %T, Isi: %v", err, err)
		if errors.Is(err, validation.ErrNotFound) {
			c.JSON(http.StatusNotFound, web.ApiResponse{
				Code:    "404",
				Message: err.Error(), // Menggunakan pesan dari variabel ErrNotFound
			})
			return
		}

		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "Terjadi kesalahan pada server.",
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "OK",
		Data:    users,
	})
}

func (s *MoodleController) CoreUserUpdateUsers(c *gin.Context) {
	var req []web.MoodleUserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	//periksa jika ada data
	err := s.moodleService.CoreUserUpdateUsers(req)
	if err != nil {
		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "An internal error occurred",
			Data:    err.Error(), // Kirim pesan error internal untuk debug
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "Users updated successfully",
	})
}

func (s *MoodleController) UserSync(c *gin.Context) {
	var req web.MoodleUserSyncRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: " + err.Error(),
		})
		return
	}

	// Panggil service
	err := s.moodleService.UserSync(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// Berhasil
	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "User synced successfully",
		Data:    nil,
	})
}

func (s *MoodleController) CoreRoleAssignRoles(c *gin.Context) {
	var req web.MoodleRoleAssignRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: " + err.Error(),
		})
		return
	}

	log.Printf("[DEBUG] CoreRoleAssignRoles Controller: Menerima request %+v", req)

	// panggil service
	if err := s.moodleService.CoreRoleAssignRoles(req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "User assigned successfully",
	})
}

func (s *MoodleController) CoreCourseCreateCourses(c *gin.Context) {
	var req web.MoodleCoreCourseCreateCoursesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: " + err.Error(),
		})
		return
	}

	courses, err := s.moodleService.CoreCourseCreateCourses(req)
	if err != nil {
		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "An internal error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "Courses created successfully",
		Data:    courses,
	})
}

func (s *MoodleController) EnrolManualEnrolUsers(c *gin.Context) {
	var req web.MoodleManualEnrollRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: " + err.Error(),
			Data:    nil,
		})
		return
	}

	if err := s.moodleService.EnrolManualEnrolUsers(req); err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "An internal error occurred" + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "Users enrolled successfully",
		Data:    nil,
	})
}

func (s *MoodleController) CreateCourseWithEnrollUser(c *gin.Context) {
	var req web.MoodleCreateCourseWithEnrollUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: ",
			Data:    err.Error(),
		})
		return
	}

	newCourseInfo, err := s.moodleService.CreateCourseWithEnrollUser(req)
	if err != nil {
		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "Terjadi kesalahan internal pada server.",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "201",
		Message: "Kursus berhasil di buat dan penggunal telah didaftarkan.",
		Data:    newCourseInfo,
	})

}

func (s *MoodleController) CoreCourseCreateCategories(c *gin.Context) {
	var req web.MoodleCreateCategoriesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: " + err.Error(),
			Data:    nil,
		})
		return
	}

	newCategories, err := s.moodleService.CoreCourseCreateCategories(req)
	if err != nil {
		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "An internal error occurred",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, web.ApiResponse{
		Code:    "201",
		Message: "Kategori berhasil dibuat",
		Data:    newCategories,
	})
}

func (s *MoodleController) CoreCourseUpdateCategories(c *gin.Context) {
	var req web.MoodleUpdateCategoriesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ApiResponse{
			Code:    "400",
			Message: "Data yang dikirim tidak valid: " + err.Error(),
			Data:    nil,
		})
		return
	}

	err := s.moodleService.CoreCourseUpdateCategories(req)
	if err != nil {
		if moodleErr, ok := err.(*web.MoodleException); ok {
			c.JSON(http.StatusBadRequest, web.ApiResponse{
				Code:    moodleErr.ErrorCode,
				Message: moodleErr.Message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "An internal error occurred",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, web.ApiResponse{
		Code:    "200",
		Message: "Kategori berhasil diperbarui",
	})
}

func (s *MoodleController) ServeSwaggerSpec(c *gin.Context) {
	//baca file
	file, err := os.ReadFile("./apispec.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ApiResponse{
			Code:    "500",
			Message: "Could not read API spec file",
		})
		return
	}
	c.Data(http.StatusOK, "application/json", file)
}
