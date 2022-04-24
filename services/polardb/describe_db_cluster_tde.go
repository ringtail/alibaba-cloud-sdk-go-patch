package polardb

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

// DescribeDBClusterTDE invokes the polardb.DescribeDBClusterTDE API synchronously
func (client *Client) DescribeDBClusterTDE(request *DescribeDBClusterTDERequest) (response *DescribeDBClusterTDEResponse, err error) {
	response = CreateDescribeDBClusterTDEResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterTDEWithChan invokes the polardb.DescribeDBClusterTDE API asynchronously
func (client *Client) DescribeDBClusterTDEWithChan(request *DescribeDBClusterTDERequest) (<-chan *DescribeDBClusterTDEResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterTDEResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterTDE(request)
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

// DescribeDBClusterTDEWithCallback invokes the polardb.DescribeDBClusterTDE API asynchronously
func (client *Client) DescribeDBClusterTDEWithCallback(request *DescribeDBClusterTDERequest, callback func(response *DescribeDBClusterTDEResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterTDEResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterTDE(request)
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

// DescribeDBClusterTDERequest is the request struct for api DescribeDBClusterTDE
type DescribeDBClusterTDERequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterTDEResponse is the response struct for api DescribeDBClusterTDE
type DescribeDBClusterTDEResponse struct {
	*responses.BaseResponse
	TDEStatus        string `json:"TDEStatus" xml:"TDEStatus"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBClusterId      string `json:"DBClusterId" xml:"DBClusterId"`
	EncryptionKey    string `json:"EncryptionKey" xml:"EncryptionKey"`
	EncryptNewTables string `json:"EncryptNewTables" xml:"EncryptNewTables"`
}

// CreateDescribeDBClusterTDERequest creates a request to invoke DescribeDBClusterTDE API
func CreateDescribeDBClusterTDERequest() (request *DescribeDBClusterTDERequest) {
	request = &DescribeDBClusterTDERequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterTDE", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterTDEResponse creates a response to parse from DescribeDBClusterTDE response
func CreateDescribeDBClusterTDEResponse() (response *DescribeDBClusterTDEResponse) {
	response = &DescribeDBClusterTDEResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
