package cloudfw

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

// ModifyVpcFirewallConfigure invokes the cloudfw.ModifyVpcFirewallConfigure API synchronously
func (client *Client) ModifyVpcFirewallConfigure(request *ModifyVpcFirewallConfigureRequest) (response *ModifyVpcFirewallConfigureResponse, err error) {
	response = CreateModifyVpcFirewallConfigureResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcFirewallConfigureWithChan invokes the cloudfw.ModifyVpcFirewallConfigure API asynchronously
func (client *Client) ModifyVpcFirewallConfigureWithChan(request *ModifyVpcFirewallConfigureRequest) (<-chan *ModifyVpcFirewallConfigureResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcFirewallConfigureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcFirewallConfigure(request)
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

// ModifyVpcFirewallConfigureWithCallback invokes the cloudfw.ModifyVpcFirewallConfigure API asynchronously
func (client *Client) ModifyVpcFirewallConfigureWithCallback(request *ModifyVpcFirewallConfigureRequest, callback func(response *ModifyVpcFirewallConfigureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcFirewallConfigureResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcFirewallConfigure(request)
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

// ModifyVpcFirewallConfigureRequest is the request struct for api ModifyVpcFirewallConfigure
type ModifyVpcFirewallConfigureRequest struct {
	*requests.RpcRequest
	PeerVpcCidrTableList  string `position:"Query" name:"PeerVpcCidrTableList"`
	LocalVpcCidrTableList string `position:"Query" name:"LocalVpcCidrTableList"`
	VpcFirewallName       string `position:"Query" name:"VpcFirewallName"`
	SourceIp              string `position:"Query" name:"SourceIp"`
	Lang                  string `position:"Query" name:"Lang"`
	VpcFirewallId         string `position:"Query" name:"VpcFirewallId"`
	MemberUid             string `position:"Query" name:"MemberUid"`
}

// ModifyVpcFirewallConfigureResponse is the response struct for api ModifyVpcFirewallConfigure
type ModifyVpcFirewallConfigureResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcFirewallConfigureRequest creates a request to invoke ModifyVpcFirewallConfigure API
func CreateModifyVpcFirewallConfigureRequest() (request *ModifyVpcFirewallConfigureRequest) {
	request = &ModifyVpcFirewallConfigureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "ModifyVpcFirewallConfigure", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpcFirewallConfigureResponse creates a response to parse from ModifyVpcFirewallConfigure response
func CreateModifyVpcFirewallConfigureResponse() (response *ModifyVpcFirewallConfigureResponse) {
	response = &ModifyVpcFirewallConfigureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
