package sae

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

// DeleteGreyTagRoute invokes the sae.DeleteGreyTagRoute API synchronously
func (client *Client) DeleteGreyTagRoute(request *DeleteGreyTagRouteRequest) (response *DeleteGreyTagRouteResponse, err error) {
	response = CreateDeleteGreyTagRouteResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGreyTagRouteWithChan invokes the sae.DeleteGreyTagRoute API asynchronously
func (client *Client) DeleteGreyTagRouteWithChan(request *DeleteGreyTagRouteRequest) (<-chan *DeleteGreyTagRouteResponse, <-chan error) {
	responseChan := make(chan *DeleteGreyTagRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGreyTagRoute(request)
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

// DeleteGreyTagRouteWithCallback invokes the sae.DeleteGreyTagRoute API asynchronously
func (client *Client) DeleteGreyTagRouteWithCallback(request *DeleteGreyTagRouteRequest, callback func(response *DeleteGreyTagRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGreyTagRouteResponse
		var err error
		defer close(result)
		response, err = client.DeleteGreyTagRoute(request)
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

// DeleteGreyTagRouteRequest is the request struct for api DeleteGreyTagRoute
type DeleteGreyTagRouteRequest struct {
	*requests.RoaRequest
	GreyTagRouteId requests.Integer `position:"Query" name:"GreyTagRouteId"`
}

// DeleteGreyTagRouteResponse is the response struct for api DeleteGreyTagRoute
type DeleteGreyTagRouteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteGreyTagRouteRequest creates a request to invoke DeleteGreyTagRoute API
func CreateDeleteGreyTagRouteRequest() (request *DeleteGreyTagRouteRequest) {
	request = &DeleteGreyTagRouteRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DeleteGreyTagRoute", "/pop/v1/sam/tagroute/greyTagRoute", "serverless", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteGreyTagRouteResponse creates a response to parse from DeleteGreyTagRoute response
func CreateDeleteGreyTagRouteResponse() (response *DeleteGreyTagRouteResponse) {
	response = &DeleteGreyTagRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
