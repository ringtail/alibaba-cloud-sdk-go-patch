package nas

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

// CreateProtocolMountTarget invokes the nas.CreateProtocolMountTarget API synchronously
func (client *Client) CreateProtocolMountTarget(request *CreateProtocolMountTargetRequest) (response *CreateProtocolMountTargetResponse, err error) {
	response = CreateCreateProtocolMountTargetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProtocolMountTargetWithChan invokes the nas.CreateProtocolMountTarget API asynchronously
func (client *Client) CreateProtocolMountTargetWithChan(request *CreateProtocolMountTargetRequest) (<-chan *CreateProtocolMountTargetResponse, <-chan error) {
	responseChan := make(chan *CreateProtocolMountTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProtocolMountTarget(request)
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

// CreateProtocolMountTargetWithCallback invokes the nas.CreateProtocolMountTarget API asynchronously
func (client *Client) CreateProtocolMountTargetWithCallback(request *CreateProtocolMountTargetRequest, callback func(response *CreateProtocolMountTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProtocolMountTargetResponse
		var err error
		defer close(result)
		response, err = client.CreateProtocolMountTarget(request)
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

// CreateProtocolMountTargetRequest is the request struct for api CreateProtocolMountTarget
type CreateProtocolMountTargetRequest struct {
	*requests.RpcRequest
	ProtocolServiceId string           `position:"Query" name:"ProtocolServiceId"`
	FsetId            string           `position:"Query" name:"FsetId"`
	ClientToken       string           `position:"Query" name:"ClientToken"`
	Description       string           `position:"Query" name:"Description"`
	Path              string           `position:"Query" name:"Path"`
	FileSystemId      string           `position:"Query" name:"FileSystemId"`
	DryRun            requests.Boolean `position:"Query" name:"DryRun"`
	AccessGroupName   string           `position:"Query" name:"AccessGroupName"`
	VSwitchId         string           `position:"Query" name:"VSwitchId"`
	VpcId             string           `position:"Query" name:"VpcId"`
}

// CreateProtocolMountTargetResponse is the response struct for api CreateProtocolMountTarget
type CreateProtocolMountTargetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ExportId  string `json:"ExportId" xml:"ExportId"`
}

// CreateCreateProtocolMountTargetRequest creates a request to invoke CreateProtocolMountTarget API
func CreateCreateProtocolMountTargetRequest() (request *CreateProtocolMountTargetRequest) {
	request = &CreateProtocolMountTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateProtocolMountTarget", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProtocolMountTargetResponse creates a response to parse from CreateProtocolMountTarget response
func CreateCreateProtocolMountTargetResponse() (response *CreateProtocolMountTargetResponse) {
	response = &CreateProtocolMountTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
