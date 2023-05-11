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

// DeleteFaceUserGroupAndDeviceGroupRelation invokes the linkvisual.DeleteFaceUserGroupAndDeviceGroupRelation API synchronously
func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelation(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest) (response *DeleteFaceUserGroupAndDeviceGroupRelationResponse, err error) {
	response = CreateDeleteFaceUserGroupAndDeviceGroupRelationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceUserGroupAndDeviceGroupRelationWithChan invokes the linkvisual.DeleteFaceUserGroupAndDeviceGroupRelation API asynchronously
func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelationWithChan(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest) (<-chan *DeleteFaceUserGroupAndDeviceGroupRelationResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceUserGroupAndDeviceGroupRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceUserGroupAndDeviceGroupRelation(request)
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

// DeleteFaceUserGroupAndDeviceGroupRelationWithCallback invokes the linkvisual.DeleteFaceUserGroupAndDeviceGroupRelation API asynchronously
func (client *Client) DeleteFaceUserGroupAndDeviceGroupRelationWithCallback(request *DeleteFaceUserGroupAndDeviceGroupRelationRequest, callback func(response *DeleteFaceUserGroupAndDeviceGroupRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceUserGroupAndDeviceGroupRelationResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceUserGroupAndDeviceGroupRelation(request)
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

// DeleteFaceUserGroupAndDeviceGroupRelationRequest is the request struct for api DeleteFaceUserGroupAndDeviceGroupRelation
type DeleteFaceUserGroupAndDeviceGroupRelationRequest struct {
	*requests.RpcRequest
	IsolationId string `position:"Query" name:"IsolationId"`
	ControlId   string `position:"Query" name:"ControlId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// DeleteFaceUserGroupAndDeviceGroupRelationResponse is the response struct for api DeleteFaceUserGroupAndDeviceGroupRelation
type DeleteFaceUserGroupAndDeviceGroupRelationResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDeleteFaceUserGroupAndDeviceGroupRelationRequest creates a request to invoke DeleteFaceUserGroupAndDeviceGroupRelation API
func CreateDeleteFaceUserGroupAndDeviceGroupRelationRequest() (request *DeleteFaceUserGroupAndDeviceGroupRelationRequest) {
	request = &DeleteFaceUserGroupAndDeviceGroupRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "DeleteFaceUserGroupAndDeviceGroupRelation", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFaceUserGroupAndDeviceGroupRelationResponse creates a response to parse from DeleteFaceUserGroupAndDeviceGroupRelation response
func CreateDeleteFaceUserGroupAndDeviceGroupRelationResponse() (response *DeleteFaceUserGroupAndDeviceGroupRelationResponse) {
	response = &DeleteFaceUserGroupAndDeviceGroupRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
