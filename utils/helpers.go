package utils

import (
	"github.com/gin-gonic/gin"
	"go/types"
	"gorm.io/gorm"
)

func HttpErrorResponse(c *gin.Context, httpCode int, e error, message string) {
	c.JSON(httpCode, map[string]interface{}{
		"msg_code":    "error",
		"msg_content": message,
		"msg_data":    types.Interface{},
	})
}

func HttpSuccessResponse(c *gin.Context, httpCode int, data interface{}, message string) {
	c.JSON(httpCode, map[string]interface{}{
		"msg_code":    "success",
		"msg_content": message,
		"msg_data":    data,
	})
}

func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
