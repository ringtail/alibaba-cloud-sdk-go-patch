package adb

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

// DescribeDBClusterAccessWhiteList invokes the adb.DescribeDBClusterAccessWhiteList API synchronously
// api document: https://help.aliyun.com/api/adb/describedbclusteraccesswhitelist.html
func (client *Client) DescribeDBClusterAccessWhiteList(request *DescribeDBClusterAccessWhiteListRequest) (response *DescribeDBClusterAccessWhiteListResponse, err error) {
	response = CreateDescribeDBClusterAccessWhiteListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterAccessWhiteListWithChan invokes the adb.DescribeDBClusterAccessWhiteList API asynchronously
// api document: https://help.aliyun.com/api/adb/describedbclusteraccesswhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterAccessWhiteListWithChan(request *DescribeDBClusterAccessWhiteListRequest) (<-chan *DescribeDBClusterAccessWhiteListResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterAccessWhiteListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterAccessWhiteList(request)
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

// DescribeDBClusterAccessWhiteListWithCallback invokes the adb.DescribeDBClusterAccessWhiteList API asynchronously
// api document: https://help.aliyun.com/api/adb/describedbclusteraccesswhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterAccessWhiteListWithCallback(request *DescribeDBClusterAccessWhiteListRequest, callback func(response *DescribeDBClusterAccessWhiteListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterAccessWhiteListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterAccessWhiteList(request)
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

// DescribeDBClusterAccessWhiteListRequest is the request struct for api DescribeDBClusterAccessWhiteList
type DescribeDBClusterAccessWhiteListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterAccessWhiteListResponse is the response struct for api DescribeDBClusterAccessWhiteList
type DescribeDBClusterAccessWhiteListResponse struct {
	*responses.BaseResponse
	RequestId string                                  `json:"RequestId" xml:"RequestId"`
	Items     ItemsInDescribeDBClusterAccessWhiteList `json:"Items" xml:"Items"`
}

// CreateDescribeDBClusterAccessWhiteListRequest creates a request to invoke DescribeDBClusterAccessWhiteList API
func CreateDescribeDBClusterAccessWhiteListRequest() (request *DescribeDBClusterAccessWhiteListRequest) {
	request = &DescribeDBClusterAccessWhiteListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeDBClusterAccessWhiteList", "ads", "openAPI")
	return
}

// CreateDescribeDBClusterAccessWhiteListResponse creates a response to parse from DescribeDBClusterAccessWhiteList response
func CreateDescribeDBClusterAccessWhiteListResponse() (response *DescribeDBClusterAccessWhiteListResponse) {
	response = &DescribeDBClusterAccessWhiteListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
