package storage

import (
	"classmate_reunion/config"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Cos 腾讯云对象存储服务
type Cos struct{}

func (c *Cos) id() string {
	return config.Cos().Key("secret").String()
}

func (c *Cos) key() string {
	return config.Cos().Key("key").String()
}

// Upload 上传文件到腾讯云对象存储
func (c *Cos) Upload(key string, filePath string) {
	u, _ := url.Parse("https://examplebucket-1250000000.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.id(),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			SecretKey: c.key(), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	_, _, err := client.Object.Upload(
		context.Background(), key, filePath, &cos.MultiUploadOptions{
			OptIni:          nil,
			PartSize:        1,
			ThreadPoolSize:  2,
			CheckPoint:      true,
			DisableChecksum: false,
		},
	)
	//
	if err != nil {
		panic(err)
	}

	err = os.Remove(filePath)
	if err != nil {
		log.Println("删除文件失败")
		return
	}

}
