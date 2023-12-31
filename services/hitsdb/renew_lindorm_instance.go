package hitsdb

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

// RenewLindormInstance invokes the hitsdb.RenewLindormInstance API synchronously
func (client *Client) RenewLindormInstance(request *RenewLindormInstanceRequest) (response *RenewLindormInstanceResponse, err error) {
	response = CreateRenewLindormInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RenewLindormInstanceWithChan invokes the hitsdb.RenewLindormInstance API asynchronously
func (client *Client) RenewLindormInstanceWithChan(request *RenewLindormInstanceRequest) (<-chan *RenewLindormInstanceResponse, <-chan error) {
	responseChan := make(chan *RenewLindormInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewLindormInstance(request)
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

// RenewLindormInstanceWithCallback invokes the hitsdb.RenewLindormInstance API asynchronously
func (client *Client) RenewLindormInstanceWithCallback(request *RenewLindormInstanceRequest, callback func(response *RenewLindormInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewLindormInstanceResponse
		var err error
		defer close(result)
		response, err = client.RenewLindormInstance(request)
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

// RenewLindormInstanceRequest is the request struct for api RenewLindormInstance
type RenewLindormInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Duration             requests.Integer `position:"Query" name:"Duration"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
}

// RenewLindormInstanceResponse is the response struct for api RenewLindormInstance
type RenewLindormInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	OrderId    int64  `json:"OrderId" xml:"OrderId"`
}

// CreateRenewLindormInstanceRequest creates a request to invoke RenewLindormInstance API
func CreateRenewLindormInstanceRequest() (request *RenewLindormInstanceRequest) {
	request = &RenewLindormInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "RenewLindormInstance", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewLindormInstanceResponse creates a response to parse from RenewLindormInstance response
func CreateRenewLindormInstanceResponse() (response *RenewLindormInstanceResponse) {
	response = &RenewLindormInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
