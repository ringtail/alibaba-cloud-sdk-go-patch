package dms_enterprise

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

// GetOnlineDDLProgress invokes the dms_enterprise.GetOnlineDDLProgress API synchronously
func (client *Client) GetOnlineDDLProgress(request *GetOnlineDDLProgressRequest) (response *GetOnlineDDLProgressResponse, err error) {
	response = CreateGetOnlineDDLProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetOnlineDDLProgressWithChan invokes the dms_enterprise.GetOnlineDDLProgress API asynchronously
func (client *Client) GetOnlineDDLProgressWithChan(request *GetOnlineDDLProgressRequest) (<-chan *GetOnlineDDLProgressResponse, <-chan error) {
	responseChan := make(chan *GetOnlineDDLProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOnlineDDLProgress(request)
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

// GetOnlineDDLProgressWithCallback invokes the dms_enterprise.GetOnlineDDLProgress API asynchronously
func (client *Client) GetOnlineDDLProgressWithCallback(request *GetOnlineDDLProgressRequest, callback func(response *GetOnlineDDLProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOnlineDDLProgressResponse
		var err error
		defer close(result)
		response, err = client.GetOnlineDDLProgress(request)
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

// GetOnlineDDLProgressRequest is the request struct for api GetOnlineDDLProgress
type GetOnlineDDLProgressRequest struct {
	*requests.RpcRequest
	Tid         requests.Integer `position:"Query" name:"Tid"`
	JobDetailId requests.Integer `position:"Query" name:"JobDetailId"`
}

// GetOnlineDDLProgressResponse is the response struct for api GetOnlineDDLProgress
type GetOnlineDDLProgressResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	Success             bool                `json:"Success" xml:"Success"`
	ErrorMessage        string              `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode           string              `json:"ErrorCode" xml:"ErrorCode"`
	OnlineDDLTaskDetail OnlineDDLTaskDetail `json:"OnlineDDLTaskDetail" xml:"OnlineDDLTaskDetail"`
}

// CreateGetOnlineDDLProgressRequest creates a request to invoke GetOnlineDDLProgress API
func CreateGetOnlineDDLProgressRequest() (request *GetOnlineDDLProgressRequest) {
	request = &GetOnlineDDLProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetOnlineDDLProgress", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOnlineDDLProgressResponse creates a response to parse from GetOnlineDDLProgress response
func CreateGetOnlineDDLProgressResponse() (response *GetOnlineDDLProgressResponse) {
	response = &GetOnlineDDLProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
