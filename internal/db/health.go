package db

import (
	"context"
	"time"

	"github.com/LucienVen/tech-backend/pkg/log"
)

// HealthChecker 健康检查器
type HealthChecker struct {
	db     DB
	ticker *time.Ticker
	done   chan struct{}
}

// NewHealthChecker 创建健康检查器
func NewHealthChecker(db DB) *HealthChecker {
	return &HealthChecker{
		db:   db,
		done: make(chan struct{}),
	}
}

// Start 启动健康检查
func (h *HealthChecker) Start(ctx context.Context, interval time.Duration) {
	h.ticker = time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-h.ticker.C:
				// 只有 db 实现了 Ping/Connect 方法时才调用
				if pinger, ok := h.db.(interface{ Ping() error }); ok {
					dbType := "unknown"
					switch h.db.(type) {
					case *GormDB:
						dbType = "MySQL"
					case *GormPGDB:
						dbType = "PostgreSQL"
					}

					if err := pinger.Ping(); err != nil {
						log.Errorf("[%s] 数据库心跳检测失败: %v", dbType, err)
						if connector, ok := h.db.(interface{ Connect() error }); ok {
							if err := connector.Connect(); err != nil {
								log.Errorf("[%s] 数据库重连失败: %v", dbType, err)
							} else {
								log.Infof("[%s] 数据库重连成功", dbType)
							}
						}
					} else {
						log.Infof("[%s] 数据库心跳检测正常", dbType)
					}
				}
			case <-ctx.Done():
				h.Stop()
				return
			case <-h.done:
				return
			}
		}
	}()
}

// Stop 停止健康检查
func (h *HealthChecker) Stop() {
	if h.ticker != nil {
		h.ticker.Stop()
	}
	close(h.done)
}

// Shutdown 实现 Shutdownable 接口
func (h *HealthChecker) Shutdown(ctx context.Context) error {
	h.Stop()
	return nil
}
