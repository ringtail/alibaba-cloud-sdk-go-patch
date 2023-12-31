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

// DispatchRule is a nested struct in arms response
type DispatchRule struct {
	Name                     string                             `json:"Name" xml:"Name"`
	DispatchType             string                             `json:"DispatchType" xml:"DispatchType"`
	State                    string                             `json:"State" xml:"State"`
	RuleId                   int64                              `json:"RuleId" xml:"RuleId"`
	IsRecover                bool                               `json:"IsRecover" xml:"IsRecover"`
	Id                       int64                              `json:"id" xml:"id"`
	LabelMatchExpressionGrid LabelMatchExpressionGrid           `json:"LabelMatchExpressionGrid" xml:"LabelMatchExpressionGrid"`
	GroupRules               []GroupRule                        `json:"GroupRules" xml:"GroupRules"`
	NotifyRules              []NotifyRuleInDescribeDispatchRule `json:"NotifyRules" xml:"NotifyRules"`
}
