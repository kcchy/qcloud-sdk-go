package common

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//"time"
)

// RemovalPolicy.N add index to array item
// RemovalPolicy=["a", "b"] => RemovalPolicy.1="a" RemovalPolicy.2="b"
type FlattenArray []string

type Client struct {
	SecretId    string
	SecretKey   string
	Region      Region
	httpClient  *http.Client
	endpoint    string
	requestPath string
}

// NewClient creates a new instance of ECS client
func (client *Client) Init(secretId, secretKey, endpoint string, region Region) {
	client.SecretId = secretId
	client.SecretKey = secretKey
	client.Region = region
	client.httpClient = &http.Client{}
	client.endpoint = endpoint
	client.requestPath = RequestPath

}

// SetEndpoint sets custom endpoint
func (client *Client) SetSecretId(secretId string) {
	client.SecretId = secretId
}

// SetEndpoint sets custom version
func (client *Client) SetSecretKey(secretKey string) {
	client.SecretKey = secretKey
}

func (client *Client) SetRegion(region Region) {
	client.Region = region
}

func (client *Client) Invoke(action string, args interface{}, response interface{}) error {
	request := Request{}
	request.init(client.Region, action, client.SecretId)

	params := ConvertToparamValues(request, true)

	//设置查询数据

	SetParamValues(args, &params, false)

	sign, _ := sign(RequestMethod, client.endpoint, params, client.SecretKey)

	params.Add("Signature", sign)

	urlStr := "https://" + client.endpoint + client.requestPath

	//Encode 参数 Action=CreateLoadBalancer&DomainPrefix=test&
	var requestBody io.Reader
	requestBody = bytes.NewBufferString(params.Encode())

	r, err := http.NewRequest(RequestMethod, urlStr, requestBody)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	rsp, err := client.httpClient.Do(r)

	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	retData, err := ioutil.ReadAll(rsp.Body)
	if err != err {
		panic(err)
	}

	err = json.Unmarshal([]byte(retData), response)

	jsonOut, _ := json.MarshalIndent(response, "", "  ")
	b2 := append(jsonOut, '\n')
	os.Stdout.Write(b2)

	if err != nil {
		return err
	}

	return nil
}
