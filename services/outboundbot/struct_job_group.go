package outboundbot

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

// JobGroup is a nested struct in outboundbot response
type JobGroup struct {
	ScriptName           string         `json:"ScriptName" xml:"ScriptName"`
	JobGroupId           string         `json:"JobGroupId" xml:"JobGroupId"`
	ModifyTime           string         `json:"ModifyTime" xml:"ModifyTime"`
	MinConcurrency       int64          `json:"MinConcurrency" xml:"MinConcurrency"`
	JobGroupName         string         `json:"JobGroupName" xml:"JobGroupName"`
	TotalCallNum         int            `json:"TotalCallNum" xml:"TotalCallNum"`
	JobGroupDescription  string         `json:"JobGroupDescription" xml:"JobGroupDescription"`
	Id                   string         `json:"Id" xml:"Id"`
	RingingDuration      int64          `json:"RingingDuration" xml:"RingingDuration"`
	CreationTime         int64          `json:"CreationTime" xml:"CreationTime"`
	Priority             string         `json:"Priority" xml:"Priority"`
	ScenarioId           string         `json:"ScenarioId" xml:"ScenarioId"`
	ScriptVersion        string         `json:"ScriptVersion" xml:"ScriptVersion"`
	ScriptId             string         `json:"ScriptId" xml:"ScriptId"`
	JobDataParsingTaskId string         `json:"JobDataParsingTaskId" xml:"JobDataParsingTaskId"`
	JobFilePath          string         `json:"JobFilePath" xml:"JobFilePath"`
	Status               string         `json:"Status" xml:"Status"`
	CallingNumbers       []string       `json:"CallingNumbers" xml:"CallingNumbers"`
	Progress             Progress       `json:"Progress" xml:"Progress"`
	ExportProgress       ExportProgress `json:"ExportProgress" xml:"ExportProgress"`
	RecallStrategy       RecallStrategy `json:"RecallStrategy" xml:"RecallStrategy"`
	Strategy             Strategy       `json:"Strategy" xml:"Strategy"`
	Result               Result         `json:"Result" xml:"Result"`
}
