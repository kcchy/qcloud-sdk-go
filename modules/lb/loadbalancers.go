package lb

import (
	"github.com/kcchy/qcloudapi-sdk-go/common"
)

type CommonResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

/*
#
# Load Balancer
#
*/

type CreateLoadBalancerArgs struct {
	LoadBalancerType int
	Forward          int
	LoadBalancerName  string
	DomainPrefix     string
	VpcId            string
	SubnetId         string
	ProjectId        int
	Number           int
}

type CreateLoadBalancerResponse struct {
	common.CommonResponse
	DealIds           []string          `json:"dealIds"`
	UnLoadBalancerIds UnLoadBalancerIds `json:"unLoadBalancerIds"`
}

type UnLoadBalancerIds map[string][]string

func (client *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (response *CreateLoadBalancerResponse, err error) {
	response = &CreateLoadBalancerResponse{}

	err = client.Invoke("CreateLoadBalancer", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// query tencent cloud load balancer by load balancer name

type DescribeLoadBalancerArgs struct {
	LoadBalancerIds  common.FlattenArray
	LoadBalancerName string
}

type DescribeLoadBalancerResponse struct {
	CommonResponse
	TotalCount      int `json:"totalCount"`
	LoadBalancerSet []struct {
		LoadBalancerId   string   `json:"loadBalancerId"`
		UnLoadBalancerId string   `json:"unLoadBalancerId"`
		LoadBalancerName string   `json:"loadBalancerName"`
		LoadBalancerType int    `json:"loadBalancerType"`
		Domain           string   `json:"domain"`
		LoadBalancerVips []string `json:"loadBalancerVips"`
		Status           int      `json:"status"`
		CreateTime       string   `json:"createTime"`
		StatusTime       string   `json:"statusTime"`
		ProjectId        int      `json:"projectId"`
		VpcId            int      `json:"vpcId"`
		SubnetId         int      `json:"subnetId"`
	} `json:"loadBalancerSet"`
}

func (client *Client) DescribeLoadBalancer(args *DescribeLoadBalancerArgs) (response *DescribeLoadBalancerResponse, err error) {
	response = &DescribeLoadBalancerResponse{}

	err = client.Invoke("DescribeLoadBalancers", args, response)

	if err != nil {
		return nil, err
	}
	return response, err
}

// delete tencent cloud load balancer by load balancer id

type DeleteLoadBalancerArgs struct {
	LoadBalancerIds common.FlattenArray
}

type DeleteLoadBalancerResponse struct {
	CommonResponse
	RequestId int `json:"requestId"`
}

func (client *Client) DeleteLoadBalancer(args *DeleteLoadBalancerArgs) (response *DeleteLoadBalancerResponse, err error) {
	response = &DeleteLoadBalancerResponse{}

	err = client.Invoke("DeleteLoadBalancers", args, response)

	if err != nil {
		return nil, err
	}
	return response, err
}

// update tencent cloud load balancer, the api only can update load balanacer name and domian prefix

type ModifyLoadBalancerAttributesArgs struct {
	LoadBalancerId   string
	LoadBalancerName string
	DomainPrefix     string
}

type ModifyLoadBalancerAttributesResponse struct {
	CommonResponse
	RequestId int `json:"requestId"`
}

func (client *Client) ModifyLoadBalancerAttributes(args *ModifyLoadBalancerAttributesArgs) (response *ModifyLoadBalancerAttributesResponse, err error) {
	response = &ModifyLoadBalancerAttributesResponse{}

	err = client.Invoke("ModifyLoadBalancerAttributes", args, response)

	if err != nil {
		return nil, err
	}

	return response, err
}
