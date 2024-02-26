package controller

import (
	"app_blade/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(InitializeController,
	NewBladeUserController,
	NewBladeProductController,
)

// InitializeController 初始化路由
func InitializeController(product *BladeProductController, user *BladeUserController) web.InitializeRouter {
	return func(route *gin.Engine) {
		route.GET("/product/list", product.List)
		route.POST("/product/save", product.Save)
		route.GET("/user/detail", user.Get)
		route.POST("/user/save", user.Save)
	}
}
