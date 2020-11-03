package drds

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

// DescribeRdsReadOnly invokes the drds.DescribeRdsReadOnly API synchronously
func (client *Client) DescribeRdsReadOnly(request *DescribeRdsReadOnlyRequest) (response *DescribeRdsReadOnlyResponse, err error) {
	response = CreateDescribeRdsReadOnlyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRdsReadOnlyWithChan invokes the drds.DescribeRdsReadOnly API asynchronously
func (client *Client) DescribeRdsReadOnlyWithChan(request *DescribeRdsReadOnlyRequest) (<-chan *DescribeRdsReadOnlyResponse, <-chan error) {
	responseChan := make(chan *DescribeRdsReadOnlyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRdsReadOnly(request)
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

// DescribeRdsReadOnlyWithCallback invokes the drds.DescribeRdsReadOnly API asynchronously
func (client *Client) DescribeRdsReadOnlyWithCallback(request *DescribeRdsReadOnlyRequest, callback func(response *DescribeRdsReadOnlyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRdsReadOnlyResponse
		var err error
		defer close(result)
		response, err = client.DescribeRdsReadOnly(request)
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

// DescribeRdsReadOnlyRequest is the request struct for api DescribeRdsReadOnly
type DescribeRdsReadOnlyRequest struct {
	*requests.RpcRequest
	RdsInstanceId  string `position:"Query" name:"RdsInstanceId"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbInstType     string `position:"Query" name:"DbInstType"`
}

// DescribeRdsReadOnlyResponse is the response struct for api DescribeRdsReadOnly
type DescribeRdsReadOnlyResponse struct {
	*responses.BaseResponse
	RequestId   string                           `json:"RequestId" xml:"RequestId"`
	Success     bool                             `json:"Success" xml:"Success"`
	DbInstances DbInstancesInDescribeRdsReadOnly `json:"DbInstances" xml:"DbInstances"`
}

// CreateDescribeRdsReadOnlyRequest creates a request to invoke DescribeRdsReadOnly API
func CreateDescribeRdsReadOnlyRequest() (request *DescribeRdsReadOnlyRequest) {
	request = &DescribeRdsReadOnlyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeRdsReadOnly", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRdsReadOnlyResponse creates a response to parse from DescribeRdsReadOnly response
func CreateDescribeRdsReadOnlyResponse() (response *DescribeRdsReadOnlyResponse) {
	response = &DescribeRdsReadOnlyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}