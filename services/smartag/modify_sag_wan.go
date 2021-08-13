package smartag

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

// ModifySagWan invokes the smartag.ModifySagWan API synchronously
func (client *Client) ModifySagWan(request *ModifySagWanRequest) (response *ModifySagWanResponse, err error) {
	response = CreateModifySagWanResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySagWanWithChan invokes the smartag.ModifySagWan API asynchronously
func (client *Client) ModifySagWanWithChan(request *ModifySagWanRequest) (<-chan *ModifySagWanResponse, <-chan error) {
	responseChan := make(chan *ModifySagWanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySagWan(request)
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

// ModifySagWanWithCallback invokes the smartag.ModifySagWan API asynchronously
func (client *Client) ModifySagWanWithCallback(request *ModifySagWanRequest, callback func(response *ModifySagWanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySagWanResponse
		var err error
		defer close(result)
		response, err = client.ModifySagWan(request)
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

// ModifySagWanRequest is the request struct for api ModifySagWan
type ModifySagWanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ISP                  string           `position:"Query" name:"ISP"`
	Password             string           `position:"Query" name:"Password"`
	Vlan                 string           `position:"Query" name:"Vlan"`
	Mask                 string           `position:"Query" name:"Mask"`
	StartIp              string           `position:"Query" name:"StartIp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	IP                   string           `position:"Query" name:"IP"`
	Weight               requests.Integer `position:"Query" name:"Weight"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	IPType               string           `position:"Query" name:"IPType"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	SourceIps            string           `position:"Query" name:"SourceIps"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
	PortName             string           `position:"Query" name:"PortName"`
	StopIp               string           `position:"Query" name:"StopIp"`
	Gateway              string           `position:"Query" name:"Gateway"`
	Username             string           `position:"Query" name:"Username"`
}

// ModifySagWanResponse is the response struct for api ModifySagWan
type ModifySagWanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySagWanRequest creates a request to invoke ModifySagWan API
func CreateModifySagWanRequest() (request *ModifySagWanRequest) {
	request = &ModifySagWanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySagWan", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySagWanResponse creates a response to parse from ModifySagWan response
func CreateModifySagWanResponse() (response *ModifySagWanResponse) {
	response = &ModifySagWanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
