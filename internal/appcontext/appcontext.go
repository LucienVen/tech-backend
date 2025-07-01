package appcontext

import "github.com/LucienVen/tech-backend/internal/db"

// AppContext 聚合全局依赖
type AppContext struct {
	DB    db.DB
	Redis *db.RedisClient
}
