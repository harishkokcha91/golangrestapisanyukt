package controllers

import (
	"github.com/gin-gonic/gin"
)

func CustomResponse(message string, status bool, data interface{}) gin.H {
	return gin.H{"message": message, "status": status, "data": data}
}
