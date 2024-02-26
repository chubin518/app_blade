package controller

import (
	"app_blade/internal/model"
	"app_blade/internal/service"
	"app_blade/pkg/logging"
	"app_blade/pkg/response"

	"github.com/gin-gonic/gin"
)

func NewBladeProductController(service service.BladeProductService, logger logging.Provider) *BladeProductController {
	return &BladeProductController{
		logger:  logger,
		service: service,
	}
}

type BladeProductController struct {
	logger  logging.Provider
	service service.BladeProductService
}

func (c *BladeProductController) List(ctx *gin.Context) {
	products, err := c.service.List(ctx)
	if err != nil {
		c.logger.ErrorfContext(ctx, "Failed to list products: %v", err)
		response.ServiceError().JSON(ctx)
		return
	}
	response.OK().WithData(products).JSON(ctx)
}

func (c *BladeProductController) Save(ctx *gin.Context) {
	m := &model.BladeProduct{}
	if err := ctx.BindJSON(m); err != nil {
		c.logger.ErrorfContext(ctx.Request.Context(), "bind json failed: %v", err)
		return
	}
	if err := c.service.SaveOrUpdate(ctx, m); err != nil {
		c.logger.ErrorfContext(ctx.Request.Context(), "save product failed: %v", err)
		response.ServiceError().JSON(ctx)
		return
	}
	response.OK().JSON(ctx)
}
