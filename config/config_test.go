package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// 加载配置，使用默认参数
	Load(true)
	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("GetConfig 返回 nil")
	}
	if cfg.AppEnv == "" {
		t.Error("AppEnv 应该有默认值")
	}
	if cfg.AppName == "" {
		t.Error("AppName 应该有默认值")
	}
	if cfg.HTTPPort == 0 {
		t.Error("HTTPPort 应该有默认值")
	}
	// 可选：检查部分数据库字段
	if cfg.DBType == "" {
		t.Log("DBType 未设置，默认为空字符串")
	}
}

func TestPrintConfigFromCmdEnv(t *testing.T) {
	wd, _ := os.Getwd()
	t.Logf("当前工作目录: %s", wd)
	t.Logf("cmd/.env 是否存在: %v", fileExists("cmd/.env"))
	Load(false, "/Users/liangliangtoo/code/tech-backend/cmd/.env")
	cfg := GetConfig()
	t.Logf("Config from cmd/.env: %+v", cfg)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
