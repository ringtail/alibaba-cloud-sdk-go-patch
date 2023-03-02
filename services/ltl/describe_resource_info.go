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

// DescribeResourceInfo invokes the ltl.DescribeResourceInfo API synchronously
func (client *Client) DescribeResourceInfo(request *DescribeResourceInfoRequest) (response *DescribeResourceInfoResponse, err error) {
	response = CreateDescribeResourceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceInfoWithChan invokes the ltl.DescribeResourceInfo API asynchronously
func (client *Client) DescribeResourceInfoWithChan(request *DescribeResourceInfoRequest) (<-chan *DescribeResourceInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceInfo(request)
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

// DescribeResourceInfoWithCallback invokes the ltl.DescribeResourceInfo API asynchronously
func (client *Client) DescribeResourceInfoWithCallback(request *DescribeResourceInfoRequest, callback func(response *DescribeResourceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceInfo(request)
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

// DescribeResourceInfoRequest is the request struct for api DescribeResourceInfo
type DescribeResourceInfoRequest struct {
	*requests.RpcRequest
	ApiVersion string `position:"Query" name:"ApiVersion"`
	BizChainId string `position:"Query" name:"BizChainId"`
}

// DescribeResourceInfoResponse is the response struct for api DescribeResourceInfo
type DescribeResourceInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeResourceInfoRequest creates a request to invoke DescribeResourceInfo API
func CreateDescribeResourceInfoRequest() (request *DescribeResourceInfoRequest) {
	request = &DescribeResourceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "DescribeResourceInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeResourceInfoResponse creates a response to parse from DescribeResourceInfo response
func CreateDescribeResourceInfoResponse() (response *DescribeResourceInfoResponse) {
	response = &DescribeResourceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
