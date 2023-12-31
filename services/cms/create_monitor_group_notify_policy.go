package cms

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

// CreateMonitorGroupNotifyPolicy invokes the cms.CreateMonitorGroupNotifyPolicy API synchronously
func (client *Client) CreateMonitorGroupNotifyPolicy(request *CreateMonitorGroupNotifyPolicyRequest) (response *CreateMonitorGroupNotifyPolicyResponse, err error) {
	response = CreateCreateMonitorGroupNotifyPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMonitorGroupNotifyPolicyWithChan invokes the cms.CreateMonitorGroupNotifyPolicy API asynchronously
func (client *Client) CreateMonitorGroupNotifyPolicyWithChan(request *CreateMonitorGroupNotifyPolicyRequest) (<-chan *CreateMonitorGroupNotifyPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateMonitorGroupNotifyPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMonitorGroupNotifyPolicy(request)
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

// CreateMonitorGroupNotifyPolicyWithCallback invokes the cms.CreateMonitorGroupNotifyPolicy API asynchronously
func (client *Client) CreateMonitorGroupNotifyPolicyWithCallback(request *CreateMonitorGroupNotifyPolicyRequest, callback func(response *CreateMonitorGroupNotifyPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMonitorGroupNotifyPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateMonitorGroupNotifyPolicy(request)
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

// CreateMonitorGroupNotifyPolicyRequest is the request struct for api CreateMonitorGroupNotifyPolicy
type CreateMonitorGroupNotifyPolicyRequest struct {
	*requests.RpcRequest
	PolicyType string           `position:"Query" name:"PolicyType"`
	GroupId    string           `position:"Query" name:"GroupId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
}

// CreateMonitorGroupNotifyPolicyResponse is the response struct for api CreateMonitorGroupNotifyPolicy
type CreateMonitorGroupNotifyPolicyResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    int    `json:"Result" xml:"Result"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateCreateMonitorGroupNotifyPolicyRequest creates a request to invoke CreateMonitorGroupNotifyPolicy API
func CreateCreateMonitorGroupNotifyPolicyRequest() (request *CreateMonitorGroupNotifyPolicyRequest) {
	request = &CreateMonitorGroupNotifyPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateMonitorGroupNotifyPolicy", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMonitorGroupNotifyPolicyResponse creates a response to parse from CreateMonitorGroupNotifyPolicy response
func CreateCreateMonitorGroupNotifyPolicyResponse() (response *CreateMonitorGroupNotifyPolicyResponse) {
	response = &CreateMonitorGroupNotifyPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
