package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouterAPIV1(router *gin.RouterGroup, db *gorm.DB) {

	appServiceAPI := NewAppServiceAPI(db)
	router.GET("users/:user_id/services", appServiceAPI.GetAllService)
	router.GET("users/:user_id/services/:id", appServiceAPI.GetService)

}
