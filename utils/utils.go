package utils

import (
	"CouldDisk/moudles/redis_helper"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gopkg.in/gomail.v2"
)

var (
	TencentSecretID  = os.Getenv("TencentSecretID")
	TencentSecretKey = os.Getenv("TencentSecretKey")
	CosBucket        = viper.GetString("cos.cos_bucket")
)

func _getCosClient() *cos.Client {
	u, _ := url.Parse(CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  TencentSecretID,
			SecretKey: TencentSecretKey,
		},
	})
}

func CosUpload(filePath string, ext string, file io.Reader) (string, error) {
	c := _getCosClient()
	_, err := c.Object.Put(
		context.Background(), filePath, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return CosBucket + "/" + filePath, nil
}

func CosDownload(filePath string) (io.Reader, error) {
	c := _getCosClient()
	resp, err := c.Object.Get(context.Background(), filePath, nil)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func CosDelete(filePath string) error {
	c := _getCosClient()
	_, err := c.Object.Delete(context.Background(), filePath)
	if err != nil {
		return err
	}
	return nil
}

//向邮箱里发送验证码
func SendVerifyCode(email string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "1220797826@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "coulddisk重置密码验证码")
	//生成随机code
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	m.SetBody("text/html", code)
	d := gomail.NewDialer("smtp.qq.com", 587, "1220797826@qq.com", "cxgrirwjespkbaaa")
	err := d.DialAndSend(m)
	//验证码写入redis中
	if err == nil {
		var redish redis_helper.RedisHelper
		err = redish.SaveRedisToken(email, code)
		if err != nil {
			return err
		}
	}
	return err
}

func CheckVerifyCode(email, code string) bool {
	//在redis里获取code
	var redish redis_helper.RedisHelper
	redisCode, err := redish.GetTokenFromRedisByName(email)
	if err != nil {
		return false
	}
	return redisCode == code
}
