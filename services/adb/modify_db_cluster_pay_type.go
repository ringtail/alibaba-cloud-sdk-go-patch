package adb

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

// ModifyDBClusterPayType invokes the adb.ModifyDBClusterPayType API synchronously
func (client *Client) ModifyDBClusterPayType(request *ModifyDBClusterPayTypeRequest) (response *ModifyDBClusterPayTypeResponse, err error) {
	response = CreateModifyDBClusterPayTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterPayTypeWithChan invokes the adb.ModifyDBClusterPayType API asynchronously
func (client *Client) ModifyDBClusterPayTypeWithChan(request *ModifyDBClusterPayTypeRequest) (<-chan *ModifyDBClusterPayTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterPayTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterPayType(request)
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

// ModifyDBClusterPayTypeWithCallback invokes the adb.ModifyDBClusterPayType API asynchronously
func (client *Client) ModifyDBClusterPayTypeWithCallback(request *ModifyDBClusterPayTypeRequest, callback func(response *ModifyDBClusterPayTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterPayTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterPayType(request)
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

// ModifyDBClusterPayTypeRequest is the request struct for api ModifyDBClusterPayType
type ModifyDBClusterPayTypeRequest struct {
	*requests.RpcRequest
	Period      string `position:"Query" name:"Period"`
	DbClusterId string `position:"Query" name:"DbClusterId"`
	UsedTime    string `position:"Query" name:"UsedTime"`
	PayType     string `position:"Query" name:"PayType"`
}

// ModifyDBClusterPayTypeResponse is the response struct for api ModifyDBClusterPayType
type ModifyDBClusterPayTypeResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	OrderId     string `json:"OrderId" xml:"OrderId"`
	PayType     string `json:"PayType" xml:"PayType"`
}

// CreateModifyDBClusterPayTypeRequest creates a request to invoke ModifyDBClusterPayType API
func CreateModifyDBClusterPayTypeRequest() (request *ModifyDBClusterPayTypeRequest) {
	request = &ModifyDBClusterPayTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ModifyDBClusterPayType", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterPayTypeResponse creates a response to parse from ModifyDBClusterPayType response
func CreateModifyDBClusterPayTypeResponse() (response *ModifyDBClusterPayTypeResponse) {
	response = &ModifyDBClusterPayTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
