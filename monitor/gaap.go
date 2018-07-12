package monitor

import (
	"github.com/kcchy/qcloud-sdk-go/common"
)

type GetMonitorDataArgs struct {
	Namespace  string
	MetricName string
	Period     int
	StartTime  string
	EndTime    string
	Dimensions common.FlattenMultiDimensionalArray
}

type GetMonitorDataReponse struct {
	common.CommonResponse
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	MetricName string `json:"metricName"`
	Period     int    `json:"period"`
	DataPoints []int  `json:"dataPoints"`
}

func (client *Client) GetMonitorData(args *GetMonitorDataArgs) (response *GetMonitorDataReponse, err error) {
	response = &GetMonitorDataReponse{}

	err = client.Invoke("GetMonitorData", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}
