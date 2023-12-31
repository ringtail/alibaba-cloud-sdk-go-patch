package viapi_regen

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

// DataInGetService is a nested struct in viapi_regen response
type DataInGetService struct {
	Id                 int64          `json:"Id" xml:"Id"`
	GmtCreate          int64          `json:"GmtCreate" xml:"GmtCreate"`
	ServiceName        string         `json:"ServiceName" xml:"ServiceName"`
	ServiceDescription string         `json:"ServiceDescription" xml:"ServiceDescription"`
	Status             string         `json:"Status" xml:"Status"`
	ServiceId          string         `json:"ServiceId" xml:"ServiceId"`
	InputParams        string         `json:"InputParams" xml:"InputParams"`
	OutputParams       string         `json:"OutputParams" xml:"OutputParams"`
	Errorcodes         string         `json:"Errorcodes" xml:"Errorcodes"`
	InputExample       string         `json:"InputExample" xml:"InputExample"`
	OutputExample      string         `json:"OutputExample" xml:"OutputExample"`
	DataReflowInfo     DataReflowInfo `json:"DataReflowInfo" xml:"DataReflowInfo"`
}
