package sas

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

// QueryGroupIdByGroupName invokes the sas.QueryGroupIdByGroupName API synchronously
func (client *Client) QueryGroupIdByGroupName(request *QueryGroupIdByGroupNameRequest) (response *QueryGroupIdByGroupNameResponse, err error) {
	response = CreateQueryGroupIdByGroupNameResponse()
	err = client.DoAction(request, response)
	return
}

// QueryGroupIdByGroupNameWithChan invokes the sas.QueryGroupIdByGroupName API asynchronously
func (client *Client) QueryGroupIdByGroupNameWithChan(request *QueryGroupIdByGroupNameRequest) (<-chan *QueryGroupIdByGroupNameResponse, <-chan error) {
	responseChan := make(chan *QueryGroupIdByGroupNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryGroupIdByGroupName(request)
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

// QueryGroupIdByGroupNameWithCallback invokes the sas.QueryGroupIdByGroupName API asynchronously
func (client *Client) QueryGroupIdByGroupNameWithCallback(request *QueryGroupIdByGroupNameRequest, callback func(response *QueryGroupIdByGroupNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryGroupIdByGroupNameResponse
		var err error
		defer close(result)
		response, err = client.QueryGroupIdByGroupName(request)
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

// QueryGroupIdByGroupNameRequest is the request struct for api QueryGroupIdByGroupName
type QueryGroupIdByGroupNameRequest struct {
	*requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
	SourceIp  string `position:"Query" name:"SourceIp"`
}

// QueryGroupIdByGroupNameResponse is the response struct for api QueryGroupIdByGroupName
type QueryGroupIdByGroupNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   int64  `json:"GroupId" xml:"GroupId"`
}

// CreateQueryGroupIdByGroupNameRequest creates a request to invoke QueryGroupIdByGroupName API
func CreateQueryGroupIdByGroupNameRequest() (request *QueryGroupIdByGroupNameRequest) {
	request = &QueryGroupIdByGroupNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "QueryGroupIdByGroupName", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryGroupIdByGroupNameResponse creates a response to parse from QueryGroupIdByGroupName response
func CreateQueryGroupIdByGroupNameResponse() (response *QueryGroupIdByGroupNameResponse) {
	response = &QueryGroupIdByGroupNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
