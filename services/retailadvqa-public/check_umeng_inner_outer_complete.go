package retailadvqa_public

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

// CheckUmengInnerOuterComplete invokes the retailadvqa_public.CheckUmengInnerOuterComplete API synchronously
func (client *Client) CheckUmengInnerOuterComplete(request *CheckUmengInnerOuterCompleteRequest) (response *CheckUmengInnerOuterCompleteResponse, err error) {
	response = CreateCheckUmengInnerOuterCompleteResponse()
	err = client.DoAction(request, response)
	return
}

// CheckUmengInnerOuterCompleteWithChan invokes the retailadvqa_public.CheckUmengInnerOuterComplete API asynchronously
func (client *Client) CheckUmengInnerOuterCompleteWithChan(request *CheckUmengInnerOuterCompleteRequest) (<-chan *CheckUmengInnerOuterCompleteResponse, <-chan error) {
	responseChan := make(chan *CheckUmengInnerOuterCompleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckUmengInnerOuterComplete(request)
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

// CheckUmengInnerOuterCompleteWithCallback invokes the retailadvqa_public.CheckUmengInnerOuterComplete API asynchronously
func (client *Client) CheckUmengInnerOuterCompleteWithCallback(request *CheckUmengInnerOuterCompleteRequest, callback func(response *CheckUmengInnerOuterCompleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckUmengInnerOuterCompleteResponse
		var err error
		defer close(result)
		response, err = client.CheckUmengInnerOuterComplete(request)
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

// CheckUmengInnerOuterCompleteRequest is the request struct for api CheckUmengInnerOuterComplete
type CheckUmengInnerOuterCompleteRequest struct {
	*requests.RpcRequest
	AccessId   string           `position:"Query" name:"AccessId"`
	TenantId   string           `position:"Query" name:"TenantId"`
	AudienceId string           `position:"Query" name:"AudienceId"`
	Message    string           `position:"Query" name:"Message"`
	TaskId     string           `position:"Query" name:"TaskId"`
	Status     requests.Integer `position:"Query" name:"Status"`
}

// CheckUmengInnerOuterCompleteResponse is the response struct for api CheckUmengInnerOuterComplete
type CheckUmengInnerOuterCompleteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      bool   `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCheckUmengInnerOuterCompleteRequest creates a request to invoke CheckUmengInnerOuterComplete API
func CreateCheckUmengInnerOuterCompleteRequest() (request *CheckUmengInnerOuterCompleteRequest) {
	request = &CheckUmengInnerOuterCompleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CheckUmengInnerOuterComplete", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckUmengInnerOuterCompleteResponse creates a response to parse from CheckUmengInnerOuterComplete response
func CreateCheckUmengInnerOuterCompleteResponse() (response *CheckUmengInnerOuterCompleteResponse) {
	response = &CheckUmengInnerOuterCompleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
