package conf

import (
	"fmt"
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

type MainDbCfg struct {
	UserName string
	Password string
	Database string
	Host     string
	Port     string
	Driver   string
}

type JwtCfg struct {
	TimeOut time.Duration
	Issuer  string
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

func GetDbCfgByName(dbName string) (mainDbCfg MainDbCfg) {
	mainDbCfg.UserName = viper.GetString(fmt.Sprintf("models.%s.username", dbName))
	mainDbCfg.Password = viper.GetString(fmt.Sprintf("models.%s.password", dbName))
	mainDbCfg.Database = viper.GetString(fmt.Sprintf("models.%s.database", dbName))
	mainDbCfg.Host = viper.GetString(fmt.Sprintf("models.%s.host", dbName))
	mainDbCfg.Port = viper.GetString(fmt.Sprintf("models.%s.port", dbName))
	mainDbCfg.Driver = viper.GetString(fmt.Sprintf("models.%s.driver", dbName))
	return mainDbCfg
}

func GetJwtCfg() (jwtCfg JwtCfg) {
	jwtCfg.TimeOut = viper.GetDuration("jwt.time_out")
	jwtCfg.Issuer = viper.GetString("jwt.issuer")
	return jwtCfg
}
