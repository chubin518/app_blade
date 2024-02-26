package controller

import (
	"app_blade/internal/model"
	"app_blade/internal/service"
	"app_blade/pkg/logging"
	"app_blade/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func NewBladeUserController(service service.BladeUserService, logger logging.Provider) *BladeUserController {
	return &BladeUserController{
		service: service,
		logger:  logger,
	}
}

type BladeUserController struct {
	logger  logging.Provider
	service service.BladeUserService
}

func (c *BladeUserController) Get(ctx *gin.Context) {
	id := cast.ToInt(ctx.Query("id"))
	bladeUser, err := c.service.Get(ctx.Request.Context(), id)
	if err != nil {
		c.logger.ErrorfContext(ctx.Request.Context(), "get blade user error: %v", err)
		response.ServiceError().JSON(ctx)
		return
	}
	response.OK().WithData(bladeUser).JSON(ctx)
}

func (c *BladeUserController) Save(ctx *gin.Context) {
	m := &model.BladeUser{}
	if err := ctx.BindJSON(m); err != nil {
		c.logger.ErrorfContext(ctx.Request.Context(), "bind json failed: %v", err)
		return
	}
	if err := c.service.SaveOrUpdate(ctx.Request.Context(), m); err != nil {
		c.logger.ErrorfContext(ctx.Request.Context(), "save user failed: %v", err)
		response.ServiceError().JSON(ctx)
		return
	}
	response.OK().JSON(ctx)
}
