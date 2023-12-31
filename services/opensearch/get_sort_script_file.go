package opensearch

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

// GetSortScriptFile invokes the opensearch.GetSortScriptFile API synchronously
func (client *Client) GetSortScriptFile(request *GetSortScriptFileRequest) (response *GetSortScriptFileResponse, err error) {
	response = CreateGetSortScriptFileResponse()
	err = client.DoAction(request, response)
	return
}

// GetSortScriptFileWithChan invokes the opensearch.GetSortScriptFile API asynchronously
func (client *Client) GetSortScriptFileWithChan(request *GetSortScriptFileRequest) (<-chan *GetSortScriptFileResponse, <-chan error) {
	responseChan := make(chan *GetSortScriptFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSortScriptFile(request)
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

// GetSortScriptFileWithCallback invokes the opensearch.GetSortScriptFile API asynchronously
func (client *Client) GetSortScriptFileWithCallback(request *GetSortScriptFileRequest, callback func(response *GetSortScriptFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSortScriptFileResponse
		var err error
		defer close(result)
		response, err = client.GetSortScriptFile(request)
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

// GetSortScriptFileRequest is the request struct for api GetSortScriptFile
type GetSortScriptFileRequest struct {
	*requests.RoaRequest
	AppVersionId     string `position:"Path" name:"appVersionId"`
	FileName         string `position:"Path" name:"fileName"`
	ScriptName       string `position:"Path" name:"scriptName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// GetSortScriptFileResponse is the response struct for api GetSortScriptFile
type GetSortScriptFileResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateGetSortScriptFileRequest creates a request to invoke GetSortScriptFile API
func CreateGetSortScriptFileRequest() (request *GetSortScriptFileRequest) {
	request = &GetSortScriptFileRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GetSortScriptFile", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appVersionId]/sort-scripts/[scriptName]/files/src/[fileName]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetSortScriptFileResponse creates a response to parse from GetSortScriptFile response
func CreateGetSortScriptFileResponse() (response *GetSortScriptFileResponse) {
	response = &GetSortScriptFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
