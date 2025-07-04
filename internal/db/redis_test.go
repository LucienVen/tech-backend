package db

import (
	"context"
	"testing"
	"time"

	"github.com/LucienVen/tech-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestRedisClient_Common(t *testing.T) {
	cfg := testutil.LoadTestConfig(t)
	if cfg.RedisAddr == "" {
		t.Fatal("REDIS_ADDR 未配置，请在 cmd/.env 中设置 REDIS_ADDR")
	}
	client := NewRedisClient(cfg.RedisAddr, cfg.RedisPassword, 0)
	defer client.Close()
	ctx := context.Background()

	t.Run("Ping", func(t *testing.T) {
		err := client.Ping(ctx)
		if err != nil {
			t.Skipf("Redis连接失败，跳过测试: %v", err)
		}
		assert.NoError(t, err, "Ping应该成功")
	})

	t.Run("Set/Get/Exists/Del", func(t *testing.T) {
		// Set
		err := client.Set(ctx, "hello", "world", time.Minute)
		assert.NoError(t, err, "Set操作应该成功")
		// Get
		val, err := client.Get(ctx, "hello")
		assert.NoError(t, err, "Get操作应该成功")
		assert.Equal(t, "world", val, "Get值应该为world")
		// Exists
		exists, err := client.Exists(ctx, "hello")
		assert.NoError(t, err, "Exists操作应该成功")
		assert.Equal(t, int64(1), exists, "hello键应该存在")
		// Del
		del, err := client.Del(ctx, "hello")
		assert.NoError(t, err, "Del操作应该成功")
		assert.Equal(t, int64(1), del, "应该删除1个键")
		// Exists again
		exists, err = client.Exists(ctx, "hello")
		assert.NoError(t, err, "Exists操作应该成功")
		assert.Equal(t, int64(0), exists, "hello键应该已被删除")
	})

	t.Run("过期测试", func(t *testing.T) {
		err := client.Set(ctx, "expire_key", "expire_val", time.Second)
		assert.NoError(t, err, "Set操作应该成功")
		val, err := client.Get(ctx, "expire_key")
		assert.NoError(t, err, "Get操作应该成功")
		assert.Equal(t, "expire_val", val, "值应该正确")
		time.Sleep(1500 * time.Millisecond)
		_, err = client.Get(ctx, "expire_key")
		assert.Error(t, err, "过期后应该无法获取")
	})

	t.Run("关闭连接", func(t *testing.T) {
		err := client.Close()
		assert.NoError(t, err, "关闭连接应该成功")
	})
}
