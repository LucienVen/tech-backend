package testutil

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/LucienVen/tech-backend/config"
)

// LoadTestConfig 统一加载cmd/.env配置，返回*config.Config
func LoadTestConfig(t *testing.T) *config.Config {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")
	configDir := filepath.Join(projectRoot, "cmd")
	configFile := filepath.Join(configDir, ".env")
	t.Logf("项目根目录: %s", projectRoot)
	t.Logf("配置文件路径: %s", configFile)
	wd, _ := os.Getwd()
	t.Logf("当前工作目录: %s", wd)
	config.Load(true, configFile)
	cfg := config.GetConfig()
	t.Logf("config.GetConfig() 返回: %+v", cfg)
	return cfg
}
