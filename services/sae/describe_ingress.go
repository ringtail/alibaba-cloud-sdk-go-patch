package sae

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

// DescribeIngress invokes the sae.DescribeIngress API synchronously
func (client *Client) DescribeIngress(request *DescribeIngressRequest) (response *DescribeIngressResponse, err error) {
	response = CreateDescribeIngressResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIngressWithChan invokes the sae.DescribeIngress API asynchronously
func (client *Client) DescribeIngressWithChan(request *DescribeIngressRequest) (<-chan *DescribeIngressResponse, <-chan error) {
	responseChan := make(chan *DescribeIngressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIngress(request)
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

// DescribeIngressWithCallback invokes the sae.DescribeIngress API asynchronously
func (client *Client) DescribeIngressWithCallback(request *DescribeIngressRequest, callback func(response *DescribeIngressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIngressResponse
		var err error
		defer close(result)
		response, err = client.DescribeIngress(request)
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

// DescribeIngressRequest is the request struct for api DescribeIngress
type DescribeIngressRequest struct {
	*requests.RoaRequest
	IngressId requests.Integer `position:"Query" name:"IngressId"`
}

// DescribeIngressResponse is the response struct for api DescribeIngress
type DescribeIngressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeIngressRequest creates a request to invoke DescribeIngress API
func CreateDescribeIngressRequest() (request *DescribeIngressRequest) {
	request = &DescribeIngressRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeIngress", "/pop/v1/sam/ingress/Ingress", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeIngressResponse creates a response to parse from DescribeIngress response
func CreateDescribeIngressResponse() (response *DescribeIngressResponse) {
	response = &DescribeIngressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
