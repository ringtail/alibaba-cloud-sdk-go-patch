package bssopenapi

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

// ChangeResellerConsumeAmount invokes the bssopenapi.ChangeResellerConsumeAmount API synchronously
func (client *Client) ChangeResellerConsumeAmount(request *ChangeResellerConsumeAmountRequest) (response *ChangeResellerConsumeAmountResponse, err error) {
	response = CreateChangeResellerConsumeAmountResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeResellerConsumeAmountWithChan invokes the bssopenapi.ChangeResellerConsumeAmount API asynchronously
func (client *Client) ChangeResellerConsumeAmountWithChan(request *ChangeResellerConsumeAmountRequest) (<-chan *ChangeResellerConsumeAmountResponse, <-chan error) {
	responseChan := make(chan *ChangeResellerConsumeAmountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeResellerConsumeAmount(request)
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

// ChangeResellerConsumeAmountWithCallback invokes the bssopenapi.ChangeResellerConsumeAmount API asynchronously
func (client *Client) ChangeResellerConsumeAmountWithCallback(request *ChangeResellerConsumeAmountRequest, callback func(response *ChangeResellerConsumeAmountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeResellerConsumeAmountResponse
		var err error
		defer close(result)
		response, err = client.ChangeResellerConsumeAmount(request)
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

// ChangeResellerConsumeAmountRequest is the request struct for api ChangeResellerConsumeAmount
type ChangeResellerConsumeAmountRequest struct {
	*requests.RpcRequest
	Amount       string           `position:"Query" name:"Amount"`
	OutBizId     string           `position:"Query" name:"OutBizId"`
	Source       string           `position:"Query" name:"Source"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	BusinessType string           `position:"Query" name:"BusinessType"`
	AdjustType   string           `position:"Query" name:"AdjustType"`
	ExtendMap    string           `position:"Query" name:"ExtendMap"`
	Currency     string           `position:"Query" name:"Currency"`
}

// ChangeResellerConsumeAmountResponse is the response struct for api ChangeResellerConsumeAmount
type ChangeResellerConsumeAmountResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateChangeResellerConsumeAmountRequest creates a request to invoke ChangeResellerConsumeAmount API
func CreateChangeResellerConsumeAmountRequest() (request *ChangeResellerConsumeAmountRequest) {
	request = &ChangeResellerConsumeAmountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "ChangeResellerConsumeAmount", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeResellerConsumeAmountResponse creates a response to parse from ChangeResellerConsumeAmount response
func CreateChangeResellerConsumeAmountResponse() (response *ChangeResellerConsumeAmountResponse) {
	response = &ChangeResellerConsumeAmountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
