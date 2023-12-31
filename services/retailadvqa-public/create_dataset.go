package retailadvqa_public

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

// CreateDataset invokes the retailadvqa_public.CreateDataset API synchronously
func (client *Client) CreateDataset(request *CreateDatasetRequest) (response *CreateDatasetResponse, err error) {
	response = CreateCreateDatasetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDatasetWithChan invokes the retailadvqa_public.CreateDataset API asynchronously
func (client *Client) CreateDatasetWithChan(request *CreateDatasetRequest) (<-chan *CreateDatasetResponse, <-chan error) {
	responseChan := make(chan *CreateDatasetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataset(request)
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

// CreateDatasetWithCallback invokes the retailadvqa_public.CreateDataset API asynchronously
func (client *Client) CreateDatasetWithCallback(request *CreateDatasetRequest, callback func(response *CreateDatasetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDatasetResponse
		var err error
		defer close(result)
		response, err = client.CreateDataset(request)
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

// CreateDatasetRequest is the request struct for api CreateDataset
type CreateDatasetRequest struct {
	*requests.RpcRequest
	AccessId    string `position:"Body" name:"AccessId"`
	TenantId    string `position:"Body" name:"TenantId"`
	DataSetName string `position:"Body" name:"DataSetName"`
	Type        string `position:"Body" name:"Type"`
	DataSet     string `position:"Body" name:"DataSet"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// CreateDatasetResponse is the response struct for api CreateDataset
type CreateDatasetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateDatasetRequest creates a request to invoke CreateDataset API
func CreateCreateDatasetRequest() (request *CreateDatasetRequest) {
	request = &CreateDatasetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CreateDataset", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDatasetResponse creates a response to parse from CreateDataset response
func CreateCreateDatasetResponse() (response *CreateDatasetResponse) {
	response = &CreateDatasetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
