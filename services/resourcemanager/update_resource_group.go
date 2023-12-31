package resourcemanager

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

// UpdateResourceGroup invokes the resourcemanager.UpdateResourceGroup API synchronously
func (client *Client) UpdateResourceGroup(request *UpdateResourceGroupRequest) (response *UpdateResourceGroupResponse, err error) {
	response = CreateUpdateResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateResourceGroupWithChan invokes the resourcemanager.UpdateResourceGroup API asynchronously
func (client *Client) UpdateResourceGroupWithChan(request *UpdateResourceGroupRequest) (<-chan *UpdateResourceGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateResourceGroup(request)
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

// UpdateResourceGroupWithCallback invokes the resourcemanager.UpdateResourceGroup API asynchronously
func (client *Client) UpdateResourceGroupWithCallback(request *UpdateResourceGroupRequest, callback func(response *UpdateResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateResourceGroup(request)
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

// UpdateResourceGroupRequest is the request struct for api UpdateResourceGroup
type UpdateResourceGroupRequest struct {
	*requests.RpcRequest
	NewDisplayName  string `position:"Query" name:"NewDisplayName"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
}

// UpdateResourceGroupResponse is the response struct for api UpdateResourceGroup
type UpdateResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResourceGroup ResourceGroup `json:"ResourceGroup" xml:"ResourceGroup"`
}

// CreateUpdateResourceGroupRequest creates a request to invoke UpdateResourceGroup API
func CreateUpdateResourceGroupRequest() (request *UpdateResourceGroupRequest) {
	request = &UpdateResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "UpdateResourceGroup", "resourcemanager", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateResourceGroupResponse creates a response to parse from UpdateResourceGroup response
func CreateUpdateResourceGroupResponse() (response *UpdateResourceGroupResponse) {
	response = &UpdateResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
