# eZmaxAPI\ObjectCommunicationAPI

All URIs are relative to *https://prod.api.appcluster01.ca-central-1.ezmax.com/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationSendV1**](ObjectCommunicationAPI.md#CommunicationSendV1) | **Post** /1/object/communication/send | Send a new Communication



## CommunicationSendV1

> CommunicationSendV1Response CommunicationSendV1(ctx).CommunicationSendV1Request(communicationSendV1Request).Execute()

Send a new Communication



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ezmaxinc/ezmax-sdk-go"
)

func main() {
	communicationSendV1Request := *openapiclient.NewCommunicationSendV1Request([]openapiclient.CommunicationRequestCompound{*openapiclient.NewCommunicationRequestCompound(openapiclient.Field-eCommunicationType("Email"), "TCommunicationBody_example", false, []openapiclient.CustomCommunicationattachmentRequest{*openapiclient.NewCustomCommunicationattachmentRequest()}, []openapiclient.CommunicationrecipientRequestCompound{*openapiclient.NewCommunicationrecipientRequestCompound()}, []openapiclient.CommunicationreferenceRequestCompound{*openapiclient.NewCommunicationreferenceRequestCompound()}, []openapiclient.CommunicationexternalrecipientRequestCompound{*openapiclient.NewCommunicationexternalrecipientRequestCompound()})}) // CommunicationSendV1Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectCommunicationAPI.CommunicationSendV1(context.Background()).CommunicationSendV1Request(communicationSendV1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectCommunicationAPI.CommunicationSendV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommunicationSendV1`: CommunicationSendV1Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectCommunicationAPI.CommunicationSendV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationSendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **communicationSendV1Request** | [**CommunicationSendV1Request**](CommunicationSendV1Request.md) |  | 

### Return type

[**CommunicationSendV1Response**](CommunicationSendV1Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

