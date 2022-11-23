package utils

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
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
