package quickbi_public

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

// AllotDatasetAccelerationTask invokes the quickbi_public.AllotDatasetAccelerationTask API synchronously
func (client *Client) AllotDatasetAccelerationTask(request *AllotDatasetAccelerationTaskRequest) (response *AllotDatasetAccelerationTaskResponse, err error) {
	response = CreateAllotDatasetAccelerationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// AllotDatasetAccelerationTaskWithChan invokes the quickbi_public.AllotDatasetAccelerationTask API asynchronously
func (client *Client) AllotDatasetAccelerationTaskWithChan(request *AllotDatasetAccelerationTaskRequest) (<-chan *AllotDatasetAccelerationTaskResponse, <-chan error) {
	responseChan := make(chan *AllotDatasetAccelerationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllotDatasetAccelerationTask(request)
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

// AllotDatasetAccelerationTaskWithCallback invokes the quickbi_public.AllotDatasetAccelerationTask API asynchronously
func (client *Client) AllotDatasetAccelerationTaskWithCallback(request *AllotDatasetAccelerationTaskRequest, callback func(response *AllotDatasetAccelerationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllotDatasetAccelerationTaskResponse
		var err error
		defer close(result)
		response, err = client.AllotDatasetAccelerationTask(request)
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

// AllotDatasetAccelerationTaskRequest is the request struct for api AllotDatasetAccelerationTask
type AllotDatasetAccelerationTaskRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
	CubeId      string `position:"Query" name:"CubeId"`
}

// AllotDatasetAccelerationTaskResponse is the response struct for api AllotDatasetAccelerationTask
type AllotDatasetAccelerationTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAllotDatasetAccelerationTaskRequest creates a request to invoke AllotDatasetAccelerationTask API
func CreateAllotDatasetAccelerationTaskRequest() (request *AllotDatasetAccelerationTaskRequest) {
	request = &AllotDatasetAccelerationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "AllotDatasetAccelerationTask", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllotDatasetAccelerationTaskResponse creates a response to parse from AllotDatasetAccelerationTask response
func CreateAllotDatasetAccelerationTaskResponse() (response *AllotDatasetAccelerationTaskResponse) {
	response = &AllotDatasetAccelerationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
