package monitor

import (
	"github.com/kcchy/qcloud-sdk-go/common"
)

type DescribeBaseMetricsArgs struct {
	Namespace  string
	MetricName string
}

type DescribeBaseMetricsResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	MetricSet []struct {
		Dimensions  string `json:"dimensions"`
		MetricCname string `json:"metricCname"`
		MetricName  string `json:"metricName"`
		Namespace   string `json:"namespace"`
		Period      []int  `json:"period"`
		Unit        string `json:"unit"`
		UnitCname   string `json:"unitCname"`
	} `json:"metricSet"`
}

func (client *Client) DescribeBaseMetrics(args *DescribeBaseMetricsArgs) (response *DescribeBaseMetricsResponse, err error) {
	response = &DescribeBaseMetricsResponse{}

	err = client.Invoke("DescribeBaseMetrics", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

type GetMonitorDataArgs struct {
	Namespace  string
	MetricName string
	Period     int
	StartTime  string
	EndTime    string
	Dimensions common.FlattenMultiDimensionalArray
}

type GetMonitorDataResponse struct {
	common.CommonResponse
	StartTime  string    `json:"startTime"`
	EndTime    string    `json:"endTime"`
	MetricName string    `json:"metricName"`
	Period     int       `json:"period"`
	DataPoints []float64 `json:"dataPoints"`
}

func (client *Client) GetMonitorData(args *GetMonitorDataArgs) (response *GetMonitorDataResponse, err error) {
	response = &GetMonitorDataResponse{}

	err = client.Invoke("GetMonitorData", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}
