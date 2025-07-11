package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LucienVen/tech-backend/internal/app"
)

func main() {
	// 打印当前工作目录
	if wd, err := os.Getwd(); err != nil {
		log.Printf("获取工作目录失败: %v", err)
	} else {
		log.Printf("当前工作目录: %s", wd)
	}

	// 创建应用实例
	application := app.NewApplication()

	// 启动应用
	if err := application.Start(); err != nil {
		log.Printf("应用启动失败: %v", err)
		os.Exit(1)
	}

	// 等待中断信号2
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("正在关闭服务...")

	// 创建关闭上下文
	ctx := context.Background()

	// 优雅关闭
	if err := application.Shutdown(ctx); err != nil {
		log.Printf("服务关闭错误: %v", err)
		os.Exit(1)
	}

	log.Println("服务已关闭")
}
