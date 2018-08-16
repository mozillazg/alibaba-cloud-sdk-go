package imm

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

// GroupsItem is a nested struct in imm response
type GroupsItem struct {
	UnGroupReason string `json:"UnGroupReason" xml:"UnGroupReason"`
	ModifyTime    string `json:"ModifyTime" xml:"ModifyTime"`
	CreateTime    string `json:"CreateTime" xml:"CreateTime"`
	GroupName     string `json:"GroupName" xml:"GroupName"`
	Count         string `json:"Count" xml:"Count"`
	FaceId        string `json:"FaceId" xml:"FaceId"`
	GroupId       string `json:"GroupId" xml:"GroupId"`
	Status        string `json:"Status" xml:"Status"`
}