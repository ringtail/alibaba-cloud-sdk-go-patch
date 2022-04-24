package polardb

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

// DBAccount is a nested struct in polardb response
type DBAccount struct {
	AccountStatus            string              `json:"AccountStatus" xml:"AccountStatus"`
	AccountDescription       string              `json:"AccountDescription" xml:"AccountDescription"`
	PrivilegeExceeded        string              `json:"PrivilegeExceeded" xml:"PrivilegeExceeded"`
	AccountPasswordValidTime string              `json:"AccountPasswordValidTime" xml:"AccountPasswordValidTime"`
	AccountType              string              `json:"AccountType" xml:"AccountType"`
	AccountLockState         string              `json:"AccountLockState" xml:"AccountLockState"`
	AccountName              string              `json:"AccountName" xml:"AccountName"`
	DatabasePrivileges       []DatabasePrivilege `json:"DatabasePrivileges" xml:"DatabasePrivileges"`
}
