package opensearch

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

// GetFunctionDefaultInstance invokes the opensearch.GetFunctionDefaultInstance API synchronously
func (client *Client) GetFunctionDefaultInstance(request *GetFunctionDefaultInstanceRequest) (response *GetFunctionDefaultInstanceResponse, err error) {
	response = CreateGetFunctionDefaultInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetFunctionDefaultInstanceWithChan invokes the opensearch.GetFunctionDefaultInstance API asynchronously
func (client *Client) GetFunctionDefaultInstanceWithChan(request *GetFunctionDefaultInstanceRequest) (<-chan *GetFunctionDefaultInstanceResponse, <-chan error) {
	responseChan := make(chan *GetFunctionDefaultInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFunctionDefaultInstance(request)
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

// GetFunctionDefaultInstanceWithCallback invokes the opensearch.GetFunctionDefaultInstance API asynchronously
func (client *Client) GetFunctionDefaultInstanceWithCallback(request *GetFunctionDefaultInstanceRequest, callback func(response *GetFunctionDefaultInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFunctionDefaultInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetFunctionDefaultInstance(request)
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

// GetFunctionDefaultInstanceRequest is the request struct for api GetFunctionDefaultInstance
type GetFunctionDefaultInstanceRequest struct {
	*requests.RoaRequest
	FunctionName     string `position:"Path" name:"functionName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// GetFunctionDefaultInstanceResponse is the response struct for api GetFunctionDefaultInstance
type GetFunctionDefaultInstanceResponse struct {
	*responses.BaseResponse
	Status       string `json:"Status" xml:"Status"`
	HttpCode     int64  `json:"HttpCode" xml:"HttpCode"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Message      string `json:"Message" xml:"Message"`
	Code         string `json:"Code" xml:"Code"`
	Latency      int64  `json:"Latency" xml:"Latency"`
	InstanceName string `json:"InstanceName" xml:"InstanceName"`
	FunctionName string `json:"FunctionName" xml:"FunctionName"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateGetFunctionDefaultInstanceRequest creates a request to invoke GetFunctionDefaultInstance API
func CreateGetFunctionDefaultInstanceRequest() (request *GetFunctionDefaultInstanceRequest) {
	request = &GetFunctionDefaultInstanceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetFunctionDefaultInstance", "/v4/openapi/app-groups/[appGroupIdentity]/functions/[functionName]/default-instance", "", "")
	request.Method = requests.GET
	return
}

// CreateGetFunctionDefaultInstanceResponse creates a response to parse from GetFunctionDefaultInstance response
func CreateGetFunctionDefaultInstanceResponse() (response *GetFunctionDefaultInstanceResponse) {
	response = &GetFunctionDefaultInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
