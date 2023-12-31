package live

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

// UpdateLiveDelayConfig invokes the live.UpdateLiveDelayConfig API synchronously
func (client *Client) UpdateLiveDelayConfig(request *UpdateLiveDelayConfigRequest) (response *UpdateLiveDelayConfigResponse, err error) {
	response = CreateUpdateLiveDelayConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveDelayConfigWithChan invokes the live.UpdateLiveDelayConfig API asynchronously
func (client *Client) UpdateLiveDelayConfigWithChan(request *UpdateLiveDelayConfigRequest) (<-chan *UpdateLiveDelayConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveDelayConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveDelayConfig(request)
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

// UpdateLiveDelayConfigWithCallback invokes the live.UpdateLiveDelayConfig API asynchronously
func (client *Client) UpdateLiveDelayConfigWithCallback(request *UpdateLiveDelayConfigRequest, callback func(response *UpdateLiveDelayConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveDelayConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveDelayConfig(request)
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

// UpdateLiveDelayConfigRequest is the request struct for api UpdateLiveDelayConfig
type UpdateLiveDelayConfigRequest struct {
	*requests.RpcRequest
	DelayTime       requests.Integer `position:"Query" name:"DelayTime"`
	Stream          string           `position:"Query" name:"Stream"`
	App             string           `position:"Query" name:"App"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	TaskTriggerMode string           `position:"Query" name:"TaskTriggerMode"`
	Domain          string           `position:"Query" name:"Domain"`
}

// UpdateLiveDelayConfigResponse is the response struct for api UpdateLiveDelayConfig
type UpdateLiveDelayConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveDelayConfigRequest creates a request to invoke UpdateLiveDelayConfig API
func CreateUpdateLiveDelayConfigRequest() (request *UpdateLiveDelayConfigRequest) {
	request = &UpdateLiveDelayConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLiveDelayConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLiveDelayConfigResponse creates a response to parse from UpdateLiveDelayConfig response
func CreateUpdateLiveDelayConfigResponse() (response *UpdateLiveDelayConfigResponse) {
	response = &UpdateLiveDelayConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
