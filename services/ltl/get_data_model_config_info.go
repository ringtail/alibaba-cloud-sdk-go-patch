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

// GetDataModelConfigInfo invokes the ltl.GetDataModelConfigInfo API synchronously
func (client *Client) GetDataModelConfigInfo(request *GetDataModelConfigInfoRequest) (response *GetDataModelConfigInfoResponse, err error) {
	response = CreateGetDataModelConfigInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataModelConfigInfoWithChan invokes the ltl.GetDataModelConfigInfo API asynchronously
func (client *Client) GetDataModelConfigInfoWithChan(request *GetDataModelConfigInfoRequest) (<-chan *GetDataModelConfigInfoResponse, <-chan error) {
	responseChan := make(chan *GetDataModelConfigInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataModelConfigInfo(request)
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

// GetDataModelConfigInfoWithCallback invokes the ltl.GetDataModelConfigInfo API asynchronously
func (client *Client) GetDataModelConfigInfoWithCallback(request *GetDataModelConfigInfoRequest, callback func(response *GetDataModelConfigInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataModelConfigInfoResponse
		var err error
		defer close(result)
		response, err = client.GetDataModelConfigInfo(request)
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

// GetDataModelConfigInfoRequest is the request struct for api GetDataModelConfigInfo
type GetDataModelConfigInfoRequest struct {
	*requests.RpcRequest
	ApiVersion    string `position:"Query" name:"ApiVersion"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	DataModelCode string `position:"Query" name:"DataModelCode"`
}

// GetDataModelConfigInfoResponse is the response struct for api GetDataModelConfigInfo
type GetDataModelConfigInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetDataModelConfigInfoRequest creates a request to invoke GetDataModelConfigInfo API
func CreateGetDataModelConfigInfoRequest() (request *GetDataModelConfigInfoRequest) {
	request = &GetDataModelConfigInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "GetDataModelConfigInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDataModelConfigInfoResponse creates a response to parse from GetDataModelConfigInfo response
func CreateGetDataModelConfigInfoResponse() (response *GetDataModelConfigInfoResponse) {
	response = &GetDataModelConfigInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
