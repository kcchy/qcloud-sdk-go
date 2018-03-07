package lb

import "github.com/kcchy/qcloudapi-sdk-go/common"

const (
	// type Tencent cloud credential
	TestSecretId  = ""
	TestSecretKey = ""
	TesRegion     = common.HongKong
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = LbClient(TestSecretId, TestSecretKey, TesRegion)
	}
	return testClient
}
