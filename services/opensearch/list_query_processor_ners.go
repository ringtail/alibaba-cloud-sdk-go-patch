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

// ListQueryProcessorNers invokes the opensearch.ListQueryProcessorNers API synchronously
func (client *Client) ListQueryProcessorNers(request *ListQueryProcessorNersRequest) (response *ListQueryProcessorNersResponse, err error) {
	response = CreateListQueryProcessorNersResponse()
	err = client.DoAction(request, response)
	return
}

// ListQueryProcessorNersWithChan invokes the opensearch.ListQueryProcessorNers API asynchronously
func (client *Client) ListQueryProcessorNersWithChan(request *ListQueryProcessorNersRequest) (<-chan *ListQueryProcessorNersResponse, <-chan error) {
	responseChan := make(chan *ListQueryProcessorNersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQueryProcessorNers(request)
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

// ListQueryProcessorNersWithCallback invokes the opensearch.ListQueryProcessorNers API asynchronously
func (client *Client) ListQueryProcessorNersWithCallback(request *ListQueryProcessorNersRequest, callback func(response *ListQueryProcessorNersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQueryProcessorNersResponse
		var err error
		defer close(result)
		response, err = client.ListQueryProcessorNers(request)
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

// ListQueryProcessorNersRequest is the request struct for api ListQueryProcessorNers
type ListQueryProcessorNersRequest struct {
	*requests.RoaRequest
	Domain string `position:"Query" name:"domain"`
}

// ListQueryProcessorNersResponse is the response struct for api ListQueryProcessorNers
type ListQueryProcessorNersResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListQueryProcessorNersRequest creates a request to invoke ListQueryProcessorNers API
func CreateListQueryProcessorNersRequest() (request *ListQueryProcessorNersRequest) {
	request = &ListQueryProcessorNersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListQueryProcessorNers", "/v4/openapi/query-processor/ner/default-priorities", "", "")
	request.Method = requests.GET
	return
}

// CreateListQueryProcessorNersResponse creates a response to parse from ListQueryProcessorNers response
func CreateListQueryProcessorNersResponse() (response *ListQueryProcessorNersResponse) {
	response = &ListQueryProcessorNersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
