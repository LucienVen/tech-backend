package db

import (
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
	// 加载配置
	config.Load(true, filepath.Join(projectRoot, "cmd"))
	cfg := config.GetConfig()
	t.Logf("PG配置: host=%s port=%s user=%s db=%s", cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGName)
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
