package cdn

import (
	"github.com/kcchy/qcloudapi-sdk-go/common"
)

type Client struct {
	common.Client
}

const (
	CDNDefaultEndpoint = "cdn.api.qcloud.com"
)

func CdnClient(secretId, secretKey string) *Client {

	client := &Client{}
	client.Init(secretId, secretKey, CDNDefaultEndpoint, "")
	return client
}
