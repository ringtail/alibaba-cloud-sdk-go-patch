package sgw

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

// ExpandGatewayNetworkBandwidth invokes the sgw.ExpandGatewayNetworkBandwidth API synchronously
func (client *Client) ExpandGatewayNetworkBandwidth(request *ExpandGatewayNetworkBandwidthRequest) (response *ExpandGatewayNetworkBandwidthResponse, err error) {
	response = CreateExpandGatewayNetworkBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// ExpandGatewayNetworkBandwidthWithChan invokes the sgw.ExpandGatewayNetworkBandwidth API asynchronously
func (client *Client) ExpandGatewayNetworkBandwidthWithChan(request *ExpandGatewayNetworkBandwidthRequest) (<-chan *ExpandGatewayNetworkBandwidthResponse, <-chan error) {
	responseChan := make(chan *ExpandGatewayNetworkBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExpandGatewayNetworkBandwidth(request)
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

// ExpandGatewayNetworkBandwidthWithCallback invokes the sgw.ExpandGatewayNetworkBandwidth API asynchronously
func (client *Client) ExpandGatewayNetworkBandwidthWithCallback(request *ExpandGatewayNetworkBandwidthRequest, callback func(response *ExpandGatewayNetworkBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExpandGatewayNetworkBandwidthResponse
		var err error
		defer close(result)
		response, err = client.ExpandGatewayNetworkBandwidth(request)
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

// ExpandGatewayNetworkBandwidthRequest is the request struct for api ExpandGatewayNetworkBandwidth
type ExpandGatewayNetworkBandwidthRequest struct {
	*requests.RpcRequest
	NewNetworkBandwidth requests.Integer `position:"Query" name:"NewNetworkBandwidth"`
	SecurityToken       string           `position:"Query" name:"SecurityToken"`
	GatewayId           string           `position:"Query" name:"GatewayId"`
}

// ExpandGatewayNetworkBandwidthResponse is the response struct for api ExpandGatewayNetworkBandwidth
type ExpandGatewayNetworkBandwidthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	BuyURL    string `json:"BuyURL" xml:"BuyURL"`
}

// CreateExpandGatewayNetworkBandwidthRequest creates a request to invoke ExpandGatewayNetworkBandwidth API
func CreateExpandGatewayNetworkBandwidthRequest() (request *ExpandGatewayNetworkBandwidthRequest) {
	request = &ExpandGatewayNetworkBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "ExpandGatewayNetworkBandwidth", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExpandGatewayNetworkBandwidthResponse creates a response to parse from ExpandGatewayNetworkBandwidth response
func CreateExpandGatewayNetworkBandwidthResponse() (response *ExpandGatewayNetworkBandwidthResponse) {
	response = &ExpandGatewayNetworkBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}