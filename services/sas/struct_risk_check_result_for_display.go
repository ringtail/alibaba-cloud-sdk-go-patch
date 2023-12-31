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

// RiskCheckResultForDisplay is a nested struct in sas response
type RiskCheckResultForDisplay struct {
	RiskLevel         string             `json:"RiskLevel" xml:"RiskLevel"`
	Status            string             `json:"Status" xml:"Status"`
	Type              string             `json:"Type" xml:"Type"`
	Sort              int                `json:"Sort" xml:"Sort"`
	RepairStatus      string             `json:"RepairStatus" xml:"RepairStatus"`
	RemainingTime     int                `json:"RemainingTime" xml:"RemainingTime"`
	ItemId            int64              `json:"ItemId" xml:"ItemId"`
	StartStatus       string             `json:"StartStatus" xml:"StartStatus"`
	AffectedCount     int                `json:"AffectedCount" xml:"AffectedCount"`
	RiskAssertType    string             `json:"RiskAssertType" xml:"RiskAssertType"`
	Title             string             `json:"Title" xml:"Title"`
	TaskId            int64              `json:"TaskId" xml:"TaskId"`
	CheckTime         int64              `json:"CheckTime" xml:"CheckTime"`
	RiskItemResources []RiskItemResource `json:"RiskItemResources" xml:"RiskItemResources"`
}
