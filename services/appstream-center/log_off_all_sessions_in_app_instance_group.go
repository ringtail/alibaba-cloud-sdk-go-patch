package appstream_center

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

// LogOffAllSessionsInAppInstanceGroup invokes the appstream_center.LogOffAllSessionsInAppInstanceGroup API synchronously
func (client *Client) LogOffAllSessionsInAppInstanceGroup(request *LogOffAllSessionsInAppInstanceGroupRequest) (response *LogOffAllSessionsInAppInstanceGroupResponse, err error) {
	response = CreateLogOffAllSessionsInAppInstanceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// LogOffAllSessionsInAppInstanceGroupWithChan invokes the appstream_center.LogOffAllSessionsInAppInstanceGroup API asynchronously
func (client *Client) LogOffAllSessionsInAppInstanceGroupWithChan(request *LogOffAllSessionsInAppInstanceGroupRequest) (<-chan *LogOffAllSessionsInAppInstanceGroupResponse, <-chan error) {
	responseChan := make(chan *LogOffAllSessionsInAppInstanceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LogOffAllSessionsInAppInstanceGroup(request)
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

// LogOffAllSessionsInAppInstanceGroupWithCallback invokes the appstream_center.LogOffAllSessionsInAppInstanceGroup API asynchronously
func (client *Client) LogOffAllSessionsInAppInstanceGroupWithCallback(request *LogOffAllSessionsInAppInstanceGroupRequest, callback func(response *LogOffAllSessionsInAppInstanceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LogOffAllSessionsInAppInstanceGroupResponse
		var err error
		defer close(result)
		response, err = client.LogOffAllSessionsInAppInstanceGroup(request)
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

// LogOffAllSessionsInAppInstanceGroupRequest is the request struct for api LogOffAllSessionsInAppInstanceGroup
type LogOffAllSessionsInAppInstanceGroupRequest struct {
	*requests.RpcRequest
	ProductType        string `position:"Body" name:"ProductType"`
	AppInstanceGroupId string `position:"Body" name:"AppInstanceGroupId"`
}

// LogOffAllSessionsInAppInstanceGroupResponse is the response struct for api LogOffAllSessionsInAppInstanceGroup
type LogOffAllSessionsInAppInstanceGroupResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateLogOffAllSessionsInAppInstanceGroupRequest creates a request to invoke LogOffAllSessionsInAppInstanceGroup API
func CreateLogOffAllSessionsInAppInstanceGroupRequest() (request *LogOffAllSessionsInAppInstanceGroupRequest) {
	request = &LogOffAllSessionsInAppInstanceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "LogOffAllSessionsInAppInstanceGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateLogOffAllSessionsInAppInstanceGroupResponse creates a response to parse from LogOffAllSessionsInAppInstanceGroup response
func CreateLogOffAllSessionsInAppInstanceGroupResponse() (response *LogOffAllSessionsInAppInstanceGroupResponse) {
	response = &LogOffAllSessionsInAppInstanceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
