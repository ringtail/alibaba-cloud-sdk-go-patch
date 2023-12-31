package sas

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

// Data is a nested struct in sas response
type Data struct {
	ControlNodeName          string        `json:"ControlNodeName" xml:"ControlNodeName"`
	DataSource               string        `json:"DataSource" xml:"DataSource"`
	SeriousAlarmCount        int           `json:"SeriousAlarmCount" xml:"SeriousAlarmCount"`
	EndTime                  int64         `json:"EndTime" xml:"EndTime"`
	RemindAlarmCount         int           `json:"RemindAlarmCount" xml:"RemindAlarmCount"`
	ContainerId              string        `json:"ContainerId" xml:"ContainerId"`
	StartTime                int64         `json:"StartTime" xml:"StartTime"`
	TaskId                   string        `json:"TaskId" xml:"TaskId"`
	FinishCount              int           `json:"FinishCount" xml:"FinishCount"`
	TotalCount               int           `json:"TotalCount" xml:"TotalCount"`
	Progress                 int           `json:"Progress" xml:"Progress"`
	IntranetIp               string        `json:"IntranetIp" xml:"IntranetIp"`
	Type                     string        `json:"Type" xml:"Type"`
	Solution                 string        `json:"Solution" xml:"Solution"`
	K8sClusterName           string        `json:"K8sClusterName" xml:"K8sClusterName"`
	TotalAlarmCount          int           `json:"TotalAlarmCount" xml:"TotalAlarmCount"`
	K8sClusterId             string        `json:"K8sClusterId" xml:"K8sClusterId"`
	InstanceName             string        `json:"InstanceName" xml:"InstanceName"`
	HoneypotImageName        string        `json:"HoneypotImageName" xml:"HoneypotImageName"`
	ContainerImageName       string        `json:"ContainerImageName" xml:"ContainerImageName"`
	AlarmEventDesc           string        `json:"AlarmEventDesc" xml:"AlarmEventDesc"`
	CollectTime              int64         `json:"CollectTime" xml:"CollectTime"`
	ScanImageCount           int           `json:"ScanImageCount" xml:"ScanImageCount"`
	AlarmEventAliasName      string        `json:"AlarmEventAliasName" xml:"AlarmEventAliasName"`
	Status                   string        `json:"Status" xml:"Status"`
	NeedAuthCount            int           `json:"NeedAuthCount" xml:"NeedAuthCount"`
	CanBeDealOnLine          bool          `json:"CanBeDealOnLine" xml:"CanBeDealOnLine"`
	HoneypotImageDisplayName string        `json:"HoneypotImageDisplayName" xml:"HoneypotImageDisplayName"`
	TotalNode                int           `json:"TotalNode" xml:"TotalNode"`
	K8sNodeName              string        `json:"K8sNodeName" xml:"K8sNodeName"`
	AlarmUniqueInfo          string        `json:"AlarmUniqueInfo" xml:"AlarmUniqueInfo"`
	ContainHwMode            bool          `json:"ContainHwMode" xml:"ContainHwMode"`
	Uuid                     string        `json:"Uuid" xml:"Uuid"`
	NodeId                   string        `json:"NodeId" xml:"NodeId"`
	HoneypotId               string        `json:"HoneypotId" xml:"HoneypotId"`
	CanCreate                bool          `json:"CanCreate" xml:"CanCreate"`
	SuspiciousAlarmCount     int           `json:"SuspiciousAlarmCount" xml:"SuspiciousAlarmCount"`
	HasRiskNode              int           `json:"hasRiskNode" xml:"hasRiskNode"`
	CanCancelFault           bool          `json:"CanCancelFault" xml:"CanCancelFault"`
	InternetIp               string        `json:"InternetIp" xml:"InternetIp"`
	Result                   string        `json:"Result" xml:"Result"`
	Level                    string        `json:"Level" xml:"Level"`
	ContainerImageId         string        `json:"ContainerImageId" xml:"ContainerImageId"`
	K8sPodName               string        `json:"K8sPodName" xml:"K8sPodName"`
	K8sNamespace             string        `json:"K8sNamespace" xml:"K8sNamespace"`
	K8sNodeId                string        `json:"K8sNodeId" xml:"K8sNodeId"`
	HoneypotName             string        `json:"HoneypotName" xml:"HoneypotName"`
	PresetId                 string        `json:"PresetId" xml:"PresetId"`
	ExecTime                 int64         `json:"ExecTime" xml:"ExecTime"`
	AppName                  string        `json:"AppName" xml:"AppName"`
	State                    []string      `json:"State" xml:"State"`
	CauseDetails             []CauseDetail `json:"CauseDetails" xml:"CauseDetails"`
}
