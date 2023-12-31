package dataworks_public

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

// RuleChecksItemInListQualityResultsByRule is a nested struct in dataworks_public response
type RuleChecksItemInListQualityResultsByRule struct {
	BlockType         int                  `json:"BlockType" xml:"BlockType"`
	WarningThreshold  float64              `json:"WarningThreshold" xml:"WarningThreshold"`
	Property          string               `json:"Property" xml:"Property"`
	TableName         string               `json:"TableName" xml:"TableName"`
	Comment           string               `json:"Comment" xml:"Comment"`
	CheckResultStatus int                  `json:"CheckResultStatus" xml:"CheckResultStatus"`
	TemplateName      string               `json:"TemplateName" xml:"TemplateName"`
	CheckerName       string               `json:"CheckerName" xml:"CheckerName"`
	RuleId            int64                `json:"RuleId" xml:"RuleId"`
	FixedCheck        bool                 `json:"FixedCheck" xml:"FixedCheck"`
	Op                string               `json:"Op" xml:"Op"`
	UpperValue        float64              `json:"UpperValue" xml:"UpperValue"`
	ActualExpression  string               `json:"ActualExpression" xml:"ActualExpression"`
	ExternalId        string               `json:"ExternalId" xml:"ExternalId"`
	TimeCost          string               `json:"TimeCost" xml:"TimeCost"`
	Trend             string               `json:"Trend" xml:"Trend"`
	ExternalType      string               `json:"ExternalType" xml:"ExternalType"`
	BizDate           int64                `json:"BizDate" xml:"BizDate"`
	CheckResult       int                  `json:"CheckResult" xml:"CheckResult"`
	ResultString      string               `json:"ResultString" xml:"ResultString"`
	MatchExpression   string               `json:"MatchExpression" xml:"MatchExpression"`
	CheckerType       int                  `json:"CheckerType" xml:"CheckerType"`
	ProjectName       string               `json:"ProjectName" xml:"ProjectName"`
	BeginTime         int64                `json:"BeginTime" xml:"BeginTime"`
	DateType          string               `json:"DateType" xml:"DateType"`
	CriticalThreshold float64              `json:"CriticalThreshold" xml:"CriticalThreshold"`
	IsPrediction      bool                 `json:"IsPrediction" xml:"IsPrediction"`
	RuleName          string               `json:"RuleName" xml:"RuleName"`
	CheckerId         int                  `json:"CheckerId" xml:"CheckerId"`
	DiscreteCheck     bool                 `json:"DiscreteCheck" xml:"DiscreteCheck"`
	EndTime           int64                `json:"EndTime" xml:"EndTime"`
	MethodName        string               `json:"MethodName" xml:"MethodName"`
	LowerValue        float64              `json:"LowerValue" xml:"LowerValue"`
	EntityId          int64                `json:"EntityId" xml:"EntityId"`
	WhereCondition    string               `json:"WhereCondition" xml:"WhereCondition"`
	ExpectValue       float64              `json:"ExpectValue" xml:"ExpectValue"`
	TemplateId        int                  `json:"TemplateId" xml:"TemplateId"`
	TaskId            string               `json:"TaskId" xml:"TaskId"`
	Id                int64                `json:"Id" xml:"Id"`
	Open              bool                 `json:"Open" xml:"Open"`
	ReferenceValue    []ReferenceValueItem `json:"ReferenceValue" xml:"ReferenceValue"`
	SampleValue       []SampleValueItem    `json:"SampleValue" xml:"SampleValue"`
}
