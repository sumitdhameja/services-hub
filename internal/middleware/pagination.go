package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitdhameja/services-hub/internal/dto"
)

var (
	Limit = "Limit"
	Page  = "Page"
)

func Paginate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p dto.Pageable
		if c.Bind(&p) == nil {
			c.Set("page_options", p)
		}
		c.Next()
	}
}
