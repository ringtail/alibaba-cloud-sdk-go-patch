package das

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

// GetFullRequestOriginStatByInstanceId invokes the das.GetFullRequestOriginStatByInstanceId API synchronously
func (client *Client) GetFullRequestOriginStatByInstanceId(request *GetFullRequestOriginStatByInstanceIdRequest) (response *GetFullRequestOriginStatByInstanceIdResponse, err error) {
	response = CreateGetFullRequestOriginStatByInstanceIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetFullRequestOriginStatByInstanceIdWithChan invokes the das.GetFullRequestOriginStatByInstanceId API asynchronously
func (client *Client) GetFullRequestOriginStatByInstanceIdWithChan(request *GetFullRequestOriginStatByInstanceIdRequest) (<-chan *GetFullRequestOriginStatByInstanceIdResponse, <-chan error) {
	responseChan := make(chan *GetFullRequestOriginStatByInstanceIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFullRequestOriginStatByInstanceId(request)
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

// GetFullRequestOriginStatByInstanceIdWithCallback invokes the das.GetFullRequestOriginStatByInstanceId API asynchronously
func (client *Client) GetFullRequestOriginStatByInstanceIdWithCallback(request *GetFullRequestOriginStatByInstanceIdRequest, callback func(response *GetFullRequestOriginStatByInstanceIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFullRequestOriginStatByInstanceIdResponse
		var err error
		defer close(result)
		response, err = client.GetFullRequestOriginStatByInstanceId(request)
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

// GetFullRequestOriginStatByInstanceIdRequest is the request struct for api GetFullRequestOriginStatByInstanceId
type GetFullRequestOriginStatByInstanceIdRequest struct {
	*requests.RpcRequest
	UserId     string           `position:"Query" name:"UserId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	NodeId     string           `position:"Query" name:"NodeId"`
	Start      requests.Integer `position:"Query" name:"Start"`
	End        requests.Integer `position:"Query" name:"End"`
	OrderBy    string           `position:"Query" name:"OrderBy"`
	Asc        requests.Boolean `position:"Query" name:"Asc"`
	PageNo     requests.Integer `position:"Query" name:"PageNo"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	SqlType    string           `position:"Query" name:"SqlType"`
	Role       string           `position:"Query" name:"Role"`
}

// GetFullRequestOriginStatByInstanceIdResponse is the response struct for api GetFullRequestOriginStatByInstanceId
type GetFullRequestOriginStatByInstanceIdResponse struct {
	*responses.BaseResponse
	Message   string                                     `json:"Message" xml:"Message"`
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Code      int64                                      `json:"Code" xml:"Code"`
	Success   bool                                       `json:"Success" xml:"Success"`
	Data      DataInGetFullRequestOriginStatByInstanceId `json:"Data" xml:"Data"`
}

// CreateGetFullRequestOriginStatByInstanceIdRequest creates a request to invoke GetFullRequestOriginStatByInstanceId API
func CreateGetFullRequestOriginStatByInstanceIdRequest() (request *GetFullRequestOriginStatByInstanceIdRequest) {
	request = &GetFullRequestOriginStatByInstanceIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetFullRequestOriginStatByInstanceId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetFullRequestOriginStatByInstanceIdResponse creates a response to parse from GetFullRequestOriginStatByInstanceId response
func CreateGetFullRequestOriginStatByInstanceIdResponse() (response *GetFullRequestOriginStatByInstanceIdResponse) {
	response = &GetFullRequestOriginStatByInstanceIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
