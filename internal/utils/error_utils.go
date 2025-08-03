package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/kunjbosamia/sgnp-ticket-booking/internal/constants"
)

func RespondWithError(c *gin.Context, err constants.ErrorResponse) {
	c.JSON(err.Status, gin.H{
		"error_code":    err.Code,
		"error_message": err.Message,
	})
}
