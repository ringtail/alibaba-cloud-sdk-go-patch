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

// GetOcIcSimpleCancel invokes the dt_oc_info.GetOcIcSimpleCancel API synchronously
func (client *Client) GetOcIcSimpleCancel(request *GetOcIcSimpleCancelRequest) (response *GetOcIcSimpleCancelResponse, err error) {
	response = CreateGetOcIcSimpleCancelResponse()
	err = client.DoAction(request, response)
	return
}

// GetOcIcSimpleCancelWithChan invokes the dt_oc_info.GetOcIcSimpleCancel API asynchronously
func (client *Client) GetOcIcSimpleCancelWithChan(request *GetOcIcSimpleCancelRequest) (<-chan *GetOcIcSimpleCancelResponse, <-chan error) {
	responseChan := make(chan *GetOcIcSimpleCancelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOcIcSimpleCancel(request)
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

// GetOcIcSimpleCancelWithCallback invokes the dt_oc_info.GetOcIcSimpleCancel API asynchronously
func (client *Client) GetOcIcSimpleCancelWithCallback(request *GetOcIcSimpleCancelRequest, callback func(response *GetOcIcSimpleCancelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOcIcSimpleCancelResponse
		var err error
		defer close(result)
		response, err = client.GetOcIcSimpleCancel(request)
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

// GetOcIcSimpleCancelRequest is the request struct for api GetOcIcSimpleCancel
type GetOcIcSimpleCancelRequest struct {
	*requests.RpcRequest
	PageNo    requests.Integer `position:"Body" name:"PageNo"`
	PageSize  requests.Integer `position:"Body" name:"PageSize"`
	SearchKey string           `position:"Body" name:"SearchKey"`
}

// GetOcIcSimpleCancelResponse is the response struct for api GetOcIcSimpleCancel
type GetOcIcSimpleCancelResponse struct {
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

// CreateGetOcIcSimpleCancelRequest creates a request to invoke GetOcIcSimpleCancel API
func CreateGetOcIcSimpleCancelRequest() (request *GetOcIcSimpleCancelRequest) {
	request = &GetOcIcSimpleCancelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dt-oc-info", "2022-08-29", "GetOcIcSimpleCancel", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOcIcSimpleCancelResponse creates a response to parse from GetOcIcSimpleCancel response
func CreateGetOcIcSimpleCancelResponse() (response *GetOcIcSimpleCancelResponse) {
	response = &GetOcIcSimpleCancelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
