package lb

import (
	"github.com/kcchy/qcloudapi-sdk-go/common"
)

type Client struct {
	common.Client
}

const (
	LBDefaultEndpoint = "lb.api.qcloud.com"
)

func LbClient(secretId, secretKey string, region common.Region) *Client {

	client := &Client{}
	client.Init(secretId, secretKey, LBDefaultEndpoint, region)
	return client
}
