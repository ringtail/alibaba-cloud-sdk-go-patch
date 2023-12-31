package dataworks_public

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

// SubmitDataServiceApi invokes the dataworks_public.SubmitDataServiceApi API synchronously
func (client *Client) SubmitDataServiceApi(request *SubmitDataServiceApiRequest) (response *SubmitDataServiceApiResponse, err error) {
	response = CreateSubmitDataServiceApiResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitDataServiceApiWithChan invokes the dataworks_public.SubmitDataServiceApi API asynchronously
func (client *Client) SubmitDataServiceApiWithChan(request *SubmitDataServiceApiRequest) (<-chan *SubmitDataServiceApiResponse, <-chan error) {
	responseChan := make(chan *SubmitDataServiceApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitDataServiceApi(request)
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

// SubmitDataServiceApiWithCallback invokes the dataworks_public.SubmitDataServiceApi API asynchronously
func (client *Client) SubmitDataServiceApiWithCallback(request *SubmitDataServiceApiRequest, callback func(response *SubmitDataServiceApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitDataServiceApiResponse
		var err error
		defer close(result)
		response, err = client.SubmitDataServiceApi(request)
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

// SubmitDataServiceApiRequest is the request struct for api SubmitDataServiceApi
type SubmitDataServiceApiRequest struct {
	*requests.RpcRequest
	TenantId  requests.Integer `position:"Body" name:"TenantId"`
	ProjectId requests.Integer `position:"Body" name:"ProjectId"`
	ApiId     requests.Integer `position:"Body" name:"ApiId"`
}

// SubmitDataServiceApiResponse is the response struct for api SubmitDataServiceApi
type SubmitDataServiceApiResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSubmitDataServiceApiRequest creates a request to invoke SubmitDataServiceApi API
func CreateSubmitDataServiceApiRequest() (request *SubmitDataServiceApiRequest) {
	request = &SubmitDataServiceApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "SubmitDataServiceApi", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitDataServiceApiResponse creates a response to parse from SubmitDataServiceApi response
func CreateSubmitDataServiceApiResponse() (response *SubmitDataServiceApiResponse) {
	response = &SubmitDataServiceApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
