package authtx

const (
	INPUTE_RROR = "#input format Error#"
)

type IAuthTX interface {
	GetAuthTXItems(name string) []AuthTX
}

type AuthTX struct {
	SecretId        string   `json:"secretId"`
	SecretKey       string   `json:"secretKey"`
	Proxy           string   `json:"proxy"`
	DurationSeconds int      `json:"durationSeconds"`
	Bucket          string   `json:"bucket"` //mj100-1308576853
	Region          string   `json:"region"`
	AllowPrefix     string   `json:"allowPrefix"`  //'_ALLOW_DIR_/*'
	AllowActions    []string `json:"allowActions"` // 密钥的权限列表
}

func (s AuthTX) GetAuthTXItems(t string) []AuthTX {
	// allowActions := []string{"name/cos:PutObject",
	// 	"name/cos:PostObject",
	// 	"name/cos:InitiateMultipartUpload",
	// 	"name/cos:ListMultipartUploads",
	// 	"name/cos:ListParts",
	// 	"name/cos:UploadPart",
	// 	"name/cos:CompleteMultipartUpload",
	// 	"name/cos:AbortMultipartUpload"}
	var secretId, secretKey, bucket, region string
	if t == "gpwm" {
		secretId = "AKIDgBPN4UcT3MO9esoTryoNfkY5LL4XmlTu"
		secretKey = "B5i48vRgeKjriIcfQu73h49xUjiPr2wZ"
		bucket = "mj100-1308576853"
		region = "ap-shanghai"
	}
	ss := []AuthTX{
		{
			SecretId:        secretId,
			SecretKey:       secretKey,
			Proxy:           "",
			DurationSeconds: 1800,
			Bucket:          bucket,
			Region:          region,
			AllowPrefix:     "",
			AllowActions:    []string{"name/cos:GetServic"},
		},
	}
	return ss
}
