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

// ModifyCreateVulWhitelist invokes the sas.ModifyCreateVulWhitelist API synchronously
func (client *Client) ModifyCreateVulWhitelist(request *ModifyCreateVulWhitelistRequest) (response *ModifyCreateVulWhitelistResponse, err error) {
	response = CreateModifyCreateVulWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCreateVulWhitelistWithChan invokes the sas.ModifyCreateVulWhitelist API asynchronously
func (client *Client) ModifyCreateVulWhitelistWithChan(request *ModifyCreateVulWhitelistRequest) (<-chan *ModifyCreateVulWhitelistResponse, <-chan error) {
	responseChan := make(chan *ModifyCreateVulWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCreateVulWhitelist(request)
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

// ModifyCreateVulWhitelistWithCallback invokes the sas.ModifyCreateVulWhitelist API asynchronously
func (client *Client) ModifyCreateVulWhitelistWithCallback(request *ModifyCreateVulWhitelistRequest, callback func(response *ModifyCreateVulWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCreateVulWhitelistResponse
		var err error
		defer close(result)
		response, err = client.ModifyCreateVulWhitelist(request)
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

// ModifyCreateVulWhitelistRequest is the request struct for api ModifyCreateVulWhitelist
type ModifyCreateVulWhitelistRequest struct {
	*requests.RpcRequest
	Reason     string `position:"Query" name:"Reason"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Whitelist  string `position:"Query" name:"Whitelist"`
	TargetInfo string `position:"Query" name:"TargetInfo"`
}

// ModifyCreateVulWhitelistResponse is the response struct for api ModifyCreateVulWhitelist
type ModifyCreateVulWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCreateVulWhitelistRequest creates a request to invoke ModifyCreateVulWhitelist API
func CreateModifyCreateVulWhitelistRequest() (request *ModifyCreateVulWhitelistRequest) {
	request = &ModifyCreateVulWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyCreateVulWhitelist", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCreateVulWhitelistResponse creates a response to parse from ModifyCreateVulWhitelist response
func CreateModifyCreateVulWhitelistResponse() (response *ModifyCreateVulWhitelistResponse) {
	response = &ModifyCreateVulWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
