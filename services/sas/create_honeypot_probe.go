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

// CreateHoneypotProbe invokes the sas.CreateHoneypotProbe API synchronously
func (client *Client) CreateHoneypotProbe(request *CreateHoneypotProbeRequest) (response *CreateHoneypotProbeResponse, err error) {
	response = CreateCreateHoneypotProbeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateHoneypotProbeWithChan invokes the sas.CreateHoneypotProbe API asynchronously
func (client *Client) CreateHoneypotProbeWithChan(request *CreateHoneypotProbeRequest) (<-chan *CreateHoneypotProbeResponse, <-chan error) {
	responseChan := make(chan *CreateHoneypotProbeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateHoneypotProbe(request)
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

// CreateHoneypotProbeWithCallback invokes the sas.CreateHoneypotProbe API asynchronously
func (client *Client) CreateHoneypotProbeWithCallback(request *CreateHoneypotProbeRequest, callback func(response *CreateHoneypotProbeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateHoneypotProbeResponse
		var err error
		defer close(result)
		response, err = client.CreateHoneypotProbe(request)
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

// CreateHoneypotProbeRequest is the request struct for api CreateHoneypotProbe
type CreateHoneypotProbeRequest struct {
	*requests.RpcRequest
	ControlNodeId    string                                 `position:"Query" name:"ControlNodeId"`
	ProxyIp          string                                 `position:"Query" name:"ProxyIp"`
	Ping             requests.Boolean                       `position:"Query" name:"Ping"`
	ProbeId          string                                 `position:"Query" name:"ProbeId"`
	Uuid             string                                 `position:"Query" name:"Uuid"`
	ProbeVersion     string                                 `position:"Query" name:"ProbeVersion"`
	ServiceIpList    *[]string                              `position:"Query" name:"ServiceIpList"  type:"Repeated"`
	HoneypotBindList *[]CreateHoneypotProbeHoneypotBindList `position:"Query" name:"HoneypotBindList"  type:"Repeated"`
	Lang             string                                 `position:"Query" name:"Lang"`
	Arp              requests.Boolean                       `position:"Query" name:"Arp"`
	ProbeType        string                                 `position:"Query" name:"ProbeType"`
	ProbeStatus      string                                 `position:"Query" name:"ProbeStatus"`
	BusinessGroupId  string                                 `position:"Query" name:"BusinessGroupId"`
	DisplayName      string                                 `position:"Query" name:"DisplayName"`
	VpcId            string                                 `position:"Query" name:"VpcId"`
}

// CreateHoneypotProbeHoneypotBindList is a repeated param struct in CreateHoneypotProbeRequest
type CreateHoneypotProbeHoneypotBindList struct {
	BindPortList *[]CreateHoneypotProbeHoneypotBindListBindPortList `name:"BindPortList" type:"Repeated"`
	HoneypotId   string                                             `name:"HoneypotId"`
}

// CreateHoneypotProbeHoneypotBindListBindPortList is a repeated param struct in CreateHoneypotProbeRequest
type CreateHoneypotProbeHoneypotBindListBindPortList struct {
	StartPort  string `name:"StartPort"`
	BindPort   string `name:"BindPort"`
	Fixed      string `name:"Fixed"`
	EndPort    string `name:"EndPort"`
	TargetPort string `name:"TargetPort"`
}

// CreateHoneypotProbeResponse is the response struct for api CreateHoneypotProbe
type CreateHoneypotProbeResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateHoneypotProbeRequest creates a request to invoke CreateHoneypotProbe API
func CreateCreateHoneypotProbeRequest() (request *CreateHoneypotProbeRequest) {
	request = &CreateHoneypotProbeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "CreateHoneypotProbe", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateHoneypotProbeResponse creates a response to parse from CreateHoneypotProbe response
func CreateCreateHoneypotProbeResponse() (response *CreateHoneypotProbeResponse) {
	response = &CreateHoneypotProbeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
