package quickbi_public

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

// QueryEmbeddedStaus invokes the quickbi_public.QueryEmbeddedStaus API synchronously
func (client *Client) QueryEmbeddedStaus(request *QueryEmbeddedStausRequest) (response *QueryEmbeddedStausResponse, err error) {
	response = CreateQueryEmbeddedStausResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEmbeddedStausWithChan invokes the quickbi_public.QueryEmbeddedStaus API asynchronously
func (client *Client) QueryEmbeddedStausWithChan(request *QueryEmbeddedStausRequest) (<-chan *QueryEmbeddedStausResponse, <-chan error) {
	responseChan := make(chan *QueryEmbeddedStausResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEmbeddedStaus(request)
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

// QueryEmbeddedStausWithCallback invokes the quickbi_public.QueryEmbeddedStaus API asynchronously
func (client *Client) QueryEmbeddedStausWithCallback(request *QueryEmbeddedStausRequest, callback func(response *QueryEmbeddedStausResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEmbeddedStausResponse
		var err error
		defer close(result)
		response, err = client.QueryEmbeddedStaus(request)
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

// QueryEmbeddedStausRequest is the request struct for api QueryEmbeddedStaus
type QueryEmbeddedStausRequest struct {
	*requests.RpcRequest
	WorksId     string `position:"Query" name:"WorksId"`
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryEmbeddedStausResponse is the response struct for api QueryEmbeddedStaus
type QueryEmbeddedStausResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateQueryEmbeddedStausRequest creates a request to invoke QueryEmbeddedStaus API
func CreateQueryEmbeddedStausRequest() (request *QueryEmbeddedStausRequest) {
	request = &QueryEmbeddedStausRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2020-08-07", "QueryEmbeddedStaus", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEmbeddedStausResponse creates a response to parse from QueryEmbeddedStaus response
func CreateQueryEmbeddedStausResponse() (response *QueryEmbeddedStausResponse) {
	response = &QueryEmbeddedStausResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
