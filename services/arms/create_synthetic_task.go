package arms

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

// CreateSyntheticTask invokes the arms.CreateSyntheticTask API synchronously
func (client *Client) CreateSyntheticTask(request *CreateSyntheticTaskRequest) (response *CreateSyntheticTaskResponse, err error) {
	response = CreateCreateSyntheticTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSyntheticTaskWithChan invokes the arms.CreateSyntheticTask API asynchronously
func (client *Client) CreateSyntheticTaskWithChan(request *CreateSyntheticTaskRequest) (<-chan *CreateSyntheticTaskResponse, <-chan error) {
	responseChan := make(chan *CreateSyntheticTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSyntheticTask(request)
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

// CreateSyntheticTaskWithCallback invokes the arms.CreateSyntheticTask API asynchronously
func (client *Client) CreateSyntheticTaskWithCallback(request *CreateSyntheticTaskRequest, callback func(response *CreateSyntheticTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSyntheticTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateSyntheticTask(request)
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

// CreateSyntheticTaskRequest is the request struct for api CreateSyntheticTask
type CreateSyntheticTaskRequest struct {
	*requests.RpcRequest
	TaskType       requests.Integer                  `position:"Query" name:"TaskType"`
	IntervalType   string                            `position:"Query" name:"IntervalType"`
	TaskName       string                            `position:"Query" name:"TaskName"`
	MonitorList    *[]CreateSyntheticTaskMonitorList `position:"Query" name:"MonitorList"  type:"Json"`
	IpType         requests.Integer                  `position:"Query" name:"IpType"`
	Url            string                            `position:"Query" name:"Url"`
	IntervalTime   string                            `position:"Query" name:"IntervalTime"`
	CommonParam    CreateSyntheticTaskCommonParam    `position:"Query" name:"CommonParam"  type:"Struct"`
	ExtendInterval CreateSyntheticTaskExtendInterval `position:"Query" name:"ExtendInterval"  type:"Struct"`
	Net            CreateSyntheticTaskNet            `position:"Query" name:"Net"  type:"Struct"`
}

// CreateSyntheticTaskMonitorList is a repeated param struct in CreateSyntheticTaskRequest
type CreateSyntheticTaskMonitorList struct {
	NetServiceId string `name:"NetServiceId"`
	MonitorType  string `name:"MonitorType"`
	CityCode     string `name:"CityCode"`
}

// CreateSyntheticTaskCommonParam is a repeated param struct in CreateSyntheticTaskRequest
type CreateSyntheticTaskCommonParam struct {
	AlertNotifierId    string                                         `name:"AlertNotifierId"`
	AlertPolicyId      string                                         `name:"AlertPolicyId"`
	AlertList          *[]CreateSyntheticTaskCommonParamAlertListItem `name:"AlertList" type:"Repeated"`
	AlarmFlag          string                                         `name:"AlarmFlag"`
	StartExecutionTime string                                         `name:"StartExecutionTime"`
	MonitorSamples     string                                         `name:"MonitorSamples"`
}

// CreateSyntheticTaskExtendInterval is a repeated param struct in CreateSyntheticTaskRequest
type CreateSyntheticTaskExtendInterval struct {
	StartMinute string    `name:"StartMinute"`
	EndHour     string    `name:"EndHour"`
	EndMinute   string    `name:"EndMinute"`
	StartHour   string    `name:"StartHour"`
	EndTime     string    `name:"EndTime"`
	Days        *[]string `name:"Days" type:"Repeated"`
	StartTime   string    `name:"StartTime"`
}

// CreateSyntheticTaskNet is a repeated param struct in CreateSyntheticTaskRequest
type CreateSyntheticTaskNet struct {
	NetTraceRouteTimeout string `name:"NetTraceRouteTimeout"`
	NetICMPActive        string `name:"NetICMPActive"`
	NetICMPNum           string `name:"NetICMPNum"`
	NetDNSTimeout        string `name:"NetDNSTimeout"`
	NetDNSQueryMethod    string `name:"NetDNSQueryMethod"`
	WhiteList            string `name:"WhiteList"`
	NetDNSNs             string `name:"NetDNSNs"`
	NetDNSServer         string `name:"NetDNSServer"`
	NetTraceRouteSwitch  string `name:"NetTraceRouteSwitch"`
	NetDigSwitch         string `name:"NetDigSwitch"`
	NetICMPInterval      string `name:"NetICMPInterval"`
	NetDNSSwitch         string `name:"NetDNSSwitch"`
	NetTraceRouteNum     string `name:"NetTraceRouteNum"`
	NetICMPSwitch        string `name:"NetICMPSwitch"`
	NetICMPSize          string `name:"NetICMPSize"`
}

// CreateSyntheticTaskCommonParamAlertListItem is a repeated param struct in CreateSyntheticTaskRequest
type CreateSyntheticTaskCommonParamAlertListItem struct {
	IsCritical   string `name:"IsCritical"`
	Name         string `name:"Name"`
	SeriousAlert string `name:"SeriousAlert"`
	GeneralAlert string `name:"GeneralAlert"`
	Symbols      string `name:"Symbols"`
}

// CreateSyntheticTaskResponse is the response struct for api CreateSyntheticTask
type CreateSyntheticTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateSyntheticTaskRequest creates a request to invoke CreateSyntheticTask API
func CreateCreateSyntheticTaskRequest() (request *CreateSyntheticTaskRequest) {
	request = &CreateSyntheticTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreateSyntheticTask", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSyntheticTaskResponse creates a response to parse from CreateSyntheticTask response
func CreateCreateSyntheticTaskResponse() (response *CreateSyntheticTaskResponse) {
	response = &CreateSyntheticTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
