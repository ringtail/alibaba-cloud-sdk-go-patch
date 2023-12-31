package dt_oc_info

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetOcIpTrademark invokes the dt_oc_info.GetOcIpTrademark API synchronously
func (client *Client) GetOcIpTrademark(request *GetOcIpTrademarkRequest) (response *GetOcIpTrademarkResponse, err error) {
	response = CreateGetOcIpTrademarkResponse()
	err = client.DoAction(request, response)
	return
}

// GetOcIpTrademarkWithChan invokes the dt_oc_info.GetOcIpTrademark API asynchronously
func (client *Client) GetOcIpTrademarkWithChan(request *GetOcIpTrademarkRequest) (<-chan *GetOcIpTrademarkResponse, <-chan error) {
	responseChan := make(chan *GetOcIpTrademarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOcIpTrademark(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetOcIpTrademarkWithCallback invokes the dt_oc_info.GetOcIpTrademark API asynchronously
func (client *Client) GetOcIpTrademarkWithCallback(request *GetOcIpTrademarkRequest, callback func(response *GetOcIpTrademarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOcIpTrademarkResponse
		var err error
		defer close(result)
		response, err = client.GetOcIpTrademark(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetOcIpTrademarkRequest is the request struct for api GetOcIpTrademark
type GetOcIpTrademarkRequest struct {
	*requests.RpcRequest
	PageNo    requests.Integer `position:"Body" name:"PageNo"`
	PageSize  requests.Integer `position:"Body" name:"PageSize"`
	SearchKey string           `position:"Body" name:"SearchKey"`
}

// GetOcIpTrademarkResponse is the response struct for api GetOcIpTrademark
type GetOcIpTrademarkResponse struct {
	*responses.BaseResponse
	Code      string     `json:"Code" xml:"Code"`
	Success   bool       `json:"Success" xml:"Success"`
	Message   string     `json:"Message" xml:"Message"`
	TotalNum  int        `json:"TotalNum" xml:"TotalNum"`
	PageIndex int        `json:"PageIndex" xml:"PageIndex"`
	PageNum   int        `json:"PageNum" xml:"PageNum"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetOcIpTrademarkRequest creates a request to invoke GetOcIpTrademark API
func CreateGetOcIpTrademarkRequest() (request *GetOcIpTrademarkRequest) {
	request = &GetOcIpTrademarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dt-oc-info", "2022-08-29", "GetOcIpTrademark", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOcIpTrademarkResponse creates a response to parse from GetOcIpTrademark response
func CreateGetOcIpTrademarkResponse() (response *GetOcIpTrademarkResponse) {
	response = &GetOcIpTrademarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
