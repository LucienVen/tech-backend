package db

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/LucienVen/tech-backend/config"
	"gorm.io/gorm"
)

type TestPGModel struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func loadPGConfig(t *testing.T) *config.Config {
	// 获取项目根目录
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")
	configDir := filepath.Join(projectRoot, "cmd")
	configFile := filepath.Join(configDir, ".env")
	t.Logf("项目根目录: %s", projectRoot)
	t.Logf("配置目录: %s", configDir)
	t.Logf("配置文件路径: %s", configFile)
	// 打印当前工作目录和配置文件完整路径
	wd, _ := os.Getwd()
	t.Logf("当前工作目录: %s", wd)
	t.Logf("尝试读取配置文件: %s", configFile)
	// 加载配置
	config.Load(true, configFile)
	cfg := config.GetConfig()
	t.Logf("config.GetConfig() 返回: %+v", cfg)
	t.Logf("PG配置: host=%s port=%s user=%s db=%s", cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGName)
	// 新增：打印环境变量，辅助排查
	t.Logf("ENV: PG_HOST=%s, PG_PORT=%s, PG_USER=%s, PG_PASS=%s, PG_NAME=%s", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_NAME"))
	return cfg
}

func TestGormPGDB_ConnectAndPing(t *testing.T) {
	cfg := loadPGConfig(t)
	db := NewGormPGDB(cfg)
	if err := db.Connect(); err != nil {
		t.Fatalf("连接失败: %v", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("Ping 失败: %v", err)
	}
	db.Close()
}

func TestGormPGDB_AutoMigrate(t *testing.T) {
	cfg := loadPGConfig(t)
	db := NewGormPGDB(cfg)
	if err := db.Connect(); err != nil {
		t.Fatalf("连接失败: %v", err)
	}
	defer db.Close()
	if err := db.AutoMigrate(&TestPGModel{}); err != nil {
		t.Fatalf("AutoMigrate 失败: %v", err)
	}
}

func TestGormPGDB_Transaction(t *testing.T) {
	cfg := loadPGConfig(t)
	db := NewGormPGDB(cfg)
	if err := db.Connect(); err != nil {
		t.Fatalf("连接失败: %v", err)
	}
	defer db.Close()
	err := db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&TestPGModel{Name: "事务测试"}).Error
	})
	if err != nil {
		t.Fatalf("事务失败: %v", err)
	}
}

func TestGormPGDB_Connect_Fail(t *testing.T) {
	cfg := &config.Config{
		PGHost: "127.0.0.1",
		PGPort: "15432",
		PGUser: "wronguser",
		PGPass: "wrongpass",
		PGName: "wrongdb",
	}
	db := NewGormPGDB(cfg)
	err := db.Connect()
	if err == nil {
		t.Fatal("预期连接失败，但实际连接成功")
	} else {
		t.Logf("连接失败，错误信息: %v", err)
	}
}
