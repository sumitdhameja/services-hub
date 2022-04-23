package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitdhameja/services-hub/internal/dto"
	"github.com/sumitdhameja/services-hub/internal/errors"
	"github.com/sumitdhameja/services-hub/internal/services"
	"gorm.io/gorm"
)

// appServiceAPI is a service private
type appServiceAPI struct {
	service services.AppService
}

// NewAppServiceAPI create userService
func NewAppServiceAPI(db *gorm.DB) AppServiceAPI {
	return &appServiceAPI{service: services.NewAppService(db)}
}

// AppServiceAPI interface
type AppServiceAPI interface {
	GetAllService(c *gin.Context)
	GetService(c *gin.Context)
}

func (p *appServiceAPI) GetAllService(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	services, err := p.service.GetAllService(userID, dto.Pageable{})
	if err != nil {
		ctx.Error(errors.NewError(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataResponse{Data: services})
}

// GetService return only one Service
func (p *appServiceAPI) GetService(ctx *gin.Context) {
	serviceID := ctx.Param("id")
	userID := ctx.Param("user_id")

	service, err := p.service.GetService(userID, serviceID)
	if err != nil {
		ctx.Error(errors.NewError(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dto.DataResponse{Data: service})
}
