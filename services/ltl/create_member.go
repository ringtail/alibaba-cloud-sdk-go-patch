package ltl

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

// CreateMember invokes the ltl.CreateMember API synchronously
func (client *Client) CreateMember(request *CreateMemberRequest) (response *CreateMemberResponse, err error) {
	response = CreateCreateMemberResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMemberWithChan invokes the ltl.CreateMember API asynchronously
func (client *Client) CreateMemberWithChan(request *CreateMemberRequest) (<-chan *CreateMemberResponse, <-chan error) {
	responseChan := make(chan *CreateMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMember(request)
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

// CreateMemberWithCallback invokes the ltl.CreateMember API asynchronously
func (client *Client) CreateMemberWithCallback(request *CreateMemberRequest, callback func(response *CreateMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMemberResponse
		var err error
		defer close(result)
		response, err = client.CreateMember(request)
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

// CreateMemberRequest is the request struct for api CreateMember
type CreateMemberRequest struct {
	*requests.RpcRequest
	ApiVersion    string `position:"Query" name:"ApiVersion"`
	Remark        string `position:"Query" name:"Remark"`
	BizChainId    string `position:"Query" name:"BizChainId"`
	MemberUid     string `position:"Query" name:"MemberUid"`
	MemberContact string `position:"Query" name:"MemberContact"`
	MemberPhone   string `position:"Query" name:"MemberPhone"`
	MemberName    string `position:"Query" name:"MemberName"`
}

// CreateMemberResponse is the response struct for api CreateMember
type CreateMemberResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateMemberRequest creates a request to invoke CreateMember API
func CreateCreateMemberRequest() (request *CreateMemberRequest) {
	request = &CreateMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "CreateMember", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMemberResponse creates a response to parse from CreateMember response
func CreateCreateMemberResponse() (response *CreateMemberResponse) {
	response = &CreateMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
