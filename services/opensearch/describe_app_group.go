package opensearch

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

// DescribeAppGroup invokes the opensearch.DescribeAppGroup API synchronously
func (client *Client) DescribeAppGroup(request *DescribeAppGroupRequest) (response *DescribeAppGroupResponse, err error) {
	response = CreateDescribeAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAppGroupWithChan invokes the opensearch.DescribeAppGroup API asynchronously
func (client *Client) DescribeAppGroupWithChan(request *DescribeAppGroupRequest) (<-chan *DescribeAppGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAppGroup(request)
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

// DescribeAppGroupWithCallback invokes the opensearch.DescribeAppGroup API asynchronously
func (client *Client) DescribeAppGroupWithCallback(request *DescribeAppGroupRequest, callback func(response *DescribeAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAppGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeAppGroup(request)
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

// DescribeAppGroupRequest is the request struct for api DescribeAppGroup
type DescribeAppGroupRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// DescribeAppGroupResponse is the response struct for api DescribeAppGroup
type DescribeAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"requestId" xml:"requestId"`
	Result    ResultInDescribeAppGroup `json:"result" xml:"result"`
}

// CreateDescribeAppGroupRequest creates a request to invoke DescribeAppGroup API
func CreateDescribeAppGroupRequest() (request *DescribeAppGroupRequest) {
	request = &DescribeAppGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeAppGroup", "/v4/openapi/app-groups/[appGroupIdentity]", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeAppGroupResponse creates a response to parse from DescribeAppGroup response
func CreateDescribeAppGroupResponse() (response *DescribeAppGroupResponse) {
	response = &DescribeAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
