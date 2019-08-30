package rds

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

// DescribeDBInstanceIpHostname invokes the rds.DescribeDBInstanceIpHostname API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceiphostname.html
func (client *Client) DescribeDBInstanceIpHostname(request *DescribeDBInstanceIpHostnameRequest) (response *DescribeDBInstanceIpHostnameResponse, err error) {
	response = CreateDescribeDBInstanceIpHostnameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceIpHostnameWithChan invokes the rds.DescribeDBInstanceIpHostname API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceiphostname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceIpHostnameWithChan(request *DescribeDBInstanceIpHostnameRequest) (<-chan *DescribeDBInstanceIpHostnameResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceIpHostnameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceIpHostname(request)
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

// DescribeDBInstanceIpHostnameWithCallback invokes the rds.DescribeDBInstanceIpHostname API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceiphostname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceIpHostnameWithCallback(request *DescribeDBInstanceIpHostnameRequest, callback func(response *DescribeDBInstanceIpHostnameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceIpHostnameResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceIpHostname(request)
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

// DescribeDBInstanceIpHostnameRequest is the request struct for api DescribeDBInstanceIpHostname
type DescribeDBInstanceIpHostnameRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceIpHostnameResponse is the response struct for api DescribeDBInstanceIpHostname
type DescribeDBInstanceIpHostnameResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	DBInstanceId    string `json:"DBInstanceId" xml:"DBInstanceId"`
	IpHostnameInfos string `json:"IpHostnameInfos" xml:"IpHostnameInfos"`
}

// CreateDescribeDBInstanceIpHostnameRequest creates a request to invoke DescribeDBInstanceIpHostname API
func CreateDescribeDBInstanceIpHostnameRequest() (request *DescribeDBInstanceIpHostnameRequest) {
	request = &DescribeDBInstanceIpHostnameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceIpHostname", "rds", "openAPI")
	return
}

// CreateDescribeDBInstanceIpHostnameResponse creates a response to parse from DescribeDBInstanceIpHostname response
func CreateDescribeDBInstanceIpHostnameResponse() (response *DescribeDBInstanceIpHostnameResponse) {
	response = &DescribeDBInstanceIpHostnameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}