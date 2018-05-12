package domain_intl

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

// QueryDomainByInstanceId invokes the domain_intl.QueryDomainByInstanceId API synchronously
// api document: https://help.aliyun.com/api/domain-intl/querydomainbyinstanceid.html
func (client *Client) QueryDomainByInstanceId(request *QueryDomainByInstanceIdRequest) (response *QueryDomainByInstanceIdResponse, err error) {
	response = CreateQueryDomainByInstanceIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainByInstanceIdWithChan invokes the domain_intl.QueryDomainByInstanceId API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/querydomainbyinstanceid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainByInstanceIdWithChan(request *QueryDomainByInstanceIdRequest) (<-chan *QueryDomainByInstanceIdResponse, <-chan error) {
	responseChan := make(chan *QueryDomainByInstanceIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainByInstanceId(request)
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

// QueryDomainByInstanceIdWithCallback invokes the domain_intl.QueryDomainByInstanceId API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/querydomainbyinstanceid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainByInstanceIdWithCallback(request *QueryDomainByInstanceIdRequest, callback func(response *QueryDomainByInstanceIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainByInstanceIdResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainByInstanceId(request)
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

// QueryDomainByInstanceIdRequest is the request struct for api QueryDomainByInstanceId
type QueryDomainByInstanceIdRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

// QueryDomainByInstanceIdResponse is the response struct for api QueryDomainByInstanceId
type QueryDomainByInstanceIdResponse struct {
	*responses.BaseResponse
	UserId                      string  `json:"UserId" xml:"UserId"`
	DomainName                  string  `json:"DomainName" xml:"DomainName"`
	InstanceId                  string  `json:"InstanceId" xml:"InstanceId"`
	RegistrationDate            string  `json:"RegistrationDate" xml:"RegistrationDate"`
	ExpirationDate              string  `json:"ExpirationDate" xml:"ExpirationDate"`
	RegistrantOrganization      string  `json:"RegistrantOrganization" xml:"RegistrantOrganization"`
	RegistrantName              string  `json:"RegistrantName" xml:"RegistrantName"`
	Email                       string  `json:"Email" xml:"Email"`
	UpdateProhibitionLock       string  `json:"UpdateProhibitionLock" xml:"UpdateProhibitionLock"`
	TransferProhibitionLock     string  `json:"TransferProhibitionLock" xml:"TransferProhibitionLock"`
	DomainNameProxyService      bool    `json:"DomainNameProxyService" xml:"DomainNameProxyService"`
	Premium                     bool    `json:"Premium" xml:"Premium"`
	EmailVerificationStatus     int     `json:"EmailVerificationStatus" xml:"EmailVerificationStatus"`
	EmailVerificationClientHold bool    `json:"EmailVerificationClientHold" xml:"EmailVerificationClientHold"`
	DnsList                     DnsList `json:"DnsList" xml:"DnsList"`
}

// CreateQueryDomainByInstanceIdRequest creates a request to invoke QueryDomainByInstanceId API
func CreateQueryDomainByInstanceIdRequest() (request *QueryDomainByInstanceIdRequest) {
	request = &QueryDomainByInstanceIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryDomainByInstanceId", "domain", "openAPI")
	return
}

// CreateQueryDomainByInstanceIdResponse creates a response to parse from QueryDomainByInstanceId response
func CreateQueryDomainByInstanceIdResponse() (response *QueryDomainByInstanceIdResponse) {
	response = &QueryDomainByInstanceIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}