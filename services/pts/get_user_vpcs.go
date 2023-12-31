package pts

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

// GetUserVpcs invokes the pts.GetUserVpcs API synchronously
func (client *Client) GetUserVpcs(request *GetUserVpcsRequest) (response *GetUserVpcsResponse, err error) {
	response = CreateGetUserVpcsResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserVpcsWithChan invokes the pts.GetUserVpcs API asynchronously
func (client *Client) GetUserVpcsWithChan(request *GetUserVpcsRequest) (<-chan *GetUserVpcsResponse, <-chan error) {
	responseChan := make(chan *GetUserVpcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserVpcs(request)
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

// GetUserVpcsWithCallback invokes the pts.GetUserVpcs API asynchronously
func (client *Client) GetUserVpcsWithCallback(request *GetUserVpcsRequest, callback func(response *GetUserVpcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserVpcsResponse
		var err error
		defer close(result)
		response, err = client.GetUserVpcs(request)
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

// GetUserVpcsRequest is the request struct for api GetUserVpcs
type GetUserVpcsRequest struct {
	*requests.RpcRequest
	VpcId      string           `position:"Query" name:"VpcId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// GetUserVpcsResponse is the response struct for api GetUserVpcs
type GetUserVpcsResponse struct {
	*responses.BaseResponse
	TotalCount     int64  `json:"TotalCount" xml:"TotalCount"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	PageSize       int    `json:"PageSize" xml:"PageSize"`
	PageNumber     int    `json:"PageNumber" xml:"PageNumber"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Vpcs           []Vpc  `json:"Vpcs" xml:"Vpcs"`
}

// CreateGetUserVpcsRequest creates a request to invoke GetUserVpcs API
func CreateGetUserVpcsRequest() (request *GetUserVpcsRequest) {
	request = &GetUserVpcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetUserVpcs", "", "")
	request.Method = requests.POST
	return
}

// CreateGetUserVpcsResponse creates a response to parse from GetUserVpcs response
func CreateGetUserVpcsResponse() (response *GetUserVpcsResponse) {
	response = &GetUserVpcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
