package alinlp

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

// ADMMU invokes the alinlp.ADMMU API synchronously
func (client *Client) ADMMU(request *ADMMURequest) (response *ADMMUResponse, err error) {
	response = CreateADMMUResponse()
	err = client.DoAction(request, response)
	return
}

// ADMMUWithChan invokes the alinlp.ADMMU API asynchronously
func (client *Client) ADMMUWithChan(request *ADMMURequest) (<-chan *ADMMUResponse, <-chan error) {
	responseChan := make(chan *ADMMUResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ADMMU(request)
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

// ADMMUWithCallback invokes the alinlp.ADMMU API asynchronously
func (client *Client) ADMMUWithCallback(request *ADMMURequest, callback func(response *ADMMUResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ADMMUResponse
		var err error
		defer close(result)
		response, err = client.ADMMU(request)
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

// ADMMURequest is the request struct for api ADMMU
type ADMMURequest struct {
	*requests.RpcRequest
	Params      string `position:"Body" name:"Params"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
}

// ADMMUResponse is the response struct for api ADMMU
type ADMMUResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateADMMURequest creates a request to invoke ADMMU API
func CreateADMMURequest() (request *ADMMURequest) {
	request = &ADMMURequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "ADMMU", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateADMMUResponse creates a response to parse from ADMMU response
func CreateADMMUResponse() (response *ADMMUResponse) {
	response = &ADMMUResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
