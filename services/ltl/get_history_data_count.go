package ltl

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

// GetHistoryDataCount invokes the ltl.GetHistoryDataCount API synchronously
func (client *Client) GetHistoryDataCount(request *GetHistoryDataCountRequest) (response *GetHistoryDataCountResponse, err error) {
	response = CreateGetHistoryDataCountResponse()
	err = client.DoAction(request, response)
	return
}

// GetHistoryDataCountWithChan invokes the ltl.GetHistoryDataCount API asynchronously
func (client *Client) GetHistoryDataCountWithChan(request *GetHistoryDataCountRequest) (<-chan *GetHistoryDataCountResponse, <-chan error) {
	responseChan := make(chan *GetHistoryDataCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetHistoryDataCount(request)
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

// GetHistoryDataCountWithCallback invokes the ltl.GetHistoryDataCount API asynchronously
func (client *Client) GetHistoryDataCountWithCallback(request *GetHistoryDataCountRequest, callback func(response *GetHistoryDataCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetHistoryDataCountResponse
		var err error
		defer close(result)
		response, err = client.GetHistoryDataCount(request)
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

// GetHistoryDataCountRequest is the request struct for api GetHistoryDataCount
type GetHistoryDataCountRequest struct {
	*requests.RpcRequest
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	ApiVersion string           `position:"Query" name:"ApiVersion"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	ProductKey string           `position:"Query" name:"ProductKey"`
	Key        string           `position:"Query" name:"Key"`
}

// GetHistoryDataCountResponse is the response struct for api GetHistoryDataCount
type GetHistoryDataCountResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetHistoryDataCountRequest creates a request to invoke GetHistoryDataCount API
func CreateGetHistoryDataCountRequest() (request *GetHistoryDataCountRequest) {
	request = &GetHistoryDataCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "GetHistoryDataCount", "", "")
	request.Method = requests.POST
	return
}

// CreateGetHistoryDataCountResponse creates a response to parse from GetHistoryDataCount response
func CreateGetHistoryDataCountResponse() (response *GetHistoryDataCountResponse) {
	response = &GetHistoryDataCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
