package cdn

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

// ListUserCustomLogConfig invokes the cdn.ListUserCustomLogConfig API synchronously
func (client *Client) ListUserCustomLogConfig(request *ListUserCustomLogConfigRequest) (response *ListUserCustomLogConfigResponse, err error) {
	response = CreateListUserCustomLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserCustomLogConfigWithChan invokes the cdn.ListUserCustomLogConfig API asynchronously
func (client *Client) ListUserCustomLogConfigWithChan(request *ListUserCustomLogConfigRequest) (<-chan *ListUserCustomLogConfigResponse, <-chan error) {
	responseChan := make(chan *ListUserCustomLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserCustomLogConfig(request)
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

// ListUserCustomLogConfigWithCallback invokes the cdn.ListUserCustomLogConfig API asynchronously
func (client *Client) ListUserCustomLogConfigWithCallback(request *ListUserCustomLogConfigRequest, callback func(response *ListUserCustomLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserCustomLogConfigResponse
		var err error
		defer close(result)
		response, err = client.ListUserCustomLogConfig(request)
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

// ListUserCustomLogConfigRequest is the request struct for api ListUserCustomLogConfig
type ListUserCustomLogConfigRequest struct {
	*requests.RpcRequest
}

// ListUserCustomLogConfigResponse is the response struct for api ListUserCustomLogConfig
type ListUserCustomLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	ConfigIds ConfigIds `json:"ConfigIds" xml:"ConfigIds"`
}

// CreateListUserCustomLogConfigRequest creates a request to invoke ListUserCustomLogConfig API
func CreateListUserCustomLogConfigRequest() (request *ListUserCustomLogConfigRequest) {
	request = &ListUserCustomLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "ListUserCustomLogConfig", "", "")
	request.Method = requests.GET
	return
}

// CreateListUserCustomLogConfigResponse creates a response to parse from ListUserCustomLogConfig response
func CreateListUserCustomLogConfigResponse() (response *ListUserCustomLogConfigResponse) {
	response = &ListUserCustomLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
