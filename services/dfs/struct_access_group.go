package dfs

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

// AccessGroup is a nested struct in dfs response
type AccessGroup struct {
	AccessGroupId   string `json:"AccessGroupId" xml:"AccessGroupId"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
	IsDefault       bool   `json:"IsDefault" xml:"IsDefault"`
	MountPointCount int    `json:"MountPointCount" xml:"MountPointCount"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	RuleCount       int    `json:"RuleCount" xml:"RuleCount"`
	AccessGroupName string `json:"AccessGroupName" xml:"AccessGroupName"`
	NetworkType     string `json:"NetworkType" xml:"NetworkType"`
	Description     string `json:"Description" xml:"Description"`
}
