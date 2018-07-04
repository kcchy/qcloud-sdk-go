package monitor

import (
	"github.com/kcchy/qcloud-sdk-go/common"
)

type Client struct {
	common.Client
}

const (
	MonitorDefaultEndpoint = "monitor.api.qcloud.com"
)

func MonitorClient(secretId, secretKey string, region common.Region) *Client {

	client := &Client{}
	client.Init(secretId, secretKey, MonitorDefaultEndpoint, region)
	return client
}
