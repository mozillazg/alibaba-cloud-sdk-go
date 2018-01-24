package ecs

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) RevokeSecurityGroupEgress(request *RevokeSecurityGroupEgressRequest) (response *RevokeSecurityGroupEgressResponse, err error) {
	response = CreateRevokeSecurityGroupEgressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RevokeSecurityGroupEgressWithChan(request *RevokeSecurityGroupEgressRequest) (<-chan *RevokeSecurityGroupEgressResponse, <-chan error) {
	responseChan := make(chan *RevokeSecurityGroupEgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeSecurityGroupEgress(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}

	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) RevokeSecurityGroupEgressWithCallback(request *RevokeSecurityGroupEgressRequest, callback func(response *RevokeSecurityGroupEgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeSecurityGroupEgressResponse
		var err error
		defer close(result)
		response, err = client.RevokeSecurityGroupEgress(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type RevokeSecurityGroupEgressRequest struct {
	*requests.RpcRequest
	NicType               string           `position:"Query" name:"NicType"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange       string           `position:"Query" name:"SourcePortRange"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	SecurityGroupId       string           `position:"Query" name:"SecurityGroupId"`
	Description           string           `position:"Query" name:"Description"`
	Policy                string           `position:"Query" name:"Policy"`
	PortRange             string           `position:"Query" name:"PortRange"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol            string           `position:"Query" name:"IpProtocol"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	SourceCidrIp          string           `position:"Query" name:"SourceCidrIp"`
	DestGroupId           string           `position:"Query" name:"DestGroupId"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	DestGroupOwnerAccount string           `position:"Query" name:"DestGroupOwnerAccount"`
	Priority              string           `position:"Query" name:"Priority"`
	DestCidrIp            string           `position:"Query" name:"DestCidrIp"`
	DestGroupOwnerId      requests.Integer `position:"Query" name:"DestGroupOwnerId"`
}

type RevokeSecurityGroupEgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateRevokeSecurityGroupEgressRequest() (request *RevokeSecurityGroupEgressRequest) {
	request = &RevokeSecurityGroupEgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RevokeSecurityGroupEgress", "ecs", "openAPI")
	return
}

func CreateRevokeSecurityGroupEgressResponse() (response *RevokeSecurityGroupEgressResponse) {
	response = &RevokeSecurityGroupEgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}