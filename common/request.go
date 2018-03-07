package common

import (
	"time"
)

const (
	SignatureMethod = "HmacSHA256"
	RequestPath     = "/v2/index.php"
	RequestMethod   = "POST"
)

type Request struct {
	Action          string
	Region          Region
	Timestamp       int64
	Nonce           int
	SecretId        string
	SignatureMethod string
}

func (request *Request) init(region Region, action string, secretId string) {
	request.Action = action
	request.Region = region
	request.Timestamp = time.Now().Unix()
	request.Nonce = CreateRandomString()
	request.SecretId = secretId
	//request.SignatureMethod = SignatureMethod
	request.SignatureMethod = ""
}
