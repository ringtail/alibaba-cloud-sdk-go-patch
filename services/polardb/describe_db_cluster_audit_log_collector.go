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

// DescribeDBClusterAuditLogCollector invokes the polardb.DescribeDBClusterAuditLogCollector API synchronously
func (client *Client) DescribeDBClusterAuditLogCollector(request *DescribeDBClusterAuditLogCollectorRequest) (response *DescribeDBClusterAuditLogCollectorResponse, err error) {
	response = CreateDescribeDBClusterAuditLogCollectorResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterAuditLogCollectorWithChan invokes the polardb.DescribeDBClusterAuditLogCollector API asynchronously
func (client *Client) DescribeDBClusterAuditLogCollectorWithChan(request *DescribeDBClusterAuditLogCollectorRequest) (<-chan *DescribeDBClusterAuditLogCollectorResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterAuditLogCollectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterAuditLogCollector(request)
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

// DescribeDBClusterAuditLogCollectorWithCallback invokes the polardb.DescribeDBClusterAuditLogCollector API asynchronously
func (client *Client) DescribeDBClusterAuditLogCollectorWithCallback(request *DescribeDBClusterAuditLogCollectorRequest, callback func(response *DescribeDBClusterAuditLogCollectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterAuditLogCollectorResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterAuditLogCollector(request)
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

// DescribeDBClusterAuditLogCollectorRequest is the request struct for api DescribeDBClusterAuditLogCollector
type DescribeDBClusterAuditLogCollectorRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterAuditLogCollectorResponse is the response struct for api DescribeDBClusterAuditLogCollector
type DescribeDBClusterAuditLogCollectorResponse struct {
	*responses.BaseResponse
	CollectorStatus string `json:"CollectorStatus" xml:"CollectorStatus"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeDBClusterAuditLogCollectorRequest creates a request to invoke DescribeDBClusterAuditLogCollector API
func CreateDescribeDBClusterAuditLogCollectorRequest() (request *DescribeDBClusterAuditLogCollectorRequest) {
	request = &DescribeDBClusterAuditLogCollectorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterAuditLogCollector", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterAuditLogCollectorResponse creates a response to parse from DescribeDBClusterAuditLogCollector response
func CreateDescribeDBClusterAuditLogCollectorResponse() (response *DescribeDBClusterAuditLogCollectorResponse) {
	response = &DescribeDBClusterAuditLogCollectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
