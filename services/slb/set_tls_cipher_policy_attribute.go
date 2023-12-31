package slb

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

// SetTLSCipherPolicyAttribute invokes the slb.SetTLSCipherPolicyAttribute API synchronously
func (client *Client) SetTLSCipherPolicyAttribute(request *SetTLSCipherPolicyAttributeRequest) (response *SetTLSCipherPolicyAttributeResponse, err error) {
	response = CreateSetTLSCipherPolicyAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// SetTLSCipherPolicyAttributeWithChan invokes the slb.SetTLSCipherPolicyAttribute API asynchronously
func (client *Client) SetTLSCipherPolicyAttributeWithChan(request *SetTLSCipherPolicyAttributeRequest) (<-chan *SetTLSCipherPolicyAttributeResponse, <-chan error) {
	responseChan := make(chan *SetTLSCipherPolicyAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetTLSCipherPolicyAttribute(request)
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

// SetTLSCipherPolicyAttributeWithCallback invokes the slb.SetTLSCipherPolicyAttribute API asynchronously
func (client *Client) SetTLSCipherPolicyAttributeWithCallback(request *SetTLSCipherPolicyAttributeRequest, callback func(response *SetTLSCipherPolicyAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetTLSCipherPolicyAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetTLSCipherPolicyAttribute(request)
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

// SetTLSCipherPolicyAttributeRequest is the request struct for api SetTLSCipherPolicyAttribute
type SetTLSCipherPolicyAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TLSCipherPolicyId    string           `position:"Query" name:"TLSCipherPolicyId"`
	Ciphers              *[]string        `position:"Query" name:"Ciphers"  type:"Repeated"`
	TLSVersions          *[]string        `position:"Query" name:"TLSVersions"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// SetTLSCipherPolicyAttributeResponse is the response struct for api SetTLSCipherPolicyAttribute
type SetTLSCipherPolicyAttributeResponse struct {
	*responses.BaseResponse
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetTLSCipherPolicyAttributeRequest creates a request to invoke SetTLSCipherPolicyAttribute API
func CreateSetTLSCipherPolicyAttributeRequest() (request *SetTLSCipherPolicyAttributeRequest) {
	request = &SetTLSCipherPolicyAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetTLSCipherPolicyAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetTLSCipherPolicyAttributeResponse creates a response to parse from SetTLSCipherPolicyAttribute response
func CreateSetTLSCipherPolicyAttributeResponse() (response *SetTLSCipherPolicyAttributeResponse) {
	response = &SetTLSCipherPolicyAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
