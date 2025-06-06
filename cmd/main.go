package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/LucienVen/tech-backend/internal/app"
	"github.com/LucienVen/tech-backend/pkg/log"
)

func main() {
	// 创建应用实例
	application := app.NewApplication()

	// 启动应用
	if err := application.Start(); err != nil {
		log.Errorf("应用启动失败: %v", err)
		os.Exit(1)
	}

	// 等待中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Info("正在关闭服务...")

	// 创建关闭上下文
	ctx := context.Background()

	// 优雅关闭
	if err := application.Shutdown(ctx); err != nil {
		log.Errorf("服务关闭错误: %v", err)
		os.Exit(1)
	}

	log.Info("服务已关闭")
}
