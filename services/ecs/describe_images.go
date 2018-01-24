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

func (client *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	response = CreateDescribeImagesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeImagesWithChan(request *DescribeImagesRequest) (<-chan *DescribeImagesResponse, <-chan error) {
	responseChan := make(chan *DescribeImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImages(request)
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

func (client *Client) DescribeImagesWithCallback(request *DescribeImagesRequest, callback func(response *DescribeImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeImages(request)
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

type DescribeImagesRequest struct {
	*requests.RpcRequest
	Tag4Value            string           `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ImageId              string           `position:"Query" name:"ImageId"`
	SnapshotId           string           `position:"Query" name:"SnapshotId"`
	Tag2Key              string           `position:"Query" name:"Tag.2.Key"`
	Filter2Value         string           `position:"Query" name:"Filter.2.Value"`
	Usage                string           `position:"Query" name:"Usage"`
	Tag3Key              string           `position:"Query" name:"Tag.3.Key"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ImageOwnerAlias      string           `position:"Query" name:"ImageOwnerAlias"`
	Tag1Value            string           `position:"Query" name:"Tag.1.Value"`
	IsSupportIoOptimized requests.Boolean `position:"Query" name:"IsSupportIoOptimized"`
	Filter1Key           string           `position:"Query" name:"Filter.1.Key"`
	ImageName            string           `position:"Query" name:"ImageName"`
	IsSupportCloudinit   requests.Boolean `position:"Query" name:"IsSupportCloudinit"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	Tag3Value            string           `position:"Query" name:"Tag.3.Value"`
	Architecture         string           `position:"Query" name:"Architecture"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	Tag5Key              string           `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ShowExpired          requests.Boolean `position:"Query" name:"ShowExpired"`
	Filter1Value         string           `position:"Query" name:"Filter.1.Value"`
	OSType               string           `position:"Query" name:"OSType"`
	Filter2Key           string           `position:"Query" name:"Filter.2.Key"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Value            string           `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string           `position:"Query" name:"Tag.1.Key"`
	Tag2Value            string           `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.Key"`
	Status               string           `position:"Query" name:"Status"`
}

type DescribeImagesResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	RegionId   string `json:"RegionId" xml:"RegionId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Images     struct {
		Image []struct {
			Progress             string `json:"Progress" xml:"Progress"`
			ImageId              string `json:"ImageId" xml:"ImageId"`
			ImageName            string `json:"ImageName" xml:"ImageName"`
			ImageVersion         string `json:"ImageVersion" xml:"ImageVersion"`
			Description          string `json:"Description" xml:"Description"`
			Size                 int    `json:"Size" xml:"Size"`
			ImageOwnerAlias      string `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
			IsSupportIoOptimized bool   `json:"IsSupportIoOptimized" xml:"IsSupportIoOptimized"`
			IsSupportCloudinit   bool   `json:"IsSupportCloudinit" xml:"IsSupportCloudinit"`
			OSName               string `json:"OSName" xml:"OSName"`
			Architecture         string `json:"Architecture" xml:"Architecture"`
			Status               string `json:"Status" xml:"Status"`
			ProductCode          string `json:"ProductCode" xml:"ProductCode"`
			IsSubscribed         bool   `json:"IsSubscribed" xml:"IsSubscribed"`
			CreationTime         string `json:"CreationTime" xml:"CreationTime"`
			IsSelfShared         string `json:"IsSelfShared" xml:"IsSelfShared"`
			OSType               string `json:"OSType" xml:"OSType"`
			Platform             string `json:"Platform" xml:"Platform"`
			Usage                string `json:"Usage" xml:"Usage"`
			IsCopied             bool   `json:"IsCopied" xml:"IsCopied"`
			DiskDeviceMappings   struct {
				DiskDeviceMapping []struct {
					SnapshotId      string `json:"SnapshotId" xml:"SnapshotId"`
					Size            string `json:"Size" xml:"Size"`
					Device          string `json:"Device" xml:"Device"`
					Type            string `json:"Type" xml:"Type"`
					Format          string `json:"Format" xml:"Format"`
					ImportOSSBucket string `json:"ImportOSSBucket" xml:"ImportOSSBucket"`
					ImportOSSObject string `json:"ImportOSSObject" xml:"ImportOSSObject"`
				} `json:"DiskDeviceMapping" xml:"DiskDeviceMapping"`
			} `json:"DiskDeviceMappings" xml:"DiskDeviceMappings"`
			Tags struct {
				Tag []struct {
					TagKey   string `json:"TagKey" xml:"TagKey"`
					TagValue string `json:"TagValue" xml:"TagValue"`
				} `json:"Tag" xml:"Tag"`
			} `json:"Tags" xml:"Tags"`
		} `json:"Image" xml:"Image"`
	} `json:"Images" xml:"Images"`
}

func CreateDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImages", "ecs", "openAPI")
	return
}

func CreateDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}