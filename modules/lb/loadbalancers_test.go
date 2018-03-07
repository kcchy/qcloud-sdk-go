package lb

import (
	//"github.com/kcchy/qcloudapi-sdk-go/common"
	"testing"
)

func TestLoadBalancer(t *testing.T) {

	client := NewTestClient()

	// create lb

	creationArgs := &CreateLoadBalancerArgs{
		LoadBalancerType: 3,
		Forward:          0,
		LoadBalancerName:  "test",
		DomainPrefix:     "test",
		VpcId:            "0",
	}

	response, err := client.CreateLoadBalancer(creationArgs)

	if err != nil {
		t.Fatalf("Failed to CreateLoadBalancer: %v", err)
	}

	t.Logf("CreateLoadBalancer result: %v", *response)

	lbId := response.UnLoadBalancerIds[response.DealIds[0]][0]

	// describe lb

	DescribeArgs := &DescribeLoadBalancerArgs{
		LoadBalancerIds: []string{lbId},
	}

	responseDescribe, err := client.DescribeLoadBalancer(DescribeArgs)

	if err != nil {
		t.Errorf("Failed to DescribeLoadBalancer: %v", err)
	}

	t.Logf("DescribeLoadBalancer successfully: %s", responseDescribe)

	// update lb

	UpdateArgs := &ModifyLoadBalancerAttributesArgs{
		LoadBalancerId:   lbId,
		LoadBalancerName: "test1",
		DomainPrefix:     "test1",
	}

	responseUpdate, err := client.ModifyLoadBalancerAttributes(UpdateArgs)

	if err != nil {
		t.Errorf("Failed to ModifyLoadBalancerAttributes: %v", err)
	}

	t.Logf("ModifyLoadBalancerAttributes successfully: %s", responseUpdate)

	// delete lb

	DeleteArgs := &DeleteLoadBalancerArgs{
		LoadBalancerIds: []string{lbId},
	}

	responseDelete, err := client.DeleteLoadBalancer(DeleteArgs)

	if err != nil {
		t.Errorf("Failed to DeleteLoadBalancer: %v", err)
	}

	t.Logf("DeleteLoadBalancer successfully: %s", responseDelete)

}
