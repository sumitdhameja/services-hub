package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitdhameja/services-hub/internal/dto"
	"github.com/sumitdhameja/services-hub/internal/logger"
)

var ErrReplyUnknown = dto.ReplyError("Unknown error")

const UnhandlerError = "[UNHANDLED_ERROR]:"
const AppError = "[APP_ERROR]:"

// GinError middleware
func GinError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if errors := c.Errors.ByType(gin.ErrorTypeAny); len(errors) > 0 {
			err := errors[0].Err
			if err, ok := err.(*Error); ok {
				logger.Error(AppError, err)
				c.AbortWithStatusJSON(err.Code, err.ToReply())
				return
			}
			logger.Error(UnhandlerError, err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrReplyUnknown)
			return
		}
	}
}
