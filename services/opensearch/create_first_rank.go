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

// CreateFirstRank invokes the opensearch.CreateFirstRank API synchronously
func (client *Client) CreateFirstRank(request *CreateFirstRankRequest) (response *CreateFirstRankResponse, err error) {
	response = CreateCreateFirstRankResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFirstRankWithChan invokes the opensearch.CreateFirstRank API asynchronously
func (client *Client) CreateFirstRankWithChan(request *CreateFirstRankRequest) (<-chan *CreateFirstRankResponse, <-chan error) {
	responseChan := make(chan *CreateFirstRankResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFirstRank(request)
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

// CreateFirstRankWithCallback invokes the opensearch.CreateFirstRank API asynchronously
func (client *Client) CreateFirstRankWithCallback(request *CreateFirstRankRequest, callback func(response *CreateFirstRankResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFirstRankResponse
		var err error
		defer close(result)
		response, err = client.CreateFirstRank(request)
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

// CreateFirstRankRequest is the request struct for api CreateFirstRank
type CreateFirstRankRequest struct {
	*requests.RoaRequest
	DryRun           requests.Boolean `position:"Query" name:"dryRun"`
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// CreateFirstRankResponse is the response struct for api CreateFirstRank
type CreateFirstRankResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateCreateFirstRankRequest creates a request to invoke CreateFirstRank API
func CreateCreateFirstRankRequest() (request *CreateFirstRankRequest) {
	request = &CreateFirstRankRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateFirstRank", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/first-ranks", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFirstRankResponse creates a response to parse from CreateFirstRank response
func CreateCreateFirstRankResponse() (response *CreateFirstRankResponse) {
	response = &CreateFirstRankResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
