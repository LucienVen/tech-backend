package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 定义 RedisClientMock 接口，模拟 db.RedisClient 需要的方法
// 便于 redisStore 依赖注入

type RedisClientMock struct {
	store map[string]string
}

func NewRedisClientMock() *RedisClientMock {
	return &RedisClientMock{store: make(map[string]string)}
}

func (m *RedisClientMock) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	m.store[key] = value.(string)
	return nil
}
func (m *RedisClientMock) Get(ctx context.Context, key string) (string, error) {
	val, ok := m.store[key]
	if !ok {
		return "", nil
	}
	return val, nil
}
func (m *RedisClientMock) Del(ctx context.Context, keys ...string) (int64, error) {
	for _, k := range keys {
		delete(m.store, k)
	}
	return int64(len(keys)), nil
}

// 让 redisStore 支持注入 mock
// 这里直接将 redis 字段类型改为 interface，便于测试

type redisStoreForTest struct {
	redis interface {
		Set(context.Context, string, interface{}, time.Duration) error
		Get(context.Context, string) (string, error)
		Del(context.Context, ...string) (int64, error)
	}
}

func (s *redisStoreForTest) Set(id, value string) error {
	return s.redis.Set(context.Background(), id, value, captchaExpire)
}
func (s *redisStoreForTest) Get(id string, clear bool) string {
	val, _ := s.redis.Get(context.Background(), id)
	if clear {
		s.redis.Del(context.Background(), id)
	}
	return val
}
func (s *redisStoreForTest) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}

func TestRedisStore_SetGetVerify(t *testing.T) {
	mock := NewRedisClientMock()
	store := &redisStoreForTest{redis: mock}

	err := store.Set("id1", "code123")
	assert.NoError(t, err)
	val := store.Get("id1", false)
	assert.Equal(t, "code123", val)
	assert.True(t, store.Verify("id1", "code123", false))
	assert.False(t, store.Verify("id1", "wrong", false))
	// 测试 clear
	assert.True(t, store.Verify("id1", "code123", true))
	assert.False(t, store.Verify("id1", "code123", false))
}

func TestCaptchaService_Verify(t *testing.T) {
	mock := NewRedisClientMock()
	store := &redisStoreForTest{redis: mock}
	// 先存入验证码
	err := store.Set("id2", "abcde")
	assert.NoError(t, err)
	// 正确校验
	assert.True(t, store.Verify("id2", "abcde", false))
	// 错误校验
	assert.False(t, store.Verify("id2", "wrong", false))
	// clear 后再次校验
	assert.True(t, store.Verify("id2", "abcde", true))
	assert.False(t, store.Verify("id2", "abcde", false))
}

// mock base64Captcha.Store 用于 Generate 测试

type mockStore struct {
	setCalled bool
}

func (m *mockStore) Set(id, value string) error {
	m.setCalled = true
	return nil
}
func (m *mockStore) Get(id string, clear bool) string          { return "" }
func (m *mockStore) Verify(id, answer string, clear bool) bool { return true }

func TestCaptchaService_Generate(t *testing.T) {
	mock := &mockStore{}
	svc := &CaptchaService{store: mock}
	id, b64s, err := svc.Generate()
	// 增加打印，显示执行细节
	println("id:", id)
	println("b64s:", b64s)
	if err != nil {
		println("err:", err.Error())
	} else {
		println("err: <nil>")
	}
	println("mockStore.Set called:", mock.setCalled)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.NotNil(t, b64s)
}
