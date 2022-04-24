package polardb

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

// RevokeAccountPrivilege invokes the polardb.RevokeAccountPrivilege API synchronously
func (client *Client) RevokeAccountPrivilege(request *RevokeAccountPrivilegeRequest) (response *RevokeAccountPrivilegeResponse, err error) {
	response = CreateRevokeAccountPrivilegeResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeAccountPrivilegeWithChan invokes the polardb.RevokeAccountPrivilege API asynchronously
func (client *Client) RevokeAccountPrivilegeWithChan(request *RevokeAccountPrivilegeRequest) (<-chan *RevokeAccountPrivilegeResponse, <-chan error) {
	responseChan := make(chan *RevokeAccountPrivilegeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeAccountPrivilege(request)
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

// RevokeAccountPrivilegeWithCallback invokes the polardb.RevokeAccountPrivilege API asynchronously
func (client *Client) RevokeAccountPrivilegeWithCallback(request *RevokeAccountPrivilegeRequest, callback func(response *RevokeAccountPrivilegeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeAccountPrivilegeResponse
		var err error
		defer close(result)
		response, err = client.RevokeAccountPrivilege(request)
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

// RevokeAccountPrivilegeRequest is the request struct for api RevokeAccountPrivilege
type RevokeAccountPrivilegeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBName               string           `position:"Query" name:"DBName"`
}

// RevokeAccountPrivilegeResponse is the response struct for api RevokeAccountPrivilege
type RevokeAccountPrivilegeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevokeAccountPrivilegeRequest creates a request to invoke RevokeAccountPrivilege API
func CreateRevokeAccountPrivilegeRequest() (request *RevokeAccountPrivilegeRequest) {
	request = &RevokeAccountPrivilegeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "RevokeAccountPrivilege", "", "")
	request.Method = requests.POST
	return
}

// CreateRevokeAccountPrivilegeResponse creates a response to parse from RevokeAccountPrivilege response
func CreateRevokeAccountPrivilegeResponse() (response *RevokeAccountPrivilegeResponse) {
	response = &RevokeAccountPrivilegeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
