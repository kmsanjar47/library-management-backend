package utils

import "github.com/gin-gonic/gin"

func NewResponse(status string, message string, result any) gin.H {
	return gin.H{"status": status, "message": message, "result": result}
}
