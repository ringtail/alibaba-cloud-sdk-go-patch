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

// GetSecretParametersByPath invokes the oos.GetSecretParametersByPath API synchronously
func (client *Client) GetSecretParametersByPath(request *GetSecretParametersByPathRequest) (response *GetSecretParametersByPathResponse, err error) {
	response = CreateGetSecretParametersByPathResponse()
	err = client.DoAction(request, response)
	return
}

// GetSecretParametersByPathWithChan invokes the oos.GetSecretParametersByPath API asynchronously
func (client *Client) GetSecretParametersByPathWithChan(request *GetSecretParametersByPathRequest) (<-chan *GetSecretParametersByPathResponse, <-chan error) {
	responseChan := make(chan *GetSecretParametersByPathResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSecretParametersByPath(request)
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

// GetSecretParametersByPathWithCallback invokes the oos.GetSecretParametersByPath API asynchronously
func (client *Client) GetSecretParametersByPathWithCallback(request *GetSecretParametersByPathRequest, callback func(response *GetSecretParametersByPathResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSecretParametersByPathResponse
		var err error
		defer close(result)
		response, err = client.GetSecretParametersByPath(request)
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

// GetSecretParametersByPathRequest is the request struct for api GetSecretParametersByPath
type GetSecretParametersByPathRequest struct {
	*requests.RpcRequest
	WithDecryption requests.Boolean `position:"Query" name:"WithDecryption"`
	Recursive      requests.Boolean `position:"Query" name:"Recursive"`
	Path           string           `position:"Query" name:"Path"`
	NextToken      string           `position:"Query" name:"NextToken"`
	MaxResults     requests.Integer `position:"Query" name:"MaxResults"`
}

// GetSecretParametersByPathResponse is the response struct for api GetSecretParametersByPath
type GetSecretParametersByPathResponse struct {
	*responses.BaseResponse
	NextToken  string      `json:"NextToken" xml:"NextToken"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TotalCount int         `json:"TotalCount" xml:"TotalCount"`
	MaxResults int         `json:"MaxResults" xml:"MaxResults"`
	Parameters []Parameter `json:"Parameters" xml:"Parameters"`
}

// CreateGetSecretParametersByPathRequest creates a request to invoke GetSecretParametersByPath API
func CreateGetSecretParametersByPathRequest() (request *GetSecretParametersByPathRequest) {
	request = &GetSecretParametersByPathRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "GetSecretParametersByPath", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSecretParametersByPathResponse creates a response to parse from GetSecretParametersByPath response
func CreateGetSecretParametersByPathResponse() (response *GetSecretParametersByPathResponse) {
	response = &GetSecretParametersByPathResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
