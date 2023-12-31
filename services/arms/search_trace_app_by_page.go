package arms

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

// SearchTraceAppByPage invokes the arms.SearchTraceAppByPage API synchronously
func (client *Client) SearchTraceAppByPage(request *SearchTraceAppByPageRequest) (response *SearchTraceAppByPageResponse, err error) {
	response = CreateSearchTraceAppByPageResponse()
	err = client.DoAction(request, response)
	return
}

// SearchTraceAppByPageWithChan invokes the arms.SearchTraceAppByPage API asynchronously
func (client *Client) SearchTraceAppByPageWithChan(request *SearchTraceAppByPageRequest) (<-chan *SearchTraceAppByPageResponse, <-chan error) {
	responseChan := make(chan *SearchTraceAppByPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchTraceAppByPage(request)
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

// SearchTraceAppByPageWithCallback invokes the arms.SearchTraceAppByPage API asynchronously
func (client *Client) SearchTraceAppByPageWithCallback(request *SearchTraceAppByPageRequest, callback func(response *SearchTraceAppByPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchTraceAppByPageResponse
		var err error
		defer close(result)
		response, err = client.SearchTraceAppByPage(request)
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

// SearchTraceAppByPageRequest is the request struct for api SearchTraceAppByPage
type SearchTraceAppByPageRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer            `position:"Query" name:"PageNumber"`
	Tags            *[]SearchTraceAppByPageTags `position:"Query" name:"Tags"  type:"Repeated"`
	ResourceGroupId string                      `position:"Query" name:"ResourceGroupId"`
	TraceAppName    string                      `position:"Query" name:"TraceAppName"`
	PageSize        requests.Integer            `position:"Query" name:"PageSize"`
}

// SearchTraceAppByPageTags is a repeated param struct in SearchTraceAppByPageRequest
type SearchTraceAppByPageTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// SearchTraceAppByPageResponse is the response struct for api SearchTraceAppByPage
type SearchTraceAppByPageResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageBean  PageBean `json:"PageBean" xml:"PageBean"`
}

// CreateSearchTraceAppByPageRequest creates a request to invoke SearchTraceAppByPage API
func CreateSearchTraceAppByPageRequest() (request *SearchTraceAppByPageRequest) {
	request = &SearchTraceAppByPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SearchTraceAppByPage", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchTraceAppByPageResponse creates a response to parse from SearchTraceAppByPage response
func CreateSearchTraceAppByPageResponse() (response *SearchTraceAppByPageResponse) {
	response = &SearchTraceAppByPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
