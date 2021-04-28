package ccc

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

// ListUserLevelsOfSkillGroup invokes the ccc.ListUserLevelsOfSkillGroup API synchronously
func (client *Client) ListUserLevelsOfSkillGroup(request *ListUserLevelsOfSkillGroupRequest) (response *ListUserLevelsOfSkillGroupResponse, err error) {
	response = CreateListUserLevelsOfSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserLevelsOfSkillGroupWithChan invokes the ccc.ListUserLevelsOfSkillGroup API asynchronously
func (client *Client) ListUserLevelsOfSkillGroupWithChan(request *ListUserLevelsOfSkillGroupRequest) (<-chan *ListUserLevelsOfSkillGroupResponse, <-chan error) {
	responseChan := make(chan *ListUserLevelsOfSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserLevelsOfSkillGroup(request)
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

// ListUserLevelsOfSkillGroupWithCallback invokes the ccc.ListUserLevelsOfSkillGroup API asynchronously
func (client *Client) ListUserLevelsOfSkillGroupWithCallback(request *ListUserLevelsOfSkillGroupRequest, callback func(response *ListUserLevelsOfSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserLevelsOfSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.ListUserLevelsOfSkillGroup(request)
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

// ListUserLevelsOfSkillGroupRequest is the request struct for api ListUserLevelsOfSkillGroup
type ListUserLevelsOfSkillGroupRequest struct {
	*requests.RpcRequest
	IsMember      requests.Boolean `position:"Query" name:"IsMember"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SearchPattern string           `position:"Query" name:"SearchPattern"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	SkillGroupId  string           `position:"Query" name:"SkillGroupId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListUserLevelsOfSkillGroupResponse is the response struct for api ListUserLevelsOfSkillGroup
type ListUserLevelsOfSkillGroupResponse struct {
	*responses.BaseResponse
	Code           string                           `json:"Code" xml:"Code"`
	HttpStatusCode int                              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                           `json:"Message" xml:"Message"`
	RequestId      string                           `json:"RequestId" xml:"RequestId"`
	Data           DataInListUserLevelsOfSkillGroup `json:"Data" xml:"Data"`
}

// CreateListUserLevelsOfSkillGroupRequest creates a request to invoke ListUserLevelsOfSkillGroup API
func CreateListUserLevelsOfSkillGroupRequest() (request *ListUserLevelsOfSkillGroupRequest) {
	request = &ListUserLevelsOfSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListUserLevelsOfSkillGroup", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUserLevelsOfSkillGroupResponse creates a response to parse from ListUserLevelsOfSkillGroup response
func CreateListUserLevelsOfSkillGroupResponse() (response *ListUserLevelsOfSkillGroupResponse) {
	response = &ListUserLevelsOfSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
