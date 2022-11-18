package redis_helper

import (
	"CouldDisk/conf"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type RedisHelper struct {
}

var (
	// Redis        *redis.Client
	// RedisCluster *redis.ClusterClient
	Redis      interface{}
	ClientMode string
)

func init() {
	redisCfg := conf.GetRedisCfg()
	if redisCfg.ClusterMode == true {
		// RedisCluster = NewRedisClusterHelper(redisCfg)
		Redis = NewRedisClusterHelper(redisCfg)
		ClientMode = "cluster"
	} else {
		Redis = NewRedisHelper(redisCfg)
		ClientMode = "signle"
	}
}

func NewRedisHelper(redisCfg conf.RedisCfg) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addrs[0],
		Password: redisCfg.Password,
		DB:       0,
		// DialTimeout:  redisCfg.DialTimeout,
		// ReadTimeout:  redisCfg.ReadTimeout,
		// WriteTimeout: redisCfg.WriteTimeout,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis connect error: %#v\n", err.Error()))
	}
	return client
}
func NewRedisClusterHelper(redisCfg conf.RedisCfg) *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redisCfg.Addrs,
		Password: redisCfg.Password,
		// DialTimeout:  redisCfg.DialTimeout,
		// ReadTimeout:  redisCfg.ReadTimeout,
		// WriteTimeout: redisCfg.WriteTimeout,
	})
	if _, err := clusterClient.Ping().Result(); err != nil {
		panic(fmt.Sprintf("redis connect error: %#v\n", err.Error()))
	}
	return clusterClient
}

func (h RedisHelper) SaveRedisToken(user string, token string) error {
	expireTime := viper.GetDuration("jwt.time_out")
	var err error
	if ClientMode == "cluster" {
		redisClient := Redis.(*redis.ClusterClient)
		_, err = redisClient.Set(user, token, expireTime*time.Minute).Result()
	} else {
		redisClient := Redis.(*redis.Client)
		_, err = redisClient.Set(user, token, expireTime*time.Minute).Result()
	}
	return err
}

func (h RedisHelper) RemoveRedisToken(user string) error {
	var err error
	if ClientMode == "cluster" {
		redisClient := Redis.(*redis.ClusterClient)
		_, err = redisClient.Del(user).Result()
	} else {
		redisClient := Redis.(*redis.Client)
		_, err = redisClient.Del(user).Result()
	}
	return err
}

func (h RedisHelper) GetTokenFromRedisByName(user string) (string, error) {
	if ClientMode == "cluster" {
		redisClient := Redis.(*redis.ClusterClient)
		val, err := redisClient.Get(user).Result()
		return val, err
	} else {
		redisClient := Redis.(*redis.Client)
		val, err := redisClient.Get(user).Result()
		return val, err
	}
}
