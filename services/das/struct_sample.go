package das

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

// Sample is a nested struct in das response
type Sample struct {
	SqlId      string   `json:"sqlId" xml:"sqlId"`
	Database   string   `json:"database" xml:"database"`
	OriginHost string   `json:"originHost" xml:"originHost"`
	InstanceId string   `json:"instanceId" xml:"instanceId"`
	ErrorCode  string   `json:"errorCode" xml:"errorCode"`
	User       string   `json:"user" xml:"user"`
	Sql        string   `json:"sql" xml:"sql"`
	Timestamp  int64    `json:"timestamp" xml:"timestamp"`
	Tables     []string `json:"tables" xml:"tables"`
}
