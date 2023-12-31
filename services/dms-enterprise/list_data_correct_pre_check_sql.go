package dms_enterprise

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

// ListDataCorrectPreCheckSQL invokes the dms_enterprise.ListDataCorrectPreCheckSQL API synchronously
func (client *Client) ListDataCorrectPreCheckSQL(request *ListDataCorrectPreCheckSQLRequest) (response *ListDataCorrectPreCheckSQLResponse, err error) {
	response = CreateListDataCorrectPreCheckSQLResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataCorrectPreCheckSQLWithChan invokes the dms_enterprise.ListDataCorrectPreCheckSQL API asynchronously
func (client *Client) ListDataCorrectPreCheckSQLWithChan(request *ListDataCorrectPreCheckSQLRequest) (<-chan *ListDataCorrectPreCheckSQLResponse, <-chan error) {
	responseChan := make(chan *ListDataCorrectPreCheckSQLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataCorrectPreCheckSQL(request)
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

// ListDataCorrectPreCheckSQLWithCallback invokes the dms_enterprise.ListDataCorrectPreCheckSQL API asynchronously
func (client *Client) ListDataCorrectPreCheckSQLWithCallback(request *ListDataCorrectPreCheckSQLRequest, callback func(response *ListDataCorrectPreCheckSQLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataCorrectPreCheckSQLResponse
		var err error
		defer close(result)
		response, err = client.ListDataCorrectPreCheckSQL(request)
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

// ListDataCorrectPreCheckSQLRequest is the request struct for api ListDataCorrectPreCheckSQL
type ListDataCorrectPreCheckSQLRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	OrderId    requests.Integer `position:"Query" name:"OrderId"`
	DbId       requests.Integer `position:"Query" name:"DbId"`
}

// ListDataCorrectPreCheckSQLResponse is the response struct for api ListDataCorrectPreCheckSQL
type ListDataCorrectPreCheckSQLResponse struct {
	*responses.BaseResponse
	RequestId       string        `json:"RequestId" xml:"RequestId"`
	Success         bool          `json:"Success" xml:"Success"`
	ErrorMessage    string        `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode       string        `json:"ErrorCode" xml:"ErrorCode"`
	PreCheckSQLList []PreCheckSQL `json:"PreCheckSQLList" xml:"PreCheckSQLList"`
}

// CreateListDataCorrectPreCheckSQLRequest creates a request to invoke ListDataCorrectPreCheckSQL API
func CreateListDataCorrectPreCheckSQLRequest() (request *ListDataCorrectPreCheckSQLRequest) {
	request = &ListDataCorrectPreCheckSQLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDataCorrectPreCheckSQL", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataCorrectPreCheckSQLResponse creates a response to parse from ListDataCorrectPreCheckSQL response
func CreateListDataCorrectPreCheckSQLResponse() (response *ListDataCorrectPreCheckSQLResponse) {
	response = &ListDataCorrectPreCheckSQLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
