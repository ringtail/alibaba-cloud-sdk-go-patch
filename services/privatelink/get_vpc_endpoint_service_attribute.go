package privatelink

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

// GetVpcEndpointServiceAttribute invokes the privatelink.GetVpcEndpointServiceAttribute API synchronously
func (client *Client) GetVpcEndpointServiceAttribute(request *GetVpcEndpointServiceAttributeRequest) (response *GetVpcEndpointServiceAttributeResponse, err error) {
	response = CreateGetVpcEndpointServiceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// GetVpcEndpointServiceAttributeWithChan invokes the privatelink.GetVpcEndpointServiceAttribute API asynchronously
func (client *Client) GetVpcEndpointServiceAttributeWithChan(request *GetVpcEndpointServiceAttributeRequest) (<-chan *GetVpcEndpointServiceAttributeResponse, <-chan error) {
	responseChan := make(chan *GetVpcEndpointServiceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVpcEndpointServiceAttribute(request)
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

// GetVpcEndpointServiceAttributeWithCallback invokes the privatelink.GetVpcEndpointServiceAttribute API asynchronously
func (client *Client) GetVpcEndpointServiceAttributeWithCallback(request *GetVpcEndpointServiceAttributeRequest, callback func(response *GetVpcEndpointServiceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVpcEndpointServiceAttributeResponse
		var err error
		defer close(result)
		response, err = client.GetVpcEndpointServiceAttribute(request)
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

// GetVpcEndpointServiceAttributeRequest is the request struct for api GetVpcEndpointServiceAttribute
type GetVpcEndpointServiceAttributeRequest struct {
	*requests.RpcRequest
	ServiceId string `position:"Query" name:"ServiceId"`
}

// GetVpcEndpointServiceAttributeResponse is the response struct for api GetVpcEndpointServiceAttribute
type GetVpcEndpointServiceAttributeResponse struct {
	*responses.BaseResponse
	Payer                              string   `json:"Payer" xml:"Payer"`
	RequestId                          string   `json:"RequestId" xml:"RequestId"`
	ServiceDescription                 string   `json:"ServiceDescription" xml:"ServiceDescription"`
	CreateTime                         string   `json:"CreateTime" xml:"CreateTime"`
	MaxBandwidth                       int      `json:"MaxBandwidth" xml:"MaxBandwidth"`
	MinBandwidth                       int      `json:"MinBandwidth" xml:"MinBandwidth"`
	ServiceDomain                      string   `json:"ServiceDomain" xml:"ServiceDomain"`
	AutoAcceptEnabled                  bool     `json:"AutoAcceptEnabled" xml:"AutoAcceptEnabled"`
	ZoneAffinityEnabled                bool     `json:"ZoneAffinityEnabled" xml:"ZoneAffinityEnabled"`
	ServiceId                          string   `json:"ServiceId" xml:"ServiceId"`
	ServiceBusinessStatus              string   `json:"ServiceBusinessStatus" xml:"ServiceBusinessStatus"`
	ServiceName                        string   `json:"ServiceName" xml:"ServiceName"`
	ServiceStatus                      string   `json:"ServiceStatus" xml:"ServiceStatus"`
	ConnectBandwidth                   int      `json:"ConnectBandwidth" xml:"ConnectBandwidth"`
	RegionId                           string   `json:"RegionId" xml:"RegionId"`
	ServiceType                        string   `json:"ServiceType" xml:"ServiceType"`
	ServiceResourceType                string   `json:"ServiceResourceType" xml:"ServiceResourceType"`
	PrivateServiceDomainEnabled        bool     `json:"PrivateServiceDomainEnabled" xml:"PrivateServiceDomainEnabled"`
	PrivateServiceDomain               string   `json:"PrivateServiceDomain" xml:"PrivateServiceDomain"`
	PrivateServiceDomainVerifyStatus   string   `json:"PrivateServiceDomainVerifyStatus" xml:"PrivateServiceDomainVerifyStatus"`
	PrivateServiceDomainBusinessStatus string   `json:"PrivateServiceDomainBusinessStatus" xml:"PrivateServiceDomainBusinessStatus"`
	PrivateServiceDomainVerifyName     string   `json:"PrivateServiceDomainVerifyName" xml:"PrivateServiceDomainVerifyName"`
	PrivateServiceDomainVerifyValue    string   `json:"PrivateServiceDomainVerifyValue" xml:"PrivateServiceDomainVerifyValue"`
	PrivateServiceName                 string   `json:"PrivateServiceName" xml:"PrivateServiceName"`
	ServiceSupportIPv6                 bool     `json:"ServiceSupportIPv6" xml:"ServiceSupportIPv6"`
	ResourceGroupId                    string   `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Zones                              []string `json:"Zones" xml:"Zones"`
}

// CreateGetVpcEndpointServiceAttributeRequest creates a request to invoke GetVpcEndpointServiceAttribute API
func CreateGetVpcEndpointServiceAttributeRequest() (request *GetVpcEndpointServiceAttributeRequest) {
	request = &GetVpcEndpointServiceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "GetVpcEndpointServiceAttribute", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVpcEndpointServiceAttributeResponse creates a response to parse from GetVpcEndpointServiceAttribute response
func CreateGetVpcEndpointServiceAttributeResponse() (response *GetVpcEndpointServiceAttributeResponse) {
	response = &GetVpcEndpointServiceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
