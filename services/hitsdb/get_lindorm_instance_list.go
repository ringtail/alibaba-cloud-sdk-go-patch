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

// GetLindormInstanceList invokes the hitsdb.GetLindormInstanceList API synchronously
func (client *Client) GetLindormInstanceList(request *GetLindormInstanceListRequest) (response *GetLindormInstanceListResponse, err error) {
	response = CreateGetLindormInstanceListResponse()
	err = client.DoAction(request, response)
	return
}

// GetLindormInstanceListWithChan invokes the hitsdb.GetLindormInstanceList API asynchronously
func (client *Client) GetLindormInstanceListWithChan(request *GetLindormInstanceListRequest) (<-chan *GetLindormInstanceListResponse, <-chan error) {
	responseChan := make(chan *GetLindormInstanceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLindormInstanceList(request)
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

// GetLindormInstanceListWithCallback invokes the hitsdb.GetLindormInstanceList API asynchronously
func (client *Client) GetLindormInstanceListWithCallback(request *GetLindormInstanceListRequest, callback func(response *GetLindormInstanceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLindormInstanceListResponse
		var err error
		defer close(result)
		response, err = client.GetLindormInstanceList(request)
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

// GetLindormInstanceListRequest is the request struct for api GetLindormInstanceList
type GetLindormInstanceListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer             `position:"Query" name:"ResourceOwnerId"`
	SupportEngine        requests.Integer             `position:"Query" name:"SupportEngine"`
	PageNumber           requests.Integer             `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                       `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string                       `position:"Query" name:"SecurityToken"`
	PageSize             requests.Integer             `position:"Query" name:"PageSize"`
	Tag                  *[]GetLindormInstanceListTag `position:"Query" name:"Tag"  type:"Repeated"`
	QueryStr             string                       `position:"Query" name:"QueryStr"`
	ResourceOwnerAccount string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                       `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer             `position:"Query" name:"OwnerId"`
	ServiceType          string                       `position:"Query" name:"ServiceType"`
}

// GetLindormInstanceListTag is a repeated param struct in GetLindormInstanceListRequest
type GetLindormInstanceListTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// GetLindormInstanceListResponse is the response struct for api GetLindormInstanceList
type GetLindormInstanceListResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	PageNumber   int                      `json:"PageNumber" xml:"PageNumber"`
	PageSize     int                      `json:"PageSize" xml:"PageSize"`
	Total        int                      `json:"Total" xml:"Total"`
	InstanceList []LindormInstanceSummary `json:"InstanceList" xml:"InstanceList"`
}

// CreateGetLindormInstanceListRequest creates a request to invoke GetLindormInstanceList API
func CreateGetLindormInstanceListRequest() (request *GetLindormInstanceListRequest) {
	request = &GetLindormInstanceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2020-06-15", "GetLindormInstanceList", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLindormInstanceListResponse creates a response to parse from GetLindormInstanceList response
func CreateGetLindormInstanceListResponse() (response *GetLindormInstanceListResponse) {
	response = &GetLindormInstanceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
