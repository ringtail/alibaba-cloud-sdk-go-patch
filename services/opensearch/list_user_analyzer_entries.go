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

// ListUserAnalyzerEntries invokes the opensearch.ListUserAnalyzerEntries API synchronously
func (client *Client) ListUserAnalyzerEntries(request *ListUserAnalyzerEntriesRequest) (response *ListUserAnalyzerEntriesResponse, err error) {
	response = CreateListUserAnalyzerEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserAnalyzerEntriesWithChan invokes the opensearch.ListUserAnalyzerEntries API asynchronously
func (client *Client) ListUserAnalyzerEntriesWithChan(request *ListUserAnalyzerEntriesRequest) (<-chan *ListUserAnalyzerEntriesResponse, <-chan error) {
	responseChan := make(chan *ListUserAnalyzerEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserAnalyzerEntries(request)
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

// ListUserAnalyzerEntriesWithCallback invokes the opensearch.ListUserAnalyzerEntries API asynchronously
func (client *Client) ListUserAnalyzerEntriesWithCallback(request *ListUserAnalyzerEntriesRequest, callback func(response *ListUserAnalyzerEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserAnalyzerEntriesResponse
		var err error
		defer close(result)
		response, err = client.ListUserAnalyzerEntries(request)
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

// ListUserAnalyzerEntriesRequest is the request struct for api ListUserAnalyzerEntries
type ListUserAnalyzerEntriesRequest struct {
	*requests.RoaRequest
	Name       string           `position:"Path" name:"name"`
	PageSize   requests.Integer `position:"Query" name:"pageSize"`
	Word       string           `position:"Query" name:"word"`
	PageNumber requests.Integer `position:"Query" name:"pageNumber"`
}

// ListUserAnalyzerEntriesResponse is the response struct for api ListUserAnalyzerEntries
type ListUserAnalyzerEntriesResponse struct {
	*responses.BaseResponse
	Result    map[string]interface{} `json:"result" xml:"result"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateListUserAnalyzerEntriesRequest creates a request to invoke ListUserAnalyzerEntries API
func CreateListUserAnalyzerEntriesRequest() (request *ListUserAnalyzerEntriesRequest) {
	request = &ListUserAnalyzerEntriesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListUserAnalyzerEntries", "/v4/openapi/user-analyzers/[name]/entries", "", "")
	request.Method = requests.GET
	return
}

// CreateListUserAnalyzerEntriesResponse creates a response to parse from ListUserAnalyzerEntries response
func CreateListUserAnalyzerEntriesResponse() (response *ListUserAnalyzerEntriesResponse) {
	response = &ListUserAnalyzerEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
