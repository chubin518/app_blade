package web

import "github.com/gin-gonic/gin"

// InitializeRouter 初始化web路由
type InitializeRouter func(*gin.Engine)
