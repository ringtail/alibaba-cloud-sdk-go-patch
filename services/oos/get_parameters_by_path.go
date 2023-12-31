package oos

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

// GetParametersByPath invokes the oos.GetParametersByPath API synchronously
func (client *Client) GetParametersByPath(request *GetParametersByPathRequest) (response *GetParametersByPathResponse, err error) {
	response = CreateGetParametersByPathResponse()
	err = client.DoAction(request, response)
	return
}

// GetParametersByPathWithChan invokes the oos.GetParametersByPath API asynchronously
func (client *Client) GetParametersByPathWithChan(request *GetParametersByPathRequest) (<-chan *GetParametersByPathResponse, <-chan error) {
	responseChan := make(chan *GetParametersByPathResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetParametersByPath(request)
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

// GetParametersByPathWithCallback invokes the oos.GetParametersByPath API asynchronously
func (client *Client) GetParametersByPathWithCallback(request *GetParametersByPathRequest, callback func(response *GetParametersByPathResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetParametersByPathResponse
		var err error
		defer close(result)
		response, err = client.GetParametersByPath(request)
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

// GetParametersByPathRequest is the request struct for api GetParametersByPath
type GetParametersByPathRequest struct {
	*requests.RpcRequest
	Recursive  requests.Boolean `position:"Query" name:"Recursive"`
	Path       string           `position:"Query" name:"Path"`
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// GetParametersByPathResponse is the response struct for api GetParametersByPath
type GetParametersByPathResponse struct {
	*responses.BaseResponse
	NextToken  string      `json:"NextToken" xml:"NextToken"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TotalCount int         `json:"TotalCount" xml:"TotalCount"`
	MaxResults int         `json:"MaxResults" xml:"MaxResults"`
	Parameters []Parameter `json:"Parameters" xml:"Parameters"`
}

// CreateGetParametersByPathRequest creates a request to invoke GetParametersByPath API
func CreateGetParametersByPathRequest() (request *GetParametersByPathRequest) {
	request = &GetParametersByPathRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "GetParametersByPath", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetParametersByPathResponse creates a response to parse from GetParametersByPath response
func CreateGetParametersByPathResponse() (response *GetParametersByPathResponse) {
	response = &GetParametersByPathResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
