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
	GetObjUrlItems(name ObjUrlRequest) []UrlObj
}

type UrlObj struct {
	Key string `json:"key"`
	Uri string `json:"uri"`
	Url string `json:"url"`
}

func (s UrlObj) GetObjUrlItems(t ObjUrlRequest) []UrlObj {
	var secretId, secretKey, bucket, region string
	if t.MallInfo.ChannelId == "gpwm" {
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

	for idx, ur := range t.UrlObjs {
		presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, ur.Uri, secretId, secretKey, time.Hour, nil)
		if err != nil {
			panic(err)
		}
		t.UrlObjs[idx].Url = presignedURL.String()
	}
	return t.UrlObjs
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
