package cc5g

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

// DeleteWirelessCloudConnectorGroup invokes the cc5g.DeleteWirelessCloudConnectorGroup API synchronously
func (client *Client) DeleteWirelessCloudConnectorGroup(request *DeleteWirelessCloudConnectorGroupRequest) (response *DeleteWirelessCloudConnectorGroupResponse, err error) {
	response = CreateDeleteWirelessCloudConnectorGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteWirelessCloudConnectorGroupWithChan invokes the cc5g.DeleteWirelessCloudConnectorGroup API asynchronously
func (client *Client) DeleteWirelessCloudConnectorGroupWithChan(request *DeleteWirelessCloudConnectorGroupRequest) (<-chan *DeleteWirelessCloudConnectorGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteWirelessCloudConnectorGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteWirelessCloudConnectorGroup(request)
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

// DeleteWirelessCloudConnectorGroupWithCallback invokes the cc5g.DeleteWirelessCloudConnectorGroup API asynchronously
func (client *Client) DeleteWirelessCloudConnectorGroupWithCallback(request *DeleteWirelessCloudConnectorGroupRequest, callback func(response *DeleteWirelessCloudConnectorGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteWirelessCloudConnectorGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteWirelessCloudConnectorGroup(request)
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

// DeleteWirelessCloudConnectorGroupRequest is the request struct for api DeleteWirelessCloudConnectorGroup
type DeleteWirelessCloudConnectorGroupRequest struct {
	*requests.RpcRequest
	WirelessCloudConnectorGroupId string           `position:"Query" name:"WirelessCloudConnectorGroupId"`
	DryRun                        requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
}

// DeleteWirelessCloudConnectorGroupResponse is the response struct for api DeleteWirelessCloudConnectorGroup
type DeleteWirelessCloudConnectorGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteWirelessCloudConnectorGroupRequest creates a request to invoke DeleteWirelessCloudConnectorGroup API
func CreateDeleteWirelessCloudConnectorGroupRequest() (request *DeleteWirelessCloudConnectorGroupRequest) {
	request = &DeleteWirelessCloudConnectorGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "DeleteWirelessCloudConnectorGroup", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteWirelessCloudConnectorGroupResponse creates a response to parse from DeleteWirelessCloudConnectorGroup response
func CreateDeleteWirelessCloudConnectorGroupResponse() (response *DeleteWirelessCloudConnectorGroupResponse) {
	response = &DeleteWirelessCloudConnectorGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
