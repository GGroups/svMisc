package cosurl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const (
	INPUTE_RROR = "#input format Error#"
)

type IObjUrl interface {
	GetObjUrlItems(inputs []UrlObj, channel string) []UrlObj
}

type UrlObj struct {
	Key string `json:"key"`
	Uri string `json:"uri"`
	Url string `json:"url"`
}

func (s UrlObj) GetObjUrlItems(inputs []UrlObj, channel string) []UrlObj {
	var secretId, secretKey, bucket, region string
	if channel == "gpwm" {
		secretId = "AKIDgBPN4UcT3MO9esoTryoNfkY5LL4XmlTu"
		secretKey = "B5i48vRgeKjriIcfQu73h49xUjiPr2wZ"
		bucket = "mj100-1308576853"
		region = "ap-shanghai"
	}

	u, _ := url.Parse(`https://` + bucket + `.cos.` + region + `.myqcloud.com`)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
		},
	})
	ctx := context.Background()

	opt := &cos.PresignedURLOptions{
		// http 请求参数，传入的请求参数需与实际请求相同，能够防止用户篡改此HTTP请求的参数
		Query: &url.Values{},
		// http 请求头部，传入的请求头部需包含在实际请求中，能够防止用户篡改签入此处的HTTP请求头部
		Header: &http.Header{},
	}
	// 添加请求参数
	//opt.Query.Add("x-cos-security-token", "<token>")
	opt.Header.Add("Host", "192.168.0.58")

	for idx, ur := range inputs {
		// e, _ := c.Object.IsExist(ctx, ur.Uri)
		// fmt.Printf("#exist: %v\n", e)
		presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, ur.Uri, secretId, secretKey, time.Hour, opt, true)
		if err != nil {
			panic(err)
		}
		inputs[idx].Url = presignedURL.String()
	}
	return inputs
}

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}
