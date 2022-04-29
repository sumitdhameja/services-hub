package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sumitdhameja/services-hub/internal/errors"
)

func Validate() gin.HandlerFunc {

	return func(c *gin.Context) {
		userID := c.Param("user_id")

		_, err := uuid.Parse(userID)
		if err != nil {
			c.Error(errors.NewError(http.StatusBadRequest, err.Error()))
			c.Abort()
			return
		}

		c.Next()
	}
}
