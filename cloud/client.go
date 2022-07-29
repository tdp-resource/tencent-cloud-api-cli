package cloud

import (
	"context"
	"encoding/json"

	tc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	th "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	tp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Client struct {
	tc.Client
}

type TakeRequest struct {
	*th.BaseRequest
}

type TakeResponse struct {
	*th.BaseResponse
	Response interface{} `json:"Response"`
}

func NewClient(region, secretId, secretKey string) *Client {

	profile := tp.NewClientProfile()
	credential := tc.NewCredential(secretId, secretKey)

	client := &Client{}
	client.Init(region).WithCredential(credential).WithProfile(profile)

	return client

}

func (c *Client) Take(service, version, action, payload string) (*TakeResponse, error) {

	request := &TakeRequest{
		BaseRequest: &th.BaseRequest{},
	}
	json.Unmarshal([]byte(payload), &request)

	request.Init().WithApiInfo(service, version, action)
	request.SetContext(context.Background())

	response := &TakeResponse{
		BaseResponse: &th.BaseResponse{},
	}

	err := c.Send(request, response)

	return response, err

}
