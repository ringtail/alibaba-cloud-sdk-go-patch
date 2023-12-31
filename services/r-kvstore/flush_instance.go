package r_kvstore

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

// FlushInstance invokes the r_kvstore.FlushInstance API synchronously
func (client *Client) FlushInstance(request *FlushInstanceRequest) (response *FlushInstanceResponse, err error) {
	response = CreateFlushInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// FlushInstanceWithChan invokes the r_kvstore.FlushInstance API asynchronously
func (client *Client) FlushInstanceWithChan(request *FlushInstanceRequest) (<-chan *FlushInstanceResponse, <-chan error) {
	responseChan := make(chan *FlushInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FlushInstance(request)
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

// FlushInstanceWithCallback invokes the r_kvstore.FlushInstance API asynchronously
func (client *Client) FlushInstanceWithCallback(request *FlushInstanceRequest, callback func(response *FlushInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FlushInstanceResponse
		var err error
		defer close(result)
		response, err = client.FlushInstance(request)
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

// FlushInstanceRequest is the request struct for api FlushInstance
type FlushInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// FlushInstanceResponse is the response struct for api FlushInstance
type FlushInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateFlushInstanceRequest creates a request to invoke FlushInstance API
func CreateFlushInstanceRequest() (request *FlushInstanceRequest) {
	request = &FlushInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "FlushInstance", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFlushInstanceResponse creates a response to parse from FlushInstance response
func CreateFlushInstanceResponse() (response *FlushInstanceResponse) {
	response = &FlushInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
