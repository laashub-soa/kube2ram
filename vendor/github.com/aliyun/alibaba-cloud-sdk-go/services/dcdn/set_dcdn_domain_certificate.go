package dcdn

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

// SetDcdnDomainCertificate invokes the dcdn.SetDcdnDomainCertificate API synchronously
// api document: https://help.aliyun.com/api/dcdn/setdcdndomaincertificate.html
func (client *Client) SetDcdnDomainCertificate(request *SetDcdnDomainCertificateRequest) (response *SetDcdnDomainCertificateResponse, err error) {
	response = CreateSetDcdnDomainCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// SetDcdnDomainCertificateWithChan invokes the dcdn.SetDcdnDomainCertificate API asynchronously
// api document: https://help.aliyun.com/api/dcdn/setdcdndomaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDcdnDomainCertificateWithChan(request *SetDcdnDomainCertificateRequest) (<-chan *SetDcdnDomainCertificateResponse, <-chan error) {
	responseChan := make(chan *SetDcdnDomainCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDcdnDomainCertificate(request)
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

// SetDcdnDomainCertificateWithCallback invokes the dcdn.SetDcdnDomainCertificate API asynchronously
// api document: https://help.aliyun.com/api/dcdn/setdcdndomaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDcdnDomainCertificateWithCallback(request *SetDcdnDomainCertificateRequest, callback func(response *SetDcdnDomainCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDcdnDomainCertificateResponse
		var err error
		defer close(result)
		response, err = client.SetDcdnDomainCertificate(request)
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

// SetDcdnDomainCertificateRequest is the request struct for api SetDcdnDomainCertificate
type SetDcdnDomainCertificateRequest struct {
	*requests.RpcRequest
	ForceSet      string           `position:"Query" name:"ForceSet"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	CertType      string           `position:"Query" name:"CertType"`
	SSLPub        string           `position:"Query" name:"SSLPub"`
	CertName      string           `position:"Query" name:"CertName"`
	SSLProtocol   string           `position:"Query" name:"SSLProtocol"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Region        string           `position:"Query" name:"Region"`
	SSLPri        string           `position:"Query" name:"SSLPri"`
}

// SetDcdnDomainCertificateResponse is the response struct for api SetDcdnDomainCertificate
type SetDcdnDomainCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDcdnDomainCertificateRequest creates a request to invoke SetDcdnDomainCertificate API
func CreateSetDcdnDomainCertificateRequest() (request *SetDcdnDomainCertificateRequest) {
	request = &SetDcdnDomainCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "SetDcdnDomainCertificate", "dcdn", "openAPI")
	return
}

// CreateSetDcdnDomainCertificateResponse creates a response to parse from SetDcdnDomainCertificate response
func CreateSetDcdnDomainCertificateResponse() (response *SetDcdnDomainCertificateResponse) {
	response = &SetDcdnDomainCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}