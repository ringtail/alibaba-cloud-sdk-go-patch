package ahas_openapi

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

// PageableQueryExperimentTaskByExperimentId invokes the ahas_openapi.PageableQueryExperimentTaskByExperimentId API synchronously
func (client *Client) PageableQueryExperimentTaskByExperimentId(request *PageableQueryExperimentTaskByExperimentIdRequest) (response *PageableQueryExperimentTaskByExperimentIdResponse, err error) {
	response = CreatePageableQueryExperimentTaskByExperimentIdResponse()
	err = client.DoAction(request, response)
	return
}

// PageableQueryExperimentTaskByExperimentIdWithChan invokes the ahas_openapi.PageableQueryExperimentTaskByExperimentId API asynchronously
func (client *Client) PageableQueryExperimentTaskByExperimentIdWithChan(request *PageableQueryExperimentTaskByExperimentIdRequest) (<-chan *PageableQueryExperimentTaskByExperimentIdResponse, <-chan error) {
	responseChan := make(chan *PageableQueryExperimentTaskByExperimentIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PageableQueryExperimentTaskByExperimentId(request)
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

// PageableQueryExperimentTaskByExperimentIdWithCallback invokes the ahas_openapi.PageableQueryExperimentTaskByExperimentId API asynchronously
func (client *Client) PageableQueryExperimentTaskByExperimentIdWithCallback(request *PageableQueryExperimentTaskByExperimentIdRequest, callback func(response *PageableQueryExperimentTaskByExperimentIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PageableQueryExperimentTaskByExperimentIdResponse
		var err error
		defer close(result)
		response, err = client.PageableQueryExperimentTaskByExperimentId(request)
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

// PageableQueryExperimentTaskByExperimentIdRequest is the request struct for api PageableQueryExperimentTaskByExperimentId
type PageableQueryExperimentTaskByExperimentIdRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	Size         requests.Integer `position:"Query" name:"Size"`
	Namespace    string           `position:"Query" name:"Namespace"`
	ExperimentId string           `position:"Query" name:"ExperimentId"`
	Page         requests.Integer `position:"Query" name:"Page"`
}

// PageableQueryExperimentTaskByExperimentIdResponse is the response struct for api PageableQueryExperimentTaskByExperimentId
type PageableQueryExperimentTaskByExperimentIdResponse struct {
	*responses.BaseResponse
	Pages           int                  `json:"Pages" xml:"Pages"`
	RequestId       string               `json:"RequestId" xml:"RequestId"`
	Message         string               `json:"Message" xml:"Message"`
	PageSize        int                  `json:"PageSize" xml:"PageSize"`
	CurrentPage     int                  `json:"CurrentPage" xml:"CurrentPage"`
	Total           int                  `json:"Total" xml:"Total"`
	HttpStatusCode  int                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code            string               `json:"Code" xml:"Code"`
	Success         bool                 `json:"Success" xml:"Success"`
	ExperimentTasks []ExperimentTaskInfo `json:"ExperimentTasks" xml:"ExperimentTasks"`
}

// CreatePageableQueryExperimentTaskByExperimentIdRequest creates a request to invoke PageableQueryExperimentTaskByExperimentId API
func CreatePageableQueryExperimentTaskByExperimentIdRequest() (request *PageableQueryExperimentTaskByExperimentIdRequest) {
	request = &PageableQueryExperimentTaskByExperimentIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "PageableQueryExperimentTaskByExperimentId", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePageableQueryExperimentTaskByExperimentIdResponse creates a response to parse from PageableQueryExperimentTaskByExperimentId response
func CreatePageableQueryExperimentTaskByExperimentIdResponse() (response *PageableQueryExperimentTaskByExperimentIdResponse) {
	response = &PageableQueryExperimentTaskByExperimentIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
