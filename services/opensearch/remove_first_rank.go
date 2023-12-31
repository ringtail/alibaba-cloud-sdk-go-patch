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

// RemoveFirstRank invokes the opensearch.RemoveFirstRank API synchronously
func (client *Client) RemoveFirstRank(request *RemoveFirstRankRequest) (response *RemoveFirstRankResponse, err error) {
	response = CreateRemoveFirstRankResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveFirstRankWithChan invokes the opensearch.RemoveFirstRank API asynchronously
func (client *Client) RemoveFirstRankWithChan(request *RemoveFirstRankRequest) (<-chan *RemoveFirstRankResponse, <-chan error) {
	responseChan := make(chan *RemoveFirstRankResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveFirstRank(request)
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

// RemoveFirstRankWithCallback invokes the opensearch.RemoveFirstRank API asynchronously
func (client *Client) RemoveFirstRankWithCallback(request *RemoveFirstRankRequest, callback func(response *RemoveFirstRankResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveFirstRankResponse
		var err error
		defer close(result)
		response, err = client.RemoveFirstRank(request)
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

// RemoveFirstRankRequest is the request struct for api RemoveFirstRank
type RemoveFirstRankRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	Name             string           `position:"Path" name:"name"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// RemoveFirstRankResponse is the response struct for api RemoveFirstRank
type RemoveFirstRankResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateRemoveFirstRankRequest creates a request to invoke RemoveFirstRank API
func CreateRemoveFirstRankRequest() (request *RemoveFirstRankRequest) {
	request = &RemoveFirstRankRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "RemoveFirstRank", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/first-ranks/[name]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateRemoveFirstRankResponse creates a response to parse from RemoveFirstRank response
func CreateRemoveFirstRankResponse() (response *RemoveFirstRankResponse) {
	response = &RemoveFirstRankResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
