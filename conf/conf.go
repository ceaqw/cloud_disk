package conf

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

// 开发环境
const BUILD = "dev"

// 生产环境
// const BUILD = "prod"

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("service_" + BUILD)
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/conf")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

type AppCfg struct {
	Port string //App 启动端口
	Host string //App
}

type RedisCfg struct {
	Addrs        []string
	Password     string
	ClusterMode  bool
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func GetAppCfg() (appCfg AppCfg) {
	appCfg.Port = viper.GetString("service.port")
	appCfg.Host = viper.GetString("service.host")
	return appCfg
}

func GetRedisCfg() (redisCfg RedisCfg) {
	redisCfg.Addrs = viper.GetStringSlice("models.redis.addrs")
	redisCfg.Password = viper.GetString("models.redis.password")
	redisCfg.ClusterMode = viper.GetBool("models.redis.cluster_mode")
	redisCfg.DialTimeout = viper.GetDuration("models.redis.dial_timeout")
	redisCfg.ReadTimeout = viper.GetDuration("models.redis.read_timeout")
	redisCfg.WriteTimeout = viper.GetDuration("models.redis.write_timeout")
	return redisCfg
}
