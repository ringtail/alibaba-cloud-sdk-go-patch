package vs

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

// DescribeDirectory invokes the vs.DescribeDirectory API synchronously
func (client *Client) DescribeDirectory(request *DescribeDirectoryRequest) (response *DescribeDirectoryResponse, err error) {
	response = CreateDescribeDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDirectoryWithChan invokes the vs.DescribeDirectory API asynchronously
func (client *Client) DescribeDirectoryWithChan(request *DescribeDirectoryRequest) (<-chan *DescribeDirectoryResponse, <-chan error) {
	responseChan := make(chan *DescribeDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDirectory(request)
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

// DescribeDirectoryWithCallback invokes the vs.DescribeDirectory API asynchronously
func (client *Client) DescribeDirectoryWithCallback(request *DescribeDirectoryRequest, callback func(response *DescribeDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDirectoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeDirectory(request)
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

// DescribeDirectoryRequest is the request struct for api DescribeDirectory
type DescribeDirectoryRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDirectoryResponse is the response struct for api DescribeDirectory
type DescribeDirectoryResponse struct {
	*responses.BaseResponse
	ParentId    string `json:"ParentId" xml:"ParentId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Description string `json:"Description" xml:"Description"`
	GroupId     string `json:"GroupId" xml:"GroupId"`
	Name        string `json:"Name" xml:"Name"`
	CreatedTime string `json:"CreatedTime" xml:"CreatedTime"`
	Id          string `json:"Id" xml:"Id"`
}

// CreateDescribeDirectoryRequest creates a request to invoke DescribeDirectory API
func CreateDescribeDirectoryRequest() (request *DescribeDirectoryRequest) {
	request = &DescribeDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeDirectory", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDirectoryResponse creates a response to parse from DescribeDirectory response
func CreateDescribeDirectoryResponse() (response *DescribeDirectoryResponse) {
	response = &DescribeDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
