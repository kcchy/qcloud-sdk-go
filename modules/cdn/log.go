package cdn

import (
	"github.com/kcchy/qcloudapi-sdk-go/common"
)

type GetCdnLogListArgs struct {
	Host      string
	StartDate string
	EndDate   string
}

type GetCdnLogListReponse struct {
	common.CommonResponse
	Data struct {
		Now  int `json:"now"`
		List []struct {
			Date string `json:"date"`
			Type int    `json:"type"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"list"`
	} `json:"data"`
}

func (client *Client) GetCdnLogList(args *GetCdnLogListArgs) (response *GetCdnLogListReponse, err error) {
	response = &GetCdnLogListReponse{}

	err = client.Invoke("GetCdnLogList", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}
