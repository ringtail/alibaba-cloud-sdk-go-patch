package linkvisual

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

// QueryFaceUserByName invokes the linkvisual.QueryFaceUserByName API synchronously
func (client *Client) QueryFaceUserByName(request *QueryFaceUserByNameRequest) (response *QueryFaceUserByNameResponse, err error) {
	response = CreateQueryFaceUserByNameResponse()
	err = client.DoAction(request, response)
	return
}

// QueryFaceUserByNameWithChan invokes the linkvisual.QueryFaceUserByName API asynchronously
func (client *Client) QueryFaceUserByNameWithChan(request *QueryFaceUserByNameRequest) (<-chan *QueryFaceUserByNameResponse, <-chan error) {
	responseChan := make(chan *QueryFaceUserByNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryFaceUserByName(request)
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

// QueryFaceUserByNameWithCallback invokes the linkvisual.QueryFaceUserByName API asynchronously
func (client *Client) QueryFaceUserByNameWithCallback(request *QueryFaceUserByNameRequest, callback func(response *QueryFaceUserByNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryFaceUserByNameResponse
		var err error
		defer close(result)
		response, err = client.QueryFaceUserByName(request)
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

// QueryFaceUserByNameRequest is the request struct for api QueryFaceUserByName
type QueryFaceUserByNameRequest struct {
	*requests.RpcRequest
	IsolationId string           `position:"Query" name:"IsolationId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Params      string           `position:"Query" name:"Params"`
	PageNo      requests.Integer `position:"Query" name:"PageNo"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	Name        string           `position:"Query" name:"Name"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
}

// QueryFaceUserByNameResponse is the response struct for api QueryFaceUserByName
type QueryFaceUserByNameResponse struct {
	*responses.BaseResponse
	Code         string                    `json:"Code" xml:"Code"`
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool                      `json:"Success" xml:"Success"`
	Data         DataInQueryFaceUserByName `json:"Data" xml:"Data"`
}

// CreateQueryFaceUserByNameRequest creates a request to invoke QueryFaceUserByName API
func CreateQueryFaceUserByNameRequest() (request *QueryFaceUserByNameRequest) {
	request = &QueryFaceUserByNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryFaceUserByName", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryFaceUserByNameResponse creates a response to parse from QueryFaceUserByName response
func CreateQueryFaceUserByNameResponse() (response *QueryFaceUserByNameResponse) {
	response = &QueryFaceUserByNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
