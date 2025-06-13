package api

import (
	"github.com/LucienVen/tech-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Router 路由管理器
type Router struct {
	engine *gin.Engine
}

// NewRouter 创建路由管理器
func NewRouter() *Router {
	// 创建 Gin 引擎
	engine := gin.New()

	// 使用中间件
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	return &Router{
		engine: engine,
	}
}

// RegisterRoutes 注册所有路由
func (r *Router) RegisterRoutes() {
	// 基础路由组
	base := r.engine.Group("/api")
	{
		// 健康检查
		base.GET("/health", HealthCheck)
		base.GET("/ping", Ping)
	}

	// TODO: 添加更多路由组
	// 例如：
	// v1 := base.Group("/v1")
	// {
	//     v1.GET("/users", GetUsers)
	//     v1.POST("/users", CreateUser)
	// }
}

// GetEngine 获取 Gin 引擎
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
