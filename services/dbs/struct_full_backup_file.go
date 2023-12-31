package dbs

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

// FullBackupFile is a nested struct in dbs response
type FullBackupFile struct {
	FinishTime           int64  `json:"FinishTime" xml:"FinishTime"`
	BackupStatus         string `json:"BackupStatus" xml:"BackupStatus"`
	SourceEndpointIpPort string `json:"SourceEndpointIpPort" xml:"SourceEndpointIpPort"`
	CreateTime           int64  `json:"CreateTime" xml:"CreateTime"`
	ErrMessage           string `json:"ErrMessage" xml:"ErrMessage"`
	BackupObjects        string `json:"BackupObjects" xml:"BackupObjects"`
	EndTime              int64  `json:"EndTime" xml:"EndTime"`
	StartTime            int64  `json:"StartTime" xml:"StartTime"`
	BackupSetExpiredTime int64  `json:"BackupSetExpiredTime" xml:"BackupSetExpiredTime"`
	StorageMethod        string `json:"StorageMethod" xml:"StorageMethod"`
	BackupSetId          string `json:"BackupSetId" xml:"BackupSetId"`
	BackupSize           int64  `json:"BackupSize" xml:"BackupSize"`
}
