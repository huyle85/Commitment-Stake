package controllers

import (
	"net/http"

	"commitment-stake/config"
	"commitment-stake/models"

	"github.com/gin-gonic/gin"
)

type QRRequest struct {
	EventID       uint   `json:"event_id" binding:"required"`
	WalletAddress string `json:"wallet_address" binding:"required"`
}

func GenerateQR(c *gin.Context) {
	var req QRRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu thông tin event_id hoặc wallet_address"})
		return
	}

	var registration models.Registration
	result := config.DB.Where("event_id = ? AND wallet_address = ?", req.EventID, req.WalletAddress).First(&registration)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy thông tin đăng ký của ví này"})
		return
	}

	
	if registration.Status != "staked" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Ví này chưa đặt cọc, không thể cấp mã QR"})
		return
	}

	signature := req.WalletAddress + "-approved-by-server"

	c.JSON(http.StatusOK, gin.H{
		"message":        "Cấp mã QR thành công",
		"event_id":       req.EventID,
		"wallet_address": req.WalletAddress,
		"signature":      signature, 
	})
}