package rdc

import (
	"encoding/json"
	"golang-project-prototype/config"
	"golang-project-prototype/library/helper"
	"golang-project-prototype/library/util/logger"
	"golang-project-prototype/model"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

const (
	// k expire 添加随机时间[0,randExpire)
	// 防止缓存雪崩
	randExpire = 300
)

func init() {
	pool = newPool(config.RedisConfig.RedisHost)
	Ping()
}

// 创建redis连接池
func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.RedisConfig.RedisMaxIdle,
		IdleTimeout: time.Duration(config.RedisConfig.RedisIdleTimeout) * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			if psd := config.RedisConfig.RedisPassword; psd != "" {
				if _, err := c.Do("AUTH", psd); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, nil
		},
	}
}

func Ping() {
	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		logger.Error("redis:ping fail", err)
		return
	}
}

func SetWithExpire(k string, v interface{}, expire int) error {
	jb, err := json.Marshal(&v)
	if err != nil {
		logger.Error("json编码失败", err)
		return model.ErrJSONEncode
	}

	conn := pool.Get()
	defer conn.Close()

	expire = expire + helper.RandIntn(randExpire)
	_, err = conn.Do("SETEX", k, expire, string(jb))
	if err != nil {
		logger.Error("redis:set fail", err)
		return model.ErrUnknown
	}

	return nil
}

func Get(k string, receiver interface{}) error {
	conn := pool.Get()
	defer conn.Close()

	v, err := redis.String(conn.Do("GET", k))
	if err != nil {
		logger.Error("redis:set fail", err)
		return model.ErrUnknown
	}

	err = json.Unmarshal([]byte(v), receiver)
	if err != nil {
		logger.Error("json解码失败", err)
		return model.ErrJSONDecode
	}

	return nil
}
