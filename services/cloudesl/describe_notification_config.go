package cloudesl

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

// DescribeNotificationConfig invokes the cloudesl.DescribeNotificationConfig API synchronously
func (client *Client) DescribeNotificationConfig(request *DescribeNotificationConfigRequest) (response *DescribeNotificationConfigResponse, err error) {
	response = CreateDescribeNotificationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNotificationConfigWithChan invokes the cloudesl.DescribeNotificationConfig API asynchronously
func (client *Client) DescribeNotificationConfigWithChan(request *DescribeNotificationConfigRequest) (<-chan *DescribeNotificationConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeNotificationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNotificationConfig(request)
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

// DescribeNotificationConfigWithCallback invokes the cloudesl.DescribeNotificationConfig API asynchronously
func (client *Client) DescribeNotificationConfigWithCallback(request *DescribeNotificationConfigRequest, callback func(response *DescribeNotificationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNotificationConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeNotificationConfig(request)
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

// DescribeNotificationConfigRequest is the request struct for api DescribeNotificationConfig
type DescribeNotificationConfigRequest struct {
	*requests.RpcRequest
}

// DescribeNotificationConfigResponse is the response struct for api DescribeNotificationConfig
type DescribeNotificationConfigResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	Topic          string `json:"Topic" xml:"Topic"`
	GroupId        string `json:"GroupId" xml:"GroupId"`
	Endpoint       string `json:"Endpoint" xml:"Endpoint"`
	Tag            string `json:"Tag" xml:"Tag"`
	Enable         bool   `json:"Enable" xml:"Enable"`
}

// CreateDescribeNotificationConfigRequest creates a request to invoke DescribeNotificationConfig API
func CreateDescribeNotificationConfigRequest() (request *DescribeNotificationConfigRequest) {
	request = &DescribeNotificationConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeNotificationConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeNotificationConfigResponse creates a response to parse from DescribeNotificationConfig response
func CreateDescribeNotificationConfigResponse() (response *DescribeNotificationConfigResponse) {
	response = &DescribeNotificationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
