package controllers

import (
	"net/http"
	"time"

	"commitment-stake/config"
	"commitment-stake/models"

	"github.com/gin-gonic/gin"
)

type CreateEventRequest struct {
	Name          string    `json:"name" binding:"required"`
	DepositAmount float64   `json:"deposit_amount" binding:"required,gt=0"`
	StartTime     time.Time `json:"start_time" binding:"required"`
	EndTime       time.Time `json:"end_time" binding:"required"`
}

func CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ hoặc thiếu trường bắt buộc"})
		return
	}

	if !req.EndTime.After(req.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thời gian kết thúc phải diễn ra sau thời gian bắt đầu"})
		return
	}

	event := models.Event{
		Name:          req.Name,
		DepositAmount: req.DepositAmount,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Status:        "upcoming",
	}

	if err := config.DB.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu sự kiện vào Database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo sự kiện thành công",
		"data":    event,
	})
}

func GetEvents(c *gin.Context) {
	var events []models.Event
	if err := config.DB.Order("start_time asc").Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi truy vấn danh sách sự kiện"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": events})
}