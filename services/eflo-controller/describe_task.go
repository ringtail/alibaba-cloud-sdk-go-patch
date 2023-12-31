package eflo_controller

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

// DescribeTask invokes the eflo_controller.DescribeTask API synchronously
func (client *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
	response = CreateDescribeTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTaskWithChan invokes the eflo_controller.DescribeTask API asynchronously
func (client *Client) DescribeTaskWithChan(request *DescribeTaskRequest) (<-chan *DescribeTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTask(request)
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

// DescribeTaskWithCallback invokes the eflo_controller.DescribeTask API asynchronously
func (client *Client) DescribeTaskWithCallback(request *DescribeTaskRequest, callback func(response *DescribeTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeTask(request)
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

// DescribeTaskRequest is the request struct for api DescribeTask
type DescribeTaskRequest struct {
	*requests.RpcRequest
	TaskId string `position:"Body" name:"TaskId"`
}

// DescribeTaskResponse is the response struct for api DescribeTask
type DescribeTaskResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ClusterId   string      `json:"ClusterId" xml:"ClusterId"`
	ClusterName string      `json:"ClusterName" xml:"ClusterName"`
	TaskState   string      `json:"TaskState" xml:"TaskState"`
	TaskType    string      `json:"TaskType" xml:"TaskType"`
	Message     string      `json:"Message" xml:"Message"`
	CreateTime  string      `json:"CreateTime" xml:"CreateTime"`
	UpdateTime  string      `json:"UpdateTime" xml:"UpdateTime"`
	Steps       []StepsItem `json:"Steps" xml:"Steps"`
}

// CreateDescribeTaskRequest creates a request to invoke DescribeTask API
func CreateDescribeTaskRequest() (request *DescribeTaskRequest) {
	request = &DescribeTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo-controller", "2022-12-15", "DescribeTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeTaskResponse creates a response to parse from DescribeTask response
func CreateDescribeTaskResponse() (response *DescribeTaskResponse) {
	response = &DescribeTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
