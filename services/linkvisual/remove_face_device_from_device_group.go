package linkvisual

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

// RemoveFaceDeviceFromDeviceGroup invokes the linkvisual.RemoveFaceDeviceFromDeviceGroup API synchronously
func (client *Client) RemoveFaceDeviceFromDeviceGroup(request *RemoveFaceDeviceFromDeviceGroupRequest) (response *RemoveFaceDeviceFromDeviceGroupResponse, err error) {
	response = CreateRemoveFaceDeviceFromDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveFaceDeviceFromDeviceGroupWithChan invokes the linkvisual.RemoveFaceDeviceFromDeviceGroup API asynchronously
func (client *Client) RemoveFaceDeviceFromDeviceGroupWithChan(request *RemoveFaceDeviceFromDeviceGroupRequest) (<-chan *RemoveFaceDeviceFromDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *RemoveFaceDeviceFromDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveFaceDeviceFromDeviceGroup(request)
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

// RemoveFaceDeviceFromDeviceGroupWithCallback invokes the linkvisual.RemoveFaceDeviceFromDeviceGroup API asynchronously
func (client *Client) RemoveFaceDeviceFromDeviceGroupWithCallback(request *RemoveFaceDeviceFromDeviceGroupRequest, callback func(response *RemoveFaceDeviceFromDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveFaceDeviceFromDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.RemoveFaceDeviceFromDeviceGroup(request)
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

// RemoveFaceDeviceFromDeviceGroupRequest is the request struct for api RemoveFaceDeviceFromDeviceGroup
type RemoveFaceDeviceFromDeviceGroupRequest struct {
	*requests.RpcRequest
	IsolationId   string `position:"Query" name:"IsolationId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DeviceGroupId string `position:"Query" name:"DeviceGroupId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
}

// RemoveFaceDeviceFromDeviceGroupResponse is the response struct for api RemoveFaceDeviceFromDeviceGroup
type RemoveFaceDeviceFromDeviceGroupResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateRemoveFaceDeviceFromDeviceGroupRequest creates a request to invoke RemoveFaceDeviceFromDeviceGroup API
func CreateRemoveFaceDeviceFromDeviceGroupRequest() (request *RemoveFaceDeviceFromDeviceGroupRequest) {
	request = &RemoveFaceDeviceFromDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "RemoveFaceDeviceFromDeviceGroup", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveFaceDeviceFromDeviceGroupResponse creates a response to parse from RemoveFaceDeviceFromDeviceGroup response
func CreateRemoveFaceDeviceFromDeviceGroupResponse() (response *RemoveFaceDeviceFromDeviceGroupResponse) {
	response = &RemoveFaceDeviceFromDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
